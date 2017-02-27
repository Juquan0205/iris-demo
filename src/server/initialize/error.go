package initialize

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/go-errors/errors"
	"maizuo.com/back-end/iris-demo/src/server/errcode"
	"maizuo.com/back-end/iris-demo/src/server/util"
)

func SetErrorDeal() {
	iris.Use(iris.HandlerFunc(func(ctx *iris.Context) {
		defer func() {
			if err := recover(); err != nil {
				msg := fmt.Sprintf("发生panic异常: %v\n", errors.Wrap(err, 2).ErrorStack())
				util.Log(ctx, errcode.SYSTEM_ERROR.ResultWithMsg(msg))
				ctx.JSON(iris.StatusOK, errcode.SYSTEM_ERROR.Result())
			}
		}()
		ctx.Next()
	}))

}
