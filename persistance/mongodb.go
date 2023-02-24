package persistance

import (
	"context"
	"jrpg-gang/persistance/repository"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConfig struct {
	Uri               string `json:"uri"`
	RequestTimeoutSec int64  `json:"requestTimeoutSec"`
}

type MongoDB struct {
	UsersRepository *repository.UserRepository
	config          MongoDBConfig
	client          *mongo.Client
}

func NewMongoDB(config MongoDBConfig) *MongoDB {
	db := &MongoDB{}
	db.config = config
	db.connect()
	db.initRepositories()
	return db
}

func (db *MongoDB) requestContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(db.config.RequestTimeoutSec)*time.Second)
}

func (db *MongoDB) connect() {
	ctx, cancel := db.requestContext()
	defer cancel()
	clientOptions := options.Client().ApplyURI(db.config.Uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Mongodb: can't establish connection: ", err)
	}
	db.client = client
}

func (db *MongoDB) initRepositories() {
	db.UsersRepository = repository.NewUserRepository(
		db.client.Database("jrpg_gang").Collection("user"),
	)
}
