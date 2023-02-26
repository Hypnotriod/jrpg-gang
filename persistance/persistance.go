package persistance

import (
	"jrpg-gang/auth"
	"jrpg-gang/persistance/model"
	"jrpg-gang/util"
	"time"

	"github.com/jellydator/ttlcache/v3"
)

type Persistance struct {
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
	go p.usersCache.Start()
	p.db = NewMongoDB(dbConfig)
	return p
}

func (p *Persistance) AddUserToCache(userModel *model.UserModel) auth.AuthenticationToken {
	token := auth.AuthenticationToken(p.rndGen.MakeUUID())
	p.usersCache.Set(token, userModel, ttlcache.DefaultTTL)
	return token
}

func (p *Persistance) GetUserFromCache(token auth.AuthenticationToken) (*model.UserModel, bool) {
	item := p.usersCache.Get(token)
	if item == nil || item.IsExpired() {
		return nil, false
	}
	return item.Value(), true
}

func (p *Persistance) RemoveUserFromCache(token auth.AuthenticationToken) {
	p.usersCache.Delete(token)
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
