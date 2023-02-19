package persistance

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseConfig struct {
	Uri               string `json:"uri"`
	RequestTimeoutSec int64  `json:"requestTimeoutSec"`
}

type Database struct {
	config DatabaseConfig
	client *mongo.Client
}

func NewDatabase(config DatabaseConfig) *Database {
	db := &Database{}
	db.config = config
	db.connect()
	return db
}

func (db *Database) connect() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(db.config.RequestTimeoutSec)*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(db.config.Uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Mongodb: can't establish connection: ", err)
	}
	db.client = client
}
