package timer

import (
	"maizuo.com/back-end/iris-demo/src/server/util"
)

func HealthCheck () {
	util.AlarmLog("back-end_demo_HealthCheck","system=back-end_demo","系统健康检查")
}
