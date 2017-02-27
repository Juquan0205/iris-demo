package controller

import (
	"encoding/json"
	"github.com/kataras/iris"
	"maizuo.com/back-end/iris-demo/src/server/util"
	"maizuo.com/back-end/iris-demo/src/server/entity"
)
/**
基础controller
 */
type BaseController struct {
}
/**
日志记录
 */
func (*BaseController) SetLog(ctx *iris.Context) {
	ctx.Next()
	//接口返回自动写入日志
	var result entity.Result
	json.Unmarshal(ctx.Response.Body(), &result)
	util.Log(ctx, &result)
}
