package initialize

import (
	"github.com/kataras/iris"
	"github.com/spf13/viper"
	"maizuo.com/back-end/iris-demo/src/server/route/api"
)

func SetupServer() {

	port := viper.GetString("server.port")
	host := viper.GetString("server.host")

	api.Api()

	iris.Listen(host + ":" + port)

}
