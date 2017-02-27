package api

import (
	"github.com/kataras/iris"
	"maizuo.com/back-end/iris-demo/src/server/controller"
)

func Api() {

	var (
		base = &controller.BaseController{}
		demo = &controller.DemoController{}

	)
	//日志记录
	api := iris.Party("/api", base.SetLog)

	api.Get("/demo/:name", demo.HelloWorld)
}
