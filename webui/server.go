package webui

import (
	"CandlestickGame/webui/controller"
	"io/fs"
	"net/http"

	"github.com/kataras/iris/v12"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func StartWebServer() error {
	ListenAddress := viper.GetString("Server.ListenAddress")

	htmlView, _ := fs.Sub(view, "view")
	app := iris.New()
	app.Use(iris.Compression)
	app.Handle(iris.MethodGet, "/api/login", controller.AdminLogin)

	app.Handle(iris.MethodGet, "/{index:path}", iris.FileServer(http.FS(htmlView), iris.DirOptions{Compress: true, ShowList: true, IndexName: "index.html"}))

	if viper.GetBool("Server.TLS.Enable") != true {
		err := app.Run(iris.Addr(ListenAddress))
		if err != nil {
			return err
		}
	}
	if viper.GetString("Server.TLS.Cert.CertFile") != "" && viper.GetString("Server.TLS.Cert.KeyFile") != "" {
		log.Infof(viper.GetString("Server.TLS.Cert.CertFile"))
		log.Infof(viper.GetString("Server.TLS.Cert.KeyFile"))
		err := app.Run(iris.TLS(ListenAddress, viper.GetString("Server.TLS.Cert.CertFile"), viper.GetString("Server.TLS.Cert.KeyFile")))
		if err != nil {
			return err
		}
	}
	err := app.Run(iris.AutoTLS(ListenAddress, viper.GetString("Server.Domain"), "admin@example.com"))
	if err != nil {
		return err
	}
	return nil
}
