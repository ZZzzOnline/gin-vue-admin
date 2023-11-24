package gm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gm"
	gmReq "github.com/flipped-aurora/gin-vue-admin/server/model/gm/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (adminModuleControlService *AdminModuleControlService) InitAdminModuleControlIndexes() {
	collection, err := global.GVA_MONGO.Database.Collection(AdminModuleControlCollection).CloneCollection()
	if err != nil {
		panic(err)
	}
	{
		indexOpts := new(options.IndexOptions)
		indexOpts.SetName("_id_")
		indexName, err := collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    bson.D{{"_id", 1}},
			Options: indexOpts,
		})
		if err != nil {
			fmt.Print(err.Error())
			panic(err)
		}
		log.Println("created index:", indexName)
	}
	{
		indexOpts := new(options.IndexOptions)
		indexOpts.SetName("AccountId_").SetUnique(true)
		indexName, err := collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    bson.D{{"AccountId", 1}},
			Options: indexOpts,
		})
		if err != nil {
			fmt.Print(err.Error())
			panic(err)
		}
		log.Println("created index:", indexName)
	}
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

	ctx := context.TODO()
	_, err = global.GVA_MONGO.DoTransaction(ctx, func(sessCtx context.Context) (interface{}, error) {
		type User struct {
			PlayerId  primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
			AccountId int64              `bson:"AccountId" json:"AccountId"` // primary key
		}
		var user User
		err = global.GVA_MONGO.Database.Collection(UserCollection).Find(ctx, bson.M{}).One(&user)
		if err != nil {
			return nil, err
		}
		adminModuleControl.AccountId = &accountResponse.ID
		adminModuleControl.PlayerId = user.PlayerId.Hex()
		adminModuleControl.CreatedAt = time.Now().UTC()
		adminModuleControl.UpdatedAt = time.Now().UTC()
		insertOneResult, err := global.GVA_MONGO.Database.Collection(AdminModuleControlCollection).InsertOne(ctx, adminModuleControl)
		if err != nil {
			return nil, err
		}
		adminModuleControl.ID = insertOneResult.InsertedID.(primitive.ObjectID)
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
	return err
}

// GetAdminModuleControl 根据id获取模块控制记录
// Author [piexlmax](https://github.com/piexlmax)
func (adminModuleControlService *AdminModuleControlService) GetAdminModuleControl(id primitive.ObjectID) (adminModuleControl gm.AdminModuleControl, err error) {
	//err = global.GVA_MONGO.Where("id = ?", id).First(&adminModuleControl).Error
	return
}

// GetAdminModuleControlInfoList 分页获取模块控制记录
// Author [piexlmax](https://github.com/piexlmax)
func (adminModuleControlService *AdminModuleControlService) GetAdminModuleControlInfoList(info gmReq.AdminModuleControlSearch) (list []gm.AdminModuleControl, total int64, err error) {
	//limit := info.PageSize
	//offset := info.PageSize * (info.Page - 1)
	//// 创建db
	//db := global.GVA_MONGO.Model(&gm.AdminModuleControl{})
	//var adminModuleControls []gm.AdminModuleControl
	//// 如果有条件搜索 下方会自动创建搜索语句
	//if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
	//	db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	//}
	//if info.Email != "" {
	//	db = db.Where("email LIKE ?", "%"+info.Email+"%")
	//}
	//if info.AccountId != nil {
	//	db = db.Where("account_id = ?", info.AccountId)
	//}
	//if info.PlayerId != "" {
	//	db = db.Where("player_id LIKE ?", "%"+info.PlayerId+"%")
	//}
	//if info.ForbiddenLogin != nil {
	//	db = db.Where("forbidden_login = ?", info.ForbiddenLogin)
	//}
	//if info.ForbiddenInGameHeroExport != nil {
	//	db = db.Where("forbidden_in_game_hero_export = ?", info.ForbiddenInGameHeroExport)
	//}
	//err = db.Count(&total).Error
	//if err != nil {
	//	return
	//}
	//
	//if limit != 0 {
	//	db = db.Limit(limit).Offset(offset)
	//}
	//
	//err = db.Find(&adminModuleControls).Error
	//return adminModuleControls, total, err
	return nil, 0, nil
}
