package counter_service

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
)

type Counters struct {
	ID             string `bson:"_id" json:"id"`                        // autoincrement field name primary key
	Sequence_value int64  `bson:"sequence_value" json:"sequence_value"` // Name of the auto-value-added field
}

type CountersService interface {
	GetCollection() *qmgo.Collection
	GetNextSequenceValue(sequenceName string) (int64, error)
	InsertOrUpdateSequenceValueByKey(key string, value int64) (*Counters, error)
}

type countersService struct {
	SetCounters *qmgo.Collection
}

var (
	countersServiceInstance CountersService
	countersOnce            sync.Once
)

func CountersServiceInstance() CountersService {
	countersOnce.Do(func() {
		countersServiceInstance = &countersService{
			SetCounters: global.GVA_MONGO.Database.Collection("counters"),
		}
	})

	return countersServiceInstance
}

func (c *countersService) GetCollection() *qmgo.Collection {
	return c.SetCounters
}

func (c *countersService) GetNextSequenceValue(sequenceName string) (int64, error) {
	var counter Counters
	cloneCollection, err := c.SetCounters.CloneCollection()
	if err != nil {
		log.Println(err.Error())
		return -1, err
	}
	findOneAndUpdate := cloneCollection.FindOneAndUpdate(context.TODO(),
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
	err = findOneAndUpdate.Decode(&counter)
	if err != nil {
		log.Println(err.Error())
		return -1, err
	}

	return counter.Sequence_value, nil
}

func (c *countersService) InsertOrUpdateSequenceValueByKey(key string, value int64) (*Counters, error) {
	cloneCollection, err := c.SetCounters.CloneCollection()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	update := cloneCollection.FindOneAndUpdate(context.TODO(),
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
	var counters Counters
	err = update.Decode(&counters)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &counters, nil
}
