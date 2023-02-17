package controller

import (
	"encoding/json"

	"github.com/kataras/iris/v12"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type errorResultJson struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
}

type successResultJson struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
}

func AdminLogin(ctx iris.Context) {
	log.Infof("有管理员尝试登录 用户名：%s 密码：%s", ctx.URLParamDefault("username", ""), ctx.URLParamDefault("password", ""))
	session := Sess.Start(ctx)
	// 检查管理员账号密码是否为空
	if ctx.URLParamDefault("username", "") == "" || ctx.URLParamDefault("password", "") == "" {
		session.Set("authenticated", false)
		MessageJson, err := json.Marshal(errorResultJson{
			false,
			"管理员账号密码为空！",
		})
		if err != nil {
			log.Error("json编码失败：", err)
			return
		}
		ctx.WriteString(string(MessageJson))
		return
	}

	// 检查管理员账号密码是否正确
	if ctx.URLParamDefault("username", "") != viper.GetString("AdminUser.username") || ctx.URLParamDefault("password", "") != viper.GetString("AdminUser.password") {
		session.Set("authenticated", false)
		MessageJson, err := json.Marshal(errorResultJson{
			false,
			"管理员账号密码错误！",
		})
		if err != nil {
			log.Error("json编码失败：", err)
			return
		}
		ctx.WriteString(string(MessageJson))
		return
	}

	session.Set("authenticated", true)
	MessageJson, err := json.Marshal(successResultJson{
		true,
		"管理员登录成功！",
	})
	if err != nil {
		log.Error("json编码失败：", err)
	}
	ctx.WriteString(string(MessageJson))
}
