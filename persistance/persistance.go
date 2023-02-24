package persistance

import (
	"jrpg-gang/auth"
	"jrpg-gang/persistance/model"
	"jrpg-gang/util"
	"sync"
	"time"

	"github.com/jellydator/ttlcache/v3"
)

type Persistance struct {
	mu         sync.RWMutex
	rndGen     *util.RndGen
	db         *MongoDB
	usersCache *ttlcache.Cache[auth.AuthenticationToken, *model.UserModel]
}

func NewPersistance(dbConfig MongoDBConfig) *Persistance {
	p := &Persistance{}
	p.rndGen = util.NewRndGen()
	p.usersCache = ttlcache.New(
		ttlcache.WithTTL[auth.AuthenticationToken, *model.UserModel](time.Hour),
	)
	p.db = NewMongoDB(dbConfig)
	return p
}

func (p *Persistance) AddUserToCache(userModel *model.UserModel) auth.AuthenticationToken {
	defer p.mu.Unlock()
	p.mu.Lock()
	var token auth.AuthenticationToken
	for {
		token = auth.AuthenticationToken(p.rndGen.MakeHex32())
		if item := p.usersCache.Get(token); item == nil {
			break
		}
	}
	p.usersCache.Set(token, userModel, ttlcache.DefaultTTL)
	return token
}

func (p *Persistance) PopUserFromCache(token auth.AuthenticationToken) (*model.UserModel, bool) {
	p.mu.RLock()
	item := p.usersCache.Get(token)
	if item != nil {
		p.usersCache.Delete(item.Key())
	}
	p.mu.RUnlock()
	if item == nil || item.IsExpired() {
		return nil, false
	}
	return item.Value(), true
}

func (p *Persistance) GetOrCreateUser(creadentials auth.UserCredentials) *model.UserModel {
	ctx, cancel := p.db.requestContext()
	defer cancel()
	user, ok := p.db.UsersRepository.FindByEmail(ctx, creadentials.Email)
	if user != nil && ok {
		return user
	}
	userToPersist := model.UserModel{
		Email:   creadentials.Email,
		Picture: creadentials.Picture,
	}
	userToPersist.OnCreate()
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
