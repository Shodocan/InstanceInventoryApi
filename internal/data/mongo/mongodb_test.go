package mongo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

// This test supose that MongoDb instance is Running and the testing data is loaded on the database

func checkMongo(t *testing.T) bool {
	if os.Getenv("DB_HOST") == "" {
		return false
	}
	return true
}

func TestMongoConnect(t *testing.T) {
	if !checkMongo(t) {
		t.Skip()
		return
	}
	mongo := NewDB()
	err := mongo.Connect("instances")
	assert.Nil(t, err, "falha ao conectar no mongo", err)
}

func TestPing(t *testing.T) {
	if !checkMongo(t) {
		t.Skip()
		return
	}
	mongo := NewDB()
	err := mongo.Connect("instances")
	assert.Nil(t, err, "falha ao conectar no mongo", err)
	assert.Nil(t, mongo.Ping(), "Deveria conseguir fazer um Ping")
}

func TestNotConnected(t *testing.T) {
	if !checkMongo(t) {
		t.Skip()
		return
	}
	mongo := NewDB()
	result := []map[string]interface{}{}
	err := mongo.Aggregate([]bson.M{
		bson.M{
			"$match": bson.M{"hostname": "server3"},
		},
	}, &result)
	assert.NotNil(t, err, "Deveria retornar um erro")
}

func TestInvalidQuery(t *testing.T) {
	if !checkMongo(t) {
		t.Skip()
		return
	}
	mongo := NewDB()
	err := mongo.Connect("instances")
	assert.Nil(t, err, "falha ao conectar no mongo", err)
	result := []map[string]interface{}{}
	err = mongo.Aggregate([]bson.M{
		bson.M{
			"match": bson.M{"hostname": "server3"},
		},
	}, &result)
	assert.NotNil(t, err, "Deveria retornar um erro")
}

func TestMongoAggregate(t *testing.T) {
	if !checkMongo(t) {
		t.Skip()
		return
	}
	mongo := NewDB()
	err := mongo.Connect("instances")
	assert.Nil(t, err, "falha ao conectar no mongo", err)

	result := []map[string]interface{}{}
	err = mongo.Aggregate([]bson.M{
		bson.M{
			"$match": bson.M{"hostname": "server3"},
		},
		bson.M{
			"$replaceRoot": bson.M{"newRoot": "$cpu_load"},
		},
		bson.M{
			"$group": bson.M{"_id": "server3", "avg": bson.M{"$avg": "$Value"}},
		},
	}, &result)
	assert.Nil(t, err, "falha ao agregar dados no mongo", err)

	if len(result) != 1 {
		t.Errorf("Deveria retornar apenas um resultado")
	}
	assert.Equal(t, 0.5045931010673991, result[0]["avg"])
}
