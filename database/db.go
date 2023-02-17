package database

import (
	"fmt"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ExportContext struct {
	// 等待 WaitGroup
	LoginWaitGroup sync.WaitGroup
	// 数据库连接
	DB *gorm.DB
}

// prepareDB 准备数据库连接
func prepareDB() *gorm.DB {
	dbUser := viper.GetString("DB.User")
	dbPassword := viper.GetString("DB.Password")
	dbHost := viper.GetString("DB.Host")
	dbPort := viper.GetString("DB.Port")
	dbName := viper.GetString("DB.Name")
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.WithError(err).Fatal("连接数据库失败")
		// 然后就直接退出了
		os.Exit(1)
	}
	return db
}

// AsyncUpdate 异步将日志写入数据库
func (c *ExportContext) AsyncUpdate(data string) {
	// 开一个 Context
	//dCtx, cancel := context.WithCancel(context.Background())
	//dCtx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	//defer cancel()
	// 拉出我们的客户端
	/*client := sCtx.Client
	gaps := sCtx.UpdateManager
	err := client.Run(
		context.Background(), func(ctx context.Context) error {
			// 拉取用户信息
			user, err := client.Self(ctx)
			//client.Config().GetThisDC()
			//client.Config().GetDCOptions()[0].
			if err != nil {
				log.WithError(err).Error("拉取用户信息失败")
				return err
			}
			// 通知更新管理器让他更新
			if err := gaps.Auth(ctx, client.API(), user.ID, user.Bot, true); err != nil {
				log.WithError(err).Error("通知管理器更新失败")
				return err
			}
			defer func() { _ = gaps.Logout() }()

			<-ctx.Done()
			return ctx.Err()
		},
	)
	if err != nil {
		log.WithError(err).Error("运行客户端失败")
		return
	}*/
}
