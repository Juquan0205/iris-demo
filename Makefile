
# 常量定义
MAIN_GO = "./main.go"
DEPLOY_GO = "./build/main.go"
CONFIG_LOCAL = "./config/local"
CONFIG_DEV = "./config/dev"
CONFIG_PROD = "./config/prod"
APIDOC_DIR = "./apidoc/"



# 定义命令包
define deploy_dev
go run ${DEPLOY_GO} -conf ${CONFIG_DEV} -flow $(flow)
endef

# 定义命令包
define deploy_prod
go run ${DEPLOY_GO} -conf ${CONFIG_PROD} -flow $(flow)
endef

define build-linux
GOOS=linux GOARCH=amd64 go build -o ${BUILD_DIR} ${MAIN_GO}
#GOOS=${OS} GOARCH=${ARCH} go build -o ${DEST}
endef

define go
go run $(MAIN_GO) -conf $(CONFIG_LOCAL)
endef

define doc
apidoc -i ./src/server/ -o ${APIDOC_DIR}
endef



# 项目初始化
.PHONY : init
init:
	@echo 正在创建项目日志目录...
#	mkdir -p log
#	mkdir -p log/pc
	@echo 正在创建项目日志文件...
#	touch ${LOGFILE}
#	touch ${BUSSINESSLOGFILE}

# 本地运行项目
.PHONY : go
go: init
	@echo 正在启动项目...
	$(go)

# 生成项目apidoc文档
.PHONY : doc
doc:
	@echo 正在生成项目apidoc文档...
	$(doc)
	@echo 生成完成!

