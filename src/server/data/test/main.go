package test

import (
	"fmt"
	"github.com/spf13/viper"
	"maizuo.com/back-end/iris-demo/src/server/initialize"
)

func init() {
	conf := "./../../../../config/local"
	viper.SetConfigName(conf)
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("err:", err)
	}
	initialize.SetErrorDeal()

	initialize.SetUpLogger()

	initialize.SetContext()

	initialize.SetupRPC()
	//准备要测试的参数
}