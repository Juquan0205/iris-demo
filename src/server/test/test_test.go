package test

import (
	"testing"
	"fmt"
)

func Test_Main(t *testing.T) {
	var a = float64(42.8) * 100
	if a == 4280 {
		t.Log("测试ok")
	} else {
		t.Error("测试不通过")
		fmt.Println(a)
	}

}

func TestMain1(t *testing.T) {
	var i float64 = 32.80
	fmt.Println(i * 100)
	t.Log("测试ok")
}

func TestAdd(t *testing.T) {
	if 5==_Add(2,3){
		t.Log("ok")
	}else {
		t.Error("error")
	}
}

func _Add(a, b int) int {
	return a + b
}