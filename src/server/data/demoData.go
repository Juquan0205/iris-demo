package data

import (
	"maizuo.com/back-end/iris-demo/src/server/util"
	"fmt"
)

type DemoData struct {

}

func (*DemoData) getDataDemo(id int) error{
	url := "http://active.maizuo.com"
	params := map[string]interface{}{}
	params["id"] = id
	//在这里调用第三方的接口
	result, err := util.HttpGet(url, params)
	if err != nil {
		return err
	}
	fmt.Println("返回的结果是: ", result)
	return nil
}
