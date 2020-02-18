package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Shodocan/InstanceInventoryApi/configs"
	"github.com/Shodocan/InstanceInventoryApi/internal/data"
	"github.com/Shodocan/InstanceInventoryApi/internal/logger"
)

type MongoDB struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoDB() data.Database {
	db := &MongoDB{}
	db.init()
	return db
}

func (db *MongoDB) init() {
	config := configs.GetMongoDB()
	client, err := mongo.NewClient(options.Client().ApplyURI(config.GetConnectionString()))
	if err != nil {
		panic(err)
	}
	db.client = client
}

func (db *MongoDB) Connect(collection string) error {
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

func (db MongoDB) Aggregate(query interface{}, result interface{}) error {
	if db.collection == nil {
		return fmt.Errorf("No connection to the databse. Please connect using Connect method before using any operation method")
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
