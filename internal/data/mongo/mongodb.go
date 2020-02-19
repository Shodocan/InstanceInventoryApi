package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/Shodocan/InstanceInventoryApi/configs"
	"github.com/Shodocan/InstanceInventoryApi/internal/data"
	"github.com/Shodocan/InstanceInventoryApi/internal/logger"
)

type DB struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewDB() data.Database {
	db := &DB{}
	db.init()
	return db
}

func (db *DB) init() {
	config := configs.GetMongoDB()
	client, err := mongo.NewClient(options.Client().ApplyURI(config.GetConnectionString()))
	if err != nil {
		panic(err)
	}
	db.client = client
}

func (db *DB) Connect(collection string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := db.client.Connect(ctx)
	if err != nil {
		return err
	}
	logger.Infof("Conectado ao mongo %s %s", configs.GetMongoDB().Host, configs.GetMongoDB().DB)
	db.collection = db.client.Database(configs.GetMongoDB().DB).Collection(collection)
	return nil
}

func (db *DB) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return db.client.Ping(ctx, readpref.Primary())
}

func (db DB) Aggregate(query, result interface{}) error {
	if db.collection == nil {
		return fmt.Errorf(
			"no connection to the databse. Please connect using Connect method before using any operation method")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := db.collection.Aggregate(ctx, query)
	if err != nil {
		return err
	}
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return cursor.All(ctx, result)
}
