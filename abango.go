package abango

import (
	"fmt"
	time "time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func init() {
	okLog("Abango Initialized")
}

func Run(params ...string) {

	if err := GetEnvConf(); err == nil {

		db, err := xorm.NewEngine(XEnv.DbType, XEnv.DbStr)
		// db, err := xorm.NewEngine(XEnv.DbType, "root:root@tcp(127.0.0.1:3306)/kangan?charset=utf8&parseTime=True")

		if err == nil {
			okLog(XEnv.DbType + ":(" + XEnv.DbHost + ":" + XEnv.DbPort + ")->" + XEnv.DbPrefix + XEnv.DbName + " DB connected !")
		} else {
			panic(fmt.Errorf("Database open error: %s \n", err))
		}

		db.ShowSQL(false)
		db.SetMaxOpenConns(100)
		db.SetMaxIdleConns(20)
		db.SetConnMaxLifetime(60 * time.Second)
		XDb = db
	}

}
