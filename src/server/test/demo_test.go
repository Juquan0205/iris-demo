package test

import (
	"testing"
	"maizuo.com/back-end/iris-demo/src/server/service"
)

/*
测试查询积分
 */
func TestQueryScore(t *testing.T){
	result, err := service.DemoService.DealBusinessLogic("eno")
	if err != nil {
		t.Error("测试出错 err=", err)
	}
	t.Log("TestQueryScore返回的结果是: ", result)
}
