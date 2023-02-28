package persistance

import (
	"context"
	"jrpg-gang/auth"
	"jrpg-gang/persistance/model"
	"jrpg-gang/util"
	"sync"
	"time"

	"github.com/jellydator/ttlcache/v3"
)

type PersistanceConfig struct {
	MongoDBConfig              MongoDBConfig `json:"mongoDBConfig,omitempty"`
	PersistanceCacheTimeoutMin int64
}

type Persistance struct {
	mu               sync.RWMutex
	rndGen           *util.RndGen
	config           PersistanceConfig
	db               *MongoDB
	userEmailToToken map[string]auth.AuthenticationToken
	usersCache       *ttlcache.Cache[auth.AuthenticationToken, *model.UserModel]
}

func NewPersistance(config PersistanceConfig) *Persistance {
	p := &Persistance{}
	p.config = config
	p.rndGen = util.NewRndGen()
	p.userEmailToToken = make(map[string]auth.AuthenticationToken)
	p.usersCache = ttlcache.New(
		ttlcache.WithTTL[auth.AuthenticationToken, *model.UserModel](time.Duration(config.PersistanceCacheTimeoutMin) * time.Minute),
	)
	p.usersCache.OnEviction(p.onUserCacheEviction)
	go p.usersCache.Start()
	p.db = NewMongoDB(config.MongoDBConfig)
	return p
}

func (p *Persistance) AddUserToCache(userModel *model.UserModel) auth.AuthenticationToken {
	p.mu.Lock()
	if oldToken, ok := p.userEmailToToken[userModel.Email]; ok {
		p.usersCache.Delete(oldToken)
	}
	token := auth.AuthenticationToken(p.rndGen.MakeUUID())
	p.userEmailToToken[userModel.Email] = token
	p.usersCache.Set(token, userModel, ttlcache.DefaultTTL)
	p.mu.Unlock()
	return token
}

func (p *Persistance) GetUserFromCache(token auth.AuthenticationToken) (*model.UserModel, bool) {
	defer p.mu.RUnlock()
	p.mu.RLock()
	item := p.usersCache.Get(token)
	if item == nil || item.IsExpired() {
		return nil, false
	}
	return item.Value(), true
}

func (p *Persistance) RemoveUserFromCache(token auth.AuthenticationToken) {
	defer p.mu.Unlock()
	p.mu.Lock()
	item := p.usersCache.Get(token)
	if item != nil {
		email := item.Value().Email
		delete(p.userEmailToToken, email)
		p.usersCache.Delete(token)
	}
}

func (p *Persistance) onUserCacheEviction(ctx context.Context, reson ttlcache.EvictionReason, item *ttlcache.Item[auth.AuthenticationToken, *model.UserModel]) {
	if reson != ttlcache.EvictionReasonDeleted {
		defer p.mu.Unlock()
		p.mu.Lock()
		email := item.Value().Email
		delete(p.userEmailToToken, email)
	}
}

func (p *Persistance) HasUserWithNickname(nickname string) bool {
	ctx, cancel := p.db.requestContext()
	defer cancel()
	user, ok := p.db.UsersRepository.FindByNickname(ctx, nickname)
	return user != nil && ok
}

func (p *Persistance) GetOrCreateUser(creadentials auth.UserCredentials) *model.UserModel {
	ctx, cancel := p.db.requestContext()
	defer cancel()
	user, ok := p.db.UsersRepository.FindByEmail(ctx, creadentials.Email)
	if !ok {
		return nil
	}
	if user != nil {
		return user
	}
	userToPersist := model.UserModel{
		Email:   creadentials.Email,
		Picture: creadentials.Picture,
	}
	id, ok := p.db.UsersRepository.InsertOne(ctx, userToPersist)
	if !ok {
		return nil
	}
	user, ok = p.db.UsersRepository.FindOneById(ctx, id)
	if !ok {
		return nil
	}
	return user
}

func (p *Persistance) UpdateUser(user model.UserModel) bool {
	ctx, cancel := p.db.requestContext()
	defer cancel()
	updated, ok := p.db.UsersRepository.UpdateOneWithUnit(ctx, user)
	return updated != 0 && ok
}
