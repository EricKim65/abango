package abango

import (
	"encoding/json"
	"os"
	time "time"

	e "github.com/EricKim65/abango/etc"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	XEnv *EnvConf
	XDB  *xorm.Engine
)

type EnvConf struct {
	AppName      string
	HttpProtocol string
	HttpAddr     string
	HttpPort     string
	SiteName     string

	DbType     string
	DbHost     string
	DbUser     string
	DbPassword string
	DbPort     string
	DbName     string
	DbPrefix   string
	DbTimezone string

	DbStr string
}

type RunConf struct {
	RunMode     string
	DevPrefix   string
	ProdPrefix  string
	ConfPostFix string
}

func init() {
	e.OkLog("Abango Initialized")
}

func Run(params ...string) {

	if err := GetEnvConf(); err == nil {
		db, err := xorm.NewEngine(XEnv.DbType, XEnv.DbStr)
		// db, err := xorm.NewEngine(XEnv.DbType, "root:root@tcp(127.0.0.1:3306)/kangan?charset=utf8&parseTime=True")

		dbStr := XEnv.DbType + ":(" + XEnv.DbHost + ":" + XEnv.DbPort + ")->[" + XEnv.DbPrefix + XEnv.DbName + "] DB Schema"
		if err == nil {
			e.OkLog(dbStr)
		} else {
			e.MyErr(dbStr, err, true)
			// panic(fmt.Errorf("Database open error: %s \n", err))
		}

		db.ShowSQL(false)
		db.SetMaxOpenConns(100)
		db.SetMaxIdleConns(20)
		db.SetConnMaxLifetime(60 * time.Second)
		if _, err := db.IsTableExist("aaa"); err != nil { //Connect Check
			e.MyErr("DATABASE DISCONNECTED", err, true)
		} else {
			e.OkLog("DATABASE CONNECTED")
		}
		XDB = db
	}

}

func GetEnvConf() error {

	conf := "conf/"
	RunFilename := conf + "run_conf.json"

	var run RunConf

	if file, err := os.Open(RunFilename); err != nil {
		e.MyErr("SDFLJDSAFJA", nil, true)
		return err
	} else {
		decoder := json.NewDecoder(file)
		if err = decoder.Decode(&run); err != nil {
			e.MyErr("LASJLDFJASFJ", err, true)
			return err
		}
	}

	filename := conf + run.RunMode + run.ConfPostFix
	if file, err := os.Open(filename); err != nil {
		e.MyErr("QERTRRTRRW", err, true)
		return err
	} else {
		decoder := json.NewDecoder(file)
		if err = decoder.Decode(&XEnv); err != nil {
			e.MyErr("LAAFDFERHY", err, true)
			return err
		}
	}

	if XEnv.DbType == "mysql" {
		XEnv.DbStr = XEnv.DbUser + ":" + XEnv.DbPassword + "@tcp(" + XEnv.DbHost + ":" + XEnv.DbPort + ")/" + XEnv.DbPrefix + XEnv.DbName + "?charset=utf8"
	} else if XEnv.DbType == "mssql" {
		// Add on more DbStr of Db types
	}

	return nil
}
