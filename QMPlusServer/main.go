package main

import (
	"net/http"
	"time"

	"gin-vue-admin/QMPlusServer/config"
	initrouter "gin-vue-admin/QMPlusServer/init/initRouter"
	"gin-vue-admin/QMPlusServer/init/qmlog"
	"gin-vue-admin/QMPlusServer/init/qmsql"
	registtable "gin-vue-admin/QMPlusServer/init/regist-table"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	qmlog.InitLog()
	registtable.RegistTable(qmsql.InitMysql(config.DBConfig.Admin))
	defer qmsql.DEFAULTDB.Close()
	r := initrouter.InitRouter()
	// qmlog.QMLog.Info("服务器开启") // 日志测试代码
	s := &http.Server{
		Addr:           "127.0.0.1:8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}
