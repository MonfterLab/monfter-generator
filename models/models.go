package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	sqlConf, err := beego.AppConfig.GetSection("mysql")
	if err != nil {
		panic(err)
	}
	user := sqlConf["user"]
	pass := sqlConf["pass"]
	host := sqlConf["host"]
	port := sqlConf["port"]
	db := sqlConf["db"]
	beego.BeeLogger.Debug("DB user:%v host:%v:%v db:%v", user, host, port, db)

	dataSource := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4", user, pass, host, port, db)
	logs.Info("central", sqlConf, dataSource)

	err = orm.RegisterDataBase("default", "mysql", dataSource)
	if err != nil {
		panic(err)
	}
	orm.DefaultTimeLoc = time.UTC
	orm.SetMaxOpenConns("default", 600)
	orm.SetMaxIdleConns("default", 300)
	orm.Debug = true
}
