package config

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type ConfigMongo struct {
	MONGO_HOST                 string
	MONGO_PORT                 string
	MONGO_USERNAME             string
	MONGGO_PASSWORD            string
	MONGO_DATABASE             string
	MONGO_POOL_MIN             int
	MONGO_POOL_MAX             int
	MONGO_MAX_IDLE_TIME_SECOND int
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func (configMongo *ConfigMongo) InitializeMongo() (*mongo.Database, error) {
	ctx, cancel := NewMongoContext()
	defer cancel()

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/", configMongo.MONGO_USERNAME, configMongo.MONGGO_PASSWORD, configMongo.MONGO_HOST, configMongo.MONGO_PORT)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI).
		SetMinPoolSize(uint64(configMongo.MONGO_POOL_MIN)).
		SetMaxPoolSize(uint64(configMongo.MONGO_POOL_MAX)).
		SetMaxConnIdleTime(time.Duration(configMongo.MONGO_MAX_IDLE_TIME_SECOND)*time.Second))

	if err != nil {
		return nil, err
	}

	// ping to check connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	database := client.Database(configMongo.MONGO_DATABASE)

	fmt.Println("Successfully connected and pinged MONGODB Database")
	return database, nil
}
