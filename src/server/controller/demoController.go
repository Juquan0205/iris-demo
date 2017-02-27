package controller

import (
	"github.com/kataras/iris"
	"maizuo.com/back-end/iris-demo/src/server/entity"
	"maizuo.com/back-end/iris-demo/src/server/errcode"
	"maizuo.com/back-end/iris-demo/src/server/service"
	"maizuo.com/back-end/iris-demo/src/server/service/serviceInter"
)

type DemoController struct {
}

/**
实例化该controller用到的对象
*/
func (*DemoController) HelloWorld(ctx *iris.Context) {
	demoService := &service.DemoService{}
	doHelloWorld(ctx, demoService)
}

/**
controller层:
只对入参进行处理,不处理业务逻辑
*/
func doHelloWorld(ctx *iris.Context, demoService serviceInter.DemoServiceInterface) {
	name := ctx.GetString("name")
	if name == "" {
		ctx.JSON(iris.StatusOK, errcode.PARAM_ERROR.Result())
	}
	result, err := demoService.DealBusinessLogic(name)
	if err != nil {
		ctx.JSON(iris.StatusOK, &entity.Result{Status: -1, Msg: "系统异常"})
	}
	ctx.JSON(iris.StatusOK, &entity.Result{Status: 0, Msg: "success", Data: result})
}
