package main

import (
	"CandlestickGame/config"
	"CandlestickGame/tgbot"
	"CandlestickGame/webui"
	"bytes"
	"io/fs"
	"os"

	"github.com/go-faster/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	go func() {
		log.Infof("开始使用 token [%s] 启动TgBot", viper.GetString("TgBot.Token"))
		err := tgbot.StartTgBot()
		if err != nil {
			log.Fatalf("无法启动TgBot [%s]", err.Error())
		}
	}()
	go func() {
		log.Infof("开始在 [%s] 上启动服务端", viper.GetString("Server.ListenAddress"))
		err := webui.StartWebServer()
		if err != nil {
			log.Fatalf("无法启动服务 [%s]", err.Error())
		}
	}()
	select {}
}

func init() {
	var cfgFilePath = "config.yaml"
	viper.SetConfigFile(cfgFilePath)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if errors.As(err, &fs.ErrNotExist) {
			log.Warn("未找到配置文件，使用默认配置")
			// 写入 DefaultConfig
			err = os.WriteFile(cfgFilePath, []byte(config.DefaultConfig), 0644)
			if err != nil {
				log.WithFields(
					log.Fields{
						"err": err,
					},
				).Error("写入默认配置文件失败")
			} else {
				log.WithFields(
					log.Fields{
						"file": cfgFilePath,
					},
				).Info("写入默认配置文件成功")
			}
			// 写完默认配置后再读一次
			err = viper.ReadConfig(bytes.NewBuffer([]byte(config.DefaultConfig)))
			if err != nil {
				log.WithFields(
					log.Fields{
						"err": err,
					},
				).Error("读取默认配置文件失败")
			} else {
				log.Info("已成功读取默认配置")
			}
		} else {
			log.WithFields(
				log.Fields{
					"err": err,
				},
			).Fatal("配置文件解析错误")
		}
	}
	log.WithFields(
		log.Fields{
			"file": viper.ConfigFileUsed(),
		},
	).Info("使用配置文件")

	// 设置日志等级
	setLogLevel()
}

func setLogLevel() {
	if viper.IsSet("LogLevel") {
		logLevel := viper.GetString("LogLevel")
		switch logLevel {
		case "DEBUG":
			log.SetLevel(log.DebugLevel)
		case "INFO":
			log.SetLevel(log.InfoLevel)
		case "WARN":
			log.SetLevel(log.WarnLevel)
		case "ERROR":
			log.SetLevel(log.ErrorLevel)
		case "FATAL":
			log.SetLevel(log.FatalLevel)
		default:
			log.Info("未知的日志等级，使用 INFO 等级")
			log.SetLevel(log.InfoLevel)
		}
		log.WithFields(
			log.Fields{
				"level": logLevel,
			},
		).Info("使用日志等级")
	} else {
		log.SetLevel(log.InfoLevel)
		log.Info("未设置日志等级，使用默认等级")
	}
}
