package abango

import (
	"fmt"
	time "time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func init() {
	OkLog("Abango Initialized")
}

func Run(params ...string) {

	if err := GetEnvConf(); err == nil {

		db, err := xorm.NewEngine(XEnv.DbType, XEnv.DbStr)
		// db, err := xorm.NewEngine(XEnv.DbType, "root:root@tcp(127.0.0.1:3306)/kangan?charset=utf8&parseTime=True")

		if err == nil {
			OkLog(XEnv.DbType + ":(" + XEnv.DbHost + ":" + XEnv.DbPort + ")->[" + XEnv.DbPrefix + XEnv.DbName + "] DB Schema ")
		} else {
			panic(fmt.Errorf("Database open error: %s \n", err))
		}

		db.ShowSQL(false)
		db.SetMaxOpenConns(100)
		db.SetMaxIdleConns(20)
		db.SetConnMaxLifetime(60 * time.Second)
		if _, err := db.IsTableExist("aaa"); err != nil { //Connect Check
			MyErr("DATABASE DISCONNECTED", err, true)
		} else {
			OkLog("DATABASE CONNECTED")
		}
		XDB = db
	}

}
