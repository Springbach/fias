package main

import (
	"context"
	"encoding/json"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Aggregator interface {
	Get(map[string]string) ([]byte, bool, error)
}

type Regions struct{}
type Cities struct{}
type Streets struct{}
type Territory struct{}
type Houses struct{}

func (r Regions) Get(q map[string]string) ([]byte, bool, error) {
	pipe := mongo.Pipeline{
		{{"$group", bson.M{"_id": "$reg"}}},
		{{"$sort", bson.M{"_id": 1}}},
		{{"$group", bson.M{"_id": "null", "regions": bson.M{"$push": "$_id"}}}},
		{{"$project", bson.M{"_id": 0}}},
	}
	//set Aggregation Timeout of 1 second
	value, timeout, err := aggregate(pipe, 1*time.Second)
	if err != nil {
		return []byte(""), false, err
	}
	if timeout {
		return []byte(""), true, nil
	}
	result, _ := json.Marshal(value)
	return result, false, nil
}

func (t Territory) Get(q map[string]string) ([]byte, bool, error) {
	pipe := mongo.Pipeline{
		{{"$match", bson.M{"zip": q["zip"]}}},
		{{"$project", bson.M{"_id": 0, "hs": 0}}},
		{{"$project", bson.M{"city": bson.M{"$arrayElemAt": bson.A{"$adr", 1}}, "street": bson.M{"$arrayElemAt": bson.A{"$adr", 0}}}}},
		{{"$sort", bson.M{"street": 1}}},
		{{"$group", bson.M{"_id": "null", "city": bson.M{"$first": "$city"}, "streets": bson.M{"$push": "$street"}}}},
		{{"$project", bson.M{"_id": 0}}},
	}
	//set Aggregation Timeout of 1 second
	value, timeout, err := aggregate(pipe, 1*time.Second)
	if err != nil {
		return []byte(""), false, err
	}
	if timeout {
		return []byte(""), true, nil
	}
	result, _ := json.Marshal(value)
	return result, false, nil
}

func (s Streets) Get(q map[string]string) ([]byte, bool, error) {
	pipe := mongo.Pipeline{
		{{"$match", bson.M{"reg": q["reg"], "adr.1": q["city"]}}},
		{{"$project", bson.M{"_id": 0, "hs": 0}}},
		{{"$project", bson.M{"street": bson.M{"$arrayElemAt": bson.A{"$adr", 0}}}}},
		{{"$group", bson.M{"_id": "$street"}}},
		{{"$sort", bson.M{"_id": 1}}},
		{{"$group", bson.M{"_id": "null", "streets": bson.M{"$push": "$_id"}}}},
		{{"$project", bson.M{"_id": 0}}},
	}
	//set Aggregation Timeout of 1 second
	value, timeout, err := aggregate(pipe, 1*time.Second)
	if err != nil {
		return []byte(""), false, err
	}
	if timeout {
		return []byte(""), true, nil
	}
	result, _ := json.Marshal(value)
	return result, false, nil
}

func (c Cities) Get(q map[string]string) ([]byte, bool, error) {
	pipe := mongo.Pipeline{
		{{"$match", bson.M{"reg": q["reg"]}}},
		{{"$project", bson.M{"_id": 0, "hs": 0}}},
		{{"$project", bson.M{"city": bson.M{"$arrayElemAt": bson.A{"$adr", 1}}}}},
		{{"$group", bson.M{"_id": "$city"}}},
		{{"$sort", bson.M{"_id": 1}}},
		{{"$group", bson.M{"_id": "null", "cities": bson.M{"$push": "$_id"}}}},
		{{"$project", bson.M{"_id": 0}}},
	}
	//set Aggregation Timeout of 1 second
	value, timeout, err := aggregate(pipe, 1*time.Second)
	if err != nil {
		return []byte(""), false, err
	}
	if timeout {
		return []byte(""), true, nil
	}
	result, _ := json.Marshal(value)
	return result, false, nil
}

func (h Houses) Get(q map[string]string) ([]byte, bool, error) {
	pipe := mongo.Pipeline{
		{{"$match", bson.M{"zip": q["zip"], "adr.0": q["street"]}}},
		{{"$project", bson.M{"hs": 1}}},
		{{"$project", bson.M{"_id": 0}}},
	}
	//set Aggregation Timeout of 1 second
	value, timeout, err := aggregate(pipe, 1*time.Second)
	if err != nil {
		return []byte(""), false, err
	}
	if timeout {
		return []byte(""), true, nil
	}
	result, _ := json.Marshal(value)
	return result, false, nil
}

func aggregate(pipe mongo.Pipeline, timeout time.Duration) (bson.M, bool, error) {
	streets := client.Database("fias").Collection("streets")
	cursor, err := streets.Aggregate(context.Background(), pipe, &options.AggregateOptions{MaxTime: &timeout})
	if err != nil {
		//check if err is a timeout or mongodb command error
		if reflect.TypeOf(err) == reflect.TypeOf(mongo.CommandError{}) {
			return bson.M{}, true, nil
		}
		return bson.M{}, false, err
	}
	defer cursor.Close(context.Background())
	cursor.Next(context.Background())
	var aggr bson.M
	err = cursor.Decode(&aggr)
	if err != nil {
		return bson.M{}, false, err
	}
	return aggr, false, nil
}
