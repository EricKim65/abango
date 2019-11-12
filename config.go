package abango

import (
	"encoding/json"
	"os"

	"github.com/go-xorm/xorm"
)

var (
	XEnv *EnvConf
	XDB  *xorm.Engine
)

type RunConf struct {
	RunMode     string
	DevPrefix   string
	ProdPrefix  string
	ConfPostFix string
}

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

func GetEnvConf() error {

	conf := "conf/"
	RunFilename := conf + "run_conf.json"

	var run RunConf

	if file, err := os.Open(RunFilename); err != nil {
		MyErr("SDFLJDSAFJA", nil, true)
		return err
	} else {
		decoder := json.NewDecoder(file)
		if err = decoder.Decode(&run); err != nil {
			MyErr("LASJLDFJASFJ", err, true)
			return err
		}
	}

	filename := conf + run.RunMode + run.ConfPostFix
	if file, err := os.Open(filename); err != nil {
		MyErr("QERTRRTRRW", err, true)
		return err
	} else {
		decoder := json.NewDecoder(file)
		if err = decoder.Decode(&XEnv); err != nil {
			MyErr("LAAFDFERHY", err, true)
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
