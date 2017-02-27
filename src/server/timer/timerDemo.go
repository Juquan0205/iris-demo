package timer

import (
	"time"
	"fmt"
)


//订单表查询开始的id,默认从0开始
var nextId uint64

func TimerDemo() {
	fmt.Println("==============模板定时器启动============nextId=",nextId,"==========",time.Now())

	//执行定时器任务


}
