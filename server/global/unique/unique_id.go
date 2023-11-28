package unique_service

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
)

const (
	AdminModuleControl = "admin_module_control"
)

type uniqueId struct {
	ID            string `bson:"_id" json:"id"`                        // autoincrement field name primary key
	SequenceValue int64  `bson:"sequence_value" json:"sequence_value"` // Name of the auto-value-added field
}

type UniqueService interface {
	GetCollection() *mongo.Collection
	GetNextSequenceValue(sequenceName string) (int64, error)
	InsertOrUpdateSequenceValueByKey(key string, value int64) (*uniqueId, error)
}

type uniqueService struct {
	SetUnique *mongo.Collection
}

var (
	uniqueServiceInstance UniqueService
	uniqueOnce            sync.Once
)

func UniqueServiceInstance() UniqueService {
	uniqueOnce.Do(func() {
		uniqueServiceInstance = &uniqueService{
			SetUnique: global.GVA_MONGODB.Collection("admin_unique_id"),
		}
	})

	return uniqueServiceInstance
}

func (c *uniqueService) GetCollection() *mongo.Collection {
	return c.SetUnique
}

func (c *uniqueService) GetNextSequenceValue(sequenceName string) (int64, error) {

	findOneAndUpdate := c.SetUnique.FindOneAndUpdate(context.TODO(),
		bson.D{{"_id", sequenceName}},
		bson.D{{"$inc", bson.D{{"sequence_value", 1}}}},
		options.FindOneAndUpdate().
			SetUpsert(true).
			SetReturnDocument(options.After),
	)
	if findOneAndUpdate.Err() != nil {
		log.Println(findOneAndUpdate.Err().Error())
		return -1, findOneAndUpdate.Err()
	}
	var uniqueId uniqueId
	err := findOneAndUpdate.Decode(&uniqueId)
	if err != nil {
		log.Println(err.Error())
		return -1, err
	}

	return uniqueId.SequenceValue, nil
}

func (c *uniqueService) InsertOrUpdateSequenceValueByKey(key string, value int64) (*uniqueId, error) {

	update := c.SetUnique.FindOneAndUpdate(context.TODO(),
		bson.D{{"_id", key}},
		bson.D{
			{"$setOnInsert", bson.D{
				//{"sequence_value", value},
			}}, {"$set", bson.D{
				{"sequence_value", value},
			}}},
		options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After))

	if update.Err() != nil {
		fmt.Println(update.Err().Error())
		return nil, update.Err()
	}
	var uid uniqueId
	err := update.Decode(&uid)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &uid, nil
}
