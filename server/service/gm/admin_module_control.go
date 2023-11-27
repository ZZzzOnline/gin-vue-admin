package gm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gm"
	gmReq "github.com/flipped-aurora/gin-vue-admin/server/model/gm/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	UserCollection               = "user"
	AdminModuleControlCollection = "admin_module_control"
)

type AdminModuleControlService struct {
}

func (adminModuleControlService *AdminModuleControlService) InitAdminModuleControlIndexes() error {
	adminModuleControlCollection := global.GVA_MONGODB.Collection(AdminModuleControlCollection)
	{
		indexOpts := new(options.IndexOptions)
		indexOpts.SetName("_id_")
		indexName, err := adminModuleControlCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    bson.D{{"_id", 1}},
			Options: indexOpts,
		})
		if err != nil {
			return err
		}
		log.Println("created index:", indexName)
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
	}
	{
		indexOpts := new(options.IndexOptions)
		indexOpts.SetName("AccountId_").SetUnique(true)
		indexName, err := adminModuleControlCollection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    bson.D{{"AccountId", 1}},
			Options: indexOpts,
		})
		if err != nil {
			return err
		}
		log.Println("created index:", indexName)
	}

	return nil
}

// CreateAdminModuleControl 创建模块控制记录
// Author [piexlmax](https://github.com/piexlmax)
func (adminModuleControlService *AdminModuleControlService) CreateAdminModuleControl(adminModuleControl *gm.AdminModuleControl) (err error) {
	type AccountResponse struct {
		ID         uint   `json:"id"`
		Wallet     []int  `json:"wallet"`
		Email      string `json:"email"`
		PlatformID int    `json:"platformId"`
		OpenID     string `json:"openId"`
		PP         string `json:"pp"`
		Ext        string `json:"ext"`
		CreatedAt  int    `json:"createdAt"`
		UpdatedAt  int    `json:"updatedAt"`
	}
	resp, err := http.DefaultClient.Get("https://tomato-seascape-auth.blocklords.com/api/account/email/zhaojihuionline@outlook.com")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %v", err)
		}
		return fmt.Errorf("unexpected status code: %d %s. Body: %s", resp.StatusCode, resp.Status, string(bodyBytes))
	}

	var accountResponse AccountResponse
	err = json.NewDecoder(resp.Body).Decode(&accountResponse)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	session, err := global.GVA_MONGO.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		type User struct {
			PlayerId  primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
			AccountId int64              `bson:"AccountId" json:"AccountId"` // primary key
		}
		var user User

		opts := options.FindOne().SetSort(bson.D{{"AccountId", 1}})
		err = global.GVA_MONGODB.Collection(UserCollection).FindOne(sessCtx, bson.M{"AccountId": accountResponse.ID}, opts).Decode(&user)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, fmt.Errorf("can not find the user in mongodb. err:%v", err)
			}
			return nil, err
		}

		adminModuleControl.ID = accountResponse.ID
		adminModuleControl.AccountId = &accountResponse.ID
		adminModuleControl.PlayerId = user.PlayerId.Hex()
		adminModuleControl.CreatedAt = time.Now().UTC()
		adminModuleControl.UpdatedAt = time.Now().UTC()

		_, err = global.GVA_MONGODB.Collection(AdminModuleControlCollection).InsertOne(sessCtx, adminModuleControl)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	return err
}

// DeleteAdminModuleControl 删除模块控制记录
// Author [piexlmax](https://github.com/piexlmax)
func (adminModuleControlService *AdminModuleControlService) DeleteAdminModuleControl(adminModuleControl gm.AdminModuleControl) (err error) {
	//err = global.GVA_MONGO.Transaction(func(tx *gorm.DB) error {
	//	if err := tx.Model(&gm.AdminModuleControl{}).Where("id = ?", adminModuleControl.ID).Update("deleted_by", adminModuleControl.DeletedBy).Error; err != nil {
	//		return err
	//	}
	//	if err = tx.Delete(&adminModuleControl).Error; err != nil {
	//		return err
	//	}
	//	return nil
	//})
	return err
}

// DeleteAdminModuleControlByIds 批量删除模块控制记录
// Author [piexlmax](https://github.com/piexlmax)
func (adminModuleControlService *AdminModuleControlService) DeleteAdminModuleControlByIds(ids request.IdsReq, deleted_by uint) (err error) {
	//err = global.GVA_MONGO.Transaction(func(tx *gorm.DB) error {
	//	if err := tx.Model(&gm.AdminModuleControl{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
	//		return err
	//	}
	//	if err := tx.Where("id in ?", ids.Ids).Delete(&gm.AdminModuleControl{}).Error; err != nil {
	//		return err
	//	}
	//	return nil
	//})
	return err
}

// UpdateAdminModuleControl 更新模块控制记录
// Author [piexlmax](https://github.com/piexlmax)
func (adminModuleControlService *AdminModuleControlService) UpdateAdminModuleControl(adminModuleControl gm.AdminModuleControl) (err error) {
	//err = global.GVA_MONGO.Save(&adminModuleControl).Error
	if adminModuleControl.ID != *adminModuleControl.AccountId {
		return fmt.Errorf("adminModuleControl.ID:%d != *adminModuleControl.AccountId:%d", adminModuleControl.ID, adminModuleControl.AccountId)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	_, err = global.GVA_MONGODB.Collection(AdminModuleControlCollection).UpdateByID(ctx, adminModuleControl.ID, bson.M{
		"$set": adminModuleControl,
	})

	return err
}

// GetAdminModuleControl 根据id获取模块控制记录
// Author [piexlmax](https://github.com/piexlmax)
func (adminModuleControlService *AdminModuleControlService) GetAdminModuleControl(id uint) (adminModuleControl gm.AdminModuleControl, err error) {
	//err = global.GVA_MONGO.Where("id = ?", id).First(&adminModuleControl).Error
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	opts := options.FindOne().SetSort(bson.D{{"AccountId", 1}})
	err = global.GVA_MONGODB.Collection(AdminModuleControlCollection).FindOne(ctx, bson.M{"AccountId": id}, opts).Decode(&adminModuleControl)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return adminModuleControl, fmt.Errorf("can not find the user in mongodb. err:%v", err)
		}
		return adminModuleControl, err
	}
	return
}

// GetAdminModuleControlInfoList 分页获取模块控制记录
// Author [piexlmax](https://github.com/piexlmax)
func (adminModuleControlService *AdminModuleControlService) GetAdminModuleControlInfoList(info gmReq.AdminModuleControlSearch) (list []gm.AdminModuleControl, total int64, err error) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 条件搜索
	filter := bson.M{}
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		filter["CreatedAt"] = bson.M{"$gte": info.StartCreatedAt, "$lte": info.EndCreatedAt}
	}
	if info.Email != "" {
		filter["Email"] = bson.M{"$regex": info.Email}
	}
	if info.AccountId != nil {
		filter["AccountId"] = info.AccountId
	}
	if info.PlayerId != "" {
		filter["PlayerId"] = bson.M{"$regex": info.PlayerId}
	}
	if info.ForbiddenLogin != nil {
		filter["ForbiddenLogin"] = info.ForbiddenLogin
	}
	if info.ForbiddenInGameHeroExport != nil {
		filter["ForbiddenInGameHeroExport"] = info.ForbiddenInGameHeroExport
	}

	var adminModuleControls []gm.AdminModuleControl

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	session, err := global.GVA_MONGO.StartSession()
	if err != nil {
		return nil, 0, err
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		total, err = global.GVA_MONGODB.Collection(AdminModuleControlCollection).CountDocuments(sessCtx, filter)
		if err != nil {
			return nil, err
		}

		opts := options.Find()
		if limit != 0 {
			opts.SetLimit(int64(limit))
			opts.SetSkip(int64(offset))
		}
		cursor, err := global.GVA_MONGODB.Collection(AdminModuleControlCollection).Find(sessCtx, filter, opts)
		if err != nil {
			return nil, err
		}

		err = cursor.All(sessCtx, &adminModuleControls)

		return nil, nil
	})

	return adminModuleControls, total, err
}
