package tgbot

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// 启动 TgBot 服务
func StartTgBot() error {
	// 从配置文件获取机器人 token
	token := viper.GetString("TgBot.Token")
	if token == "" {
		return fmt.Errorf("TgBot Token 为空！")
	}

	// 通过从配置文件获取到的 token ，创建机器人Client
	b, err := gotgbot.NewBot(token, &gotgbot.BotOpts{
		Client: http.Client{
			Transport: &http.Transport{
				// 设置代理，从环境变量中获取
				Proxy: http.ProxyFromEnvironment,
			},
		},
		DefaultRequestOpts: &gotgbot.RequestOpts{
			Timeout: gotgbot.DefaultTimeout,
			APIURL:  gotgbot.DefaultAPIURL,
		},
	})
	if err != nil {
		return err
	}

	// 创建更新程序和调度程序。
	updater := ext.NewUpdater(&ext.UpdaterOpts{
		Dispatcher: ext.NewDispatcher(&ext.DispatcherOpts{
			// If an error is returned by a handler, log it and continue going.
			Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
				log.Error("处理更新时出错:", err.Error())
				return ext.DispatcherActionNoop
			},
			MaxRoutines: ext.DefaultMaxRoutines,
		}),
	})
	dispatcher := updater.Dispatcher

	// 绑定 /start 命令的处理程序
	dispatcher.AddHandler(handlers.NewCommand("start", start))

	dispatcher.AddHandler(handlers.NewCommand("help", help))
	// 绑定点击按钮后的应答程序
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.Equal("start_callback"), startCB))

	// 开始接收更新
	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		return err
	}
	log.Printf("[%s] TgBot 已启动", b.User.Username)

	// 阻塞下线程，避免线程退出
	updater.Idle()
	return nil
}
