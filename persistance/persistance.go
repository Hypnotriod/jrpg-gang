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
	mu                   sync.RWMutex
	rndGen               *util.RndGen
	config               PersistanceConfig
	db                   *MongoDB
	userEmailToAuthToken map[model.UserEmail]auth.AuthenticationToken
	usersAuthCache       *ttlcache.Cache[auth.AuthenticationToken, *model.UserModel]
}

func NewPersistance(config PersistanceConfig) *Persistance {
	p := &Persistance{}
	p.config = config
	p.rndGen = util.NewRndGen()
	p.userEmailToAuthToken = make(map[model.UserEmail]auth.AuthenticationToken)
	p.usersAuthCache = ttlcache.New(
		ttlcache.WithTTL[auth.AuthenticationToken, *model.UserModel](time.Duration(config.PersistanceCacheTimeoutMin) * time.Minute),
	)
	p.usersAuthCache.OnEviction(p.onUserAuthCacheEviction)
	go p.usersAuthCache.Start()
	p.db = NewMongoDB(config.MongoDBConfig)
	return p
}

func (p *Persistance) AddUserToAuthCache(userModel *model.UserModel) auth.AuthenticationToken {
	p.mu.Lock()
	if oldToken, ok := p.userEmailToAuthToken[userModel.Email]; ok {
		p.usersAuthCache.Delete(oldToken)
	}
	token := auth.AuthenticationToken(p.rndGen.MakeUUID())
	p.userEmailToAuthToken[userModel.Email] = token
	p.usersAuthCache.Set(token, userModel, ttlcache.DefaultTTL)
	p.mu.Unlock()
	return token
}

func (p *Persistance) GetUserFromAuthCache(token auth.AuthenticationToken) (*model.UserModel, bool) {
	defer p.mu.RUnlock()
	p.mu.RLock()
	item := p.usersAuthCache.Get(token)
	if item == nil || item.IsExpired() {
		return nil, false
	}
	return item.Value(), true
}

func (p *Persistance) RemoveUserFromAuthCache(token auth.AuthenticationToken) {
	defer p.mu.Unlock()
	p.mu.Lock()
	item := p.usersAuthCache.Get(token)
	if item != nil {
		email := item.Value().Email
		delete(p.userEmailToAuthToken, email)
		p.usersAuthCache.Delete(token)
	}
}

func (p *Persistance) onUserAuthCacheEviction(ctx context.Context, reson ttlcache.EvictionReason, item *ttlcache.Item[auth.AuthenticationToken, *model.UserModel]) {
	if reson != ttlcache.EvictionReasonDeleted {
		defer p.mu.Unlock()
		p.mu.Lock()
		email := item.Value().Email
		delete(p.userEmailToAuthToken, email)
	}
}

func (p *Persistance) HasUserWithNickname(nickname string) bool {
	ctx, cancel := p.db.requestContext()
	defer cancel()
	user, ok := p.db.UsersRepository.FindByNickname(ctx, nickname)
	return user != nil && ok
}

func (p *Persistance) GetUserByEmail(email model.UserEmail) *model.UserModel {
	ctx, cancel := p.db.requestContext()
	defer cancel()
	user, ok := p.db.UsersRepository.FindByEmail(ctx, email)
	if !ok {
		return nil
	}
	return user
}

func (p *Persistance) GetOrCreateUser(credentials auth.UserCredentials) *model.UserModel {
	ctx, cancel := p.db.requestContext()
	defer cancel()
	user, ok := p.db.UsersRepository.FindByEmail(ctx, model.UserEmail(credentials.Email))
	if !ok {
		return nil
	}
	if user != nil {
		return user
	}
	userToPersist := model.UserModel{
		Email:   model.UserEmail(credentials.Email),
		Picture: credentials.Picture,
	}
	id, ok := p.db.UsersRepository.InsertOne(ctx, userToPersist)
	if !ok {
		return nil
	}
	user, ok = p.db.UsersRepository.FindOneById(ctx, id, &model.UserModel{})
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

func (p *Persistance) UpdateJobStatus(jobStatus model.JobStatusModel) bool {
	ctx, cancel := p.db.requestContext()
	defer cancel()
	matchedCount, ok := p.db.JobStatusRepository.UpdateOrInsertOne(ctx, jobStatus)
	return matchedCount != 0 && ok
}

func (p *Persistance) GetJobStatus(userId model.UserId) *model.JobStatusModel {
	ctx, cancel := p.db.requestContext()
	defer cancel()
	jobStatus, _ := p.db.JobStatusRepository.FindByUserId(ctx, userId)
	return jobStatus
}
