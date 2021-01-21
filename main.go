package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv24/global"
	"github.com/liuhongdi/digv24/router"
	"log"
)

//init
func init() {
	//setting
	err := global.SetupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	//logger
	err = global.SetupLogger()
	if err != nil {
		log.Fatalf("init.SetupLogger err: %v", err)
	}

	//access logger
	err = global.SetupAccessLogger()
	if err != nil {
		log.Fatalf("init.SetupAccessLogger err: %v", err)
	}

	global.Logger.Infof("------应用init结束")
	//global.Logger.
}

func main() {
	global.Logger.Infof("------应用main函数开始")
	//设置运行模式
	gin.SetMode(global.ServerSetting.RunMode)
	//引入路由
	r := router.Router()
	//run
	r.Run(":"+global.ServerSetting.HttpPort)
}




