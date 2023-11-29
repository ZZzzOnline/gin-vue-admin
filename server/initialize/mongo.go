package initialize

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var MongoOfficial = new(mongoOfficial)

type mongoOfficial struct{}

func (m *mongoOfficial) Indexes(ctx context.Context) error {

	err := service.ServiceGroupApp.GmServiceGroup.AdminModuleControlService.InitAdminModuleControlIndexes()
	if err != nil {
		return err
	}

	return nil
}

func (m *mongoOfficial) Initialization() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(global.GVA_CONFIG.Mongo.Uri()))
	if err != nil {
		return errors.Wrap(err, "链接mongodb数据库失败!")
	}

	// 设置连接池大小
	clientOptions := options.Client().
		SetMinPoolSize(global.GVA_CONFIG.Mongo.MinPoolSize).
		SetMaxPoolSize(global.GVA_CONFIG.Mongo.MaxPoolSize)

	// 设置超时时间
	clientOptions = clientOptions.
		SetSocketTimeout(time.Duration(global.GVA_CONFIG.Mongo.SocketTimeoutMs) * time.Millisecond).
		SetConnectTimeout(time.Duration(global.GVA_CONFIG.Mongo.ConnectTimeoutMs) * time.Millisecond)

	//// 设置认证
	//credential := options.Credential{
	//	Username: global.GVA_CONFIG.Mongo.Username,
	//	Password: global.GVA_CONFIG.Mongo.Password,
	//}
	//clientOptions = clientOptions.SetAuth(credential)

	//连接
	client, err = mongo.Connect(ctx, clientOptions.ApplyURI(global.GVA_CONFIG.Mongo.Uri()))
	if err != nil {
		return errors.Wrap(err, "链接mongodb数据库失败!")
	}

	global.GVA_MONGO = client
	global.GVA_MONGODB = client.Database(global.GVA_CONFIG.Mongo.Database)

	err = m.Indexes(ctx)
	if err != nil {
		return err
	}

	return nil
}
