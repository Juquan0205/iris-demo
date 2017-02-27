package main

import (
	"maizuo.com/back-end/iris-demo/src/server/initialize"
	"maizuo.com/back-end/iris-demo/src/server/timer"
)

func main() {

	initialize.SetupConfig()

	initialize.SetErrorDeal()

	initialize.SetUpLogger()

	initialize.SetContext()

	initialize.SetUpRedis()

	//initialize.SetupRPC()
	//定时器
	timer.SetupTimer()

	initialize.SetupServer()
}
