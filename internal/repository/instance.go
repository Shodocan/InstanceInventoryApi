package repository

import (
	"github.com/Shodocan/InstanceInventoryApi/internal/data"
	"github.com/Shodocan/InstanceInventoryApi/internal/data/mongo"
	"github.com/Shodocan/InstanceInventoryApi/internal/entity"
	"github.com/Shodocan/InstanceInventoryApi/internal/logger"
	"github.com/Shodocan/InstanceInventoryApi/internal/util"
	"gopkg.in/mgo.v2/bson"
)

var instance InstanceRepository

type Instance struct {
	db data.Database
}

type MetricExtractor func(hostname, field string) (float64, error)

func NewInstanceRepository() (InstanceRepository, error) {
	if instance == nil {
		db := mongo.NewDB()
		err := db.Connect("instances")
		instance = &Instance{db: db}
		return instance, err
	}
	if instance.Ping() != nil {
		instance = nil
		return NewInstanceRepository()
	}
	return instance, nil
}

func (i *Instance) Statistics(hostname string, fields ...entity.MetricField) *entity.Instance {
	instace := new(entity.Instance)
	for _, field := range fields {
		switch field {
		case entity.MetricCPU:
			i.updateMetric(hostname, field, &instace.CPU)
		case entity.MetricMemory:
			i.updateMetric(hostname, field, &instace.Memory)
		case entity.MetricDisk:
			i.updateMetric(hostname, field, &instace.Disk)
		}
	}
	return instace
}

func (i *Instance) updateMetric(hostname string, field entity.MetricField, metric *entity.Metric) {
	metric.Mean = i.extractMetric(hostname, field.MongoField(), i.extractMean)
	metric.Median = i.extractMetric(hostname, field.MongoField(), i.extractMedian)
	metric.Mode = i.extractMetric(hostname, field.MongoField(), i.extractMode)
	metric.Unit = i.extractUnit(hostname, field.MongoField())
}

func (i *Instance) Ping() error {
	return i.db.Ping()
}

func (i *Instance) extractMetric(hostname, field string, extractor MetricExtractor) float64 {
	val, err := extractor(hostname, field)
	logger.ErrorIf("Error when extracting metric %s of %s", err, hostname, field)
	return val
}

func (i *Instance) extractUnit(hostname, field string) string {
	var unit string = ""
	result := []map[string]interface{}{}
	err := i.db.Aggregate([]bson.M{
		{"$match": bson.M{"hostname": hostname}},
		{"$replaceRoot": bson.M{"newRoot": field}},
		{"$limit": 1},
	}, &result)
	logger.ErrorIf("Error when extracting unit", err)
	logger.Debugf("Mode result: -> %s", util.ToJSON(result))
	if len(result) > 0 {
		unit, _ = result[0]["Unit"].(string)
	}
	return unit
}

func (i *Instance) extractMode(hostname, field string) (float64, error) {
	result := []map[string]interface{}{}
	err := i.db.Aggregate([]bson.M{
		{"$match": bson.M{"hostname": hostname}},
		{"$replaceRoot": bson.M{"newRoot": field}},
		{"$project": bson.M{"truncatedValue": bson.M{"$trunc": []interface{}{"$Value", 2}}}},
		{"$group": bson.M{"_id": "$truncatedValue", "count": bson.M{"$sum": 1}}},
		{"$group": bson.M{"_id": "$count", "val": bson.M{"$avg": "$_id"}}},
		{"$sort": bson.M{"_id": -1}},
		{"$limit": 1},
	}, &result)
	logger.ErrorIf("Error when extracting mode", err)
	logger.Debugf("Mode result: -> %s", util.ToJSON(result))
	mean := i.collectResult("val", result)
	return mean, err
}

func (i *Instance) extractMedian(hostname, field string) (float64, error) {
	result := []map[string]interface{}{}
	var count = i.count(hostname)
	err := i.db.Aggregate([]bson.M{
		{"$match": bson.M{"hostname": hostname}},
		{"$replaceRoot": bson.M{"newRoot": field}},
		{"$sort": bson.M{"Value": 1}},
		{"$skip": count/2 - 1},
		{"$limit": 1},
	}, &result)
	logger.ErrorIf("Error when extracting median", err)
	logger.Debugf("Median result: -> %s", util.ToJSON(result))
	mean := i.collectResult("Value", result)
	return mean, err
}

func (i *Instance) extractMean(hostname, field string) (float64, error) {
	result := []map[string]interface{}{}
	err := i.db.Aggregate([]bson.M{
		{"$match": bson.M{"hostname": hostname}},
		{"$replaceRoot": bson.M{"newRoot": field}},
		{"$group": bson.M{"_id": hostname, "val": bson.M{"$avg": "$Value"}}},
		{"$unset": "_id"},
	}, &result)
	logger.ErrorIf("Error when extracting mean", err)
	logger.Debugf("Mean result: -> %s", util.ToJSON(result))
	mean := i.collectResult("val", result)
	return mean, err
}

func (i *Instance) count(hostname string) int32 {
	result := []map[string]interface{}{}
	err := i.db.Aggregate([]bson.M{
		{"$match": bson.M{"hostname": hostname}},
		{"$group": bson.M{"_id": nil, "val": bson.M{"$sum": 1}}},
	}, &result)
	logger.ErrorIf("Error counting entries of %s ", err, hostname)
	logger.Debugf("Count result: -> %s", util.ToJSON(result))
	mean := i.collectResult("val", result)
	return int32(mean)
}

func (i *Instance) collectResult(key string, result []map[string]interface{}) float64 {
	if len(result) > 0 {
		switch mean := result[0][key].(type) {
		case float64:
			return mean
		case int32:
			return float64(mean)
		default:
			logger.Debugf("Error parsing result %s", key)
		}
	}
	return 0
}
