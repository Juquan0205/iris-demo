# iris-demo
> iris框架demo

## 开发及发布
### 工具
* golang(1.7+)

#### 如何开始
1. 下载并安装[Golang](https://golang.org/dl/)

#### 启动项目
- 启动项目
##### linux或mac系统下:
```
make go
```
##### win系统下如果未安装make
```
go run main.go -conf ./config/local
```

#### 访问demo接口
> http://localhost:8080/api/demo/eno

#### 项目结构说明

```
|____iris-demo						项目根目录
| |____config       				配置文件目录
| | |____dev.json					测试环境配置   
| | |____local.json					本地环境配置
| | |____prod.json					正式环境配置
| |____Godeps						go项目打包描述目录
| | |____Godeps.json
| | |____Readme
| |____Makefile
| |____README.md
| |____src 							代码根目录
| | |____server 					后端项目根目录
| | | |____controller 				controller层 处理业务参数
| | | | |____demoController.go
| | | | |____main.go
| | | |____data 					data层 从其他系统获取数据 包括db,第三方接口等
| | | | |____dataInter 				data层接口
| | | | | |____demoDataInterface.go
| | | | |____demoData.go 			
| | | | |____test 					data层测试
| | | | | |____demoData_test.go
| | | | | |____main.go
| | | |____entity 					entity
| | | | |____result.go
| | | |____errcode 					全局错误集合
| | | | |____main.go
| | | |____initialize 				项目初始化层: 处理项目启动时需要初始化的内容:redis,db,log等
| | | | |____config.go              初始化加载项目配置
| | | | |____context.go             初始化项目上下文环境
| | | | |____error.go               全局panic异常拦截
| | | | |____logger.go              打开日志记录连接
| | | | |____redis.go               建立redis连接
| | | | |____rpc.go                 建立grpc连接
| | | | |____server.go              启动http服务
| | | |____proto 					grpc传输生成的对象
| | | | |____mallorderproto
| | | | | |____order.pb.go
| | | | |____mallproductproto
| | | | | |____product.pb.go
| | | |____route 					route层: 处理路由地址 分派到对应的controller方法
| | | | |____api 					api: 后端接口相关
| | | | | |____main.go
| | | | |____web 					web: 前端页面相关
| | | |____service 					service层: 处理业务逻辑
| | | | |____demoService.go
| | | | |____serviceInter 			service层接口
| | | | | |____demoServiceInterface.go
| | | | |____test 					service层测试
| | | |____test 					
| | | | |____demo_test.go
| | | | |____test_test.go
| | | |____timer 					定时器相关
| | | | |____healthCheck.go
| | | | |____timer.go
| | | | |____timerDemo.go
| | | |____util 					项目工具类
| | | | |____base64.go
| | | | |____http.go
| | | | |____ip.go
| | | | |____log.go
| | | | |____md5.go
| | | | |____rateLimiter.go
| | | | |____redis.go
```

