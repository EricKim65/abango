package abango

import (
	"encoding/json"
	"os"
	"strings"
	time "time"

	cf "github.com/EricKim65/abango/config"
	e "github.com/EricKim65/abango/etc"

	g "github.com/EricKim65/abango/global"
	gr "github.com/EricKim65/abango/grpc"
	kf "github.com/EricKim65/abango/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	XEnv *EnvConf     //Kangan only
	XDB  *xorm.Engine //Kangan only

	XConfig map[string]string //Kangan only
)

// type Controller struct {
// }

type EnvConf struct { //Kangan only
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
	// e.OkLog("Abango Initialized")
}

func RunServicePoint(params ...string) {
	if err := cf.GetXConfig(); err == nil {
		if g.XConfig["ApiType"] == "Kafka" {
			kf.KafkaServiveStandBy()
		} else if g.XConfig["ApiType"] == "gRpc" {
			gr.GrpcServiveStandBy()
		} else {
			e.MyErr("Error running ServicePoint", nil, true)
		}

		// 	if g.XConfig["KafkaUse"] == "Yes" {
		// 		kf.KafkaServiveStandBy()
		// 	}
		// if g.XConfig["gRpcUse"] == "Yes" {
		// 	gr.GrpcServiveStandBy()
		// }

	} else {
		e.Atp("Error running RunServicePoint")
	}
	// e.Atp(g.XConfig["Dummy"])
}

func RunEndRequest(params ...string) {
	if err := cf.GetXConfig(); err == nil {
		if g.XConfig["ApiType"] == "Kafka" {
			RunRequest(kf.KafkaSyncRequest)
			// } else if g.XConfig["ApiType"] == "gRpc" {
			// 	RunRequest(gr.GrpcRequest)
		} else {
			e.Atp("Error running RunEndPoint")
		}
	} else {

	}
	// e.Atp(g.XConfig["Dummy"])
}

func RunRequest(MsgHandler func(string, string) (string, string, error)) error {

	var a g.AbangoAsk

	unique_id := e.RandString(20)

	askfile := e.GetAskName()
	arrask := strings.Split(askfile, "@") // @앞의 문자를 askname으로 설정
	askname := arrask[0]
	jsonsend := g.XConfig["JsonSendDir"] + askname + ".json"
	jsonreceive := g.XConfig["JsonReceiveDir"] + askname + ".json"

	// kk := []g.ComVar{}
	jsonsvrvars := "conf/server-vars.json"
	if file, err := os.Open(jsonsvrvars); err == nil {
		decoder := json.NewDecoder(file)
		if err = decoder.Decode(&a.ServerVars); err == nil {

			if askstr, err := e.FileToStr(jsonsend); err == nil {

				a.AskName = askname
				a.UniqueId = unique_id
				a.Body = []byte(askstr)

				askstr, _ := json.Marshal(&a)
				if retstr, retsta, err := MsgHandler(string(askstr), unique_id); err == nil {
					e.Tp("ReturnStatus: " + retsta + "  ReturnJsonFile: " + jsonreceive)
					e.StrToFile(jsonreceive, retstr)
					if g.XConfig["ShowReceivedJson"] == "Yes" {
						e.Tp(retstr)
					}
				} else {
					e.MyErr("QWERDSFAERQRDA-MsgHandler", err, true)
				}
			} else {
				e.MyErr("WERZDSVCZSRE-JsonSendFile", err, true)
			}
		} else {
			return e.MyErr("LAAFDFERHYWE", err, true)
		}
	} else {
		return e.MyErr("LAAFDFDWDERHYWE", err, true)
	}

	// if fvar, err := cf.GetServerVarsInEnd(askname); err == nil {
	// 	if askstr, err := e.FileToStr(jsonsend); err == nil {

	// 		a.AskName = askname
	// 		a.UniqueId = unique_id
	// 		a.Ask.Body = []byte(askstr)
	// 		a.Ask.ServerVars =

	// 		delim := g.XConfig["MsgDelimiter"]
	// 		combined_msg := askname + delim + fvar + delim + askstr
	// 		// e.Tp(combined_msg)
	// 		if retstr, retsta, err := MsgHandler(combined_msg); err == nil {
	// 			e.Tp("ReturnStatus: " + retsta + "  ReturnJsonFile: " + jsonreceive)
	// 			e.StrToFile(jsonreceive, retstr)
	// 			if g.XConfig["ShowReceivedJson"] == "Yes" {
	// 				e.Tp(retstr)
	// 			}
	// 		} else {
	// 			e.MyErr("QWERDSFAERQRDA-MsgHandler", err, true)
	// 		}
	// 	} else {
	// 		e.MyErr("WERZDSVCZSRE-JsonSendFile", err, true)
	// 	}
	// } else {
	// 	e.MyErr("QERVZDBXTRFG-MsgHandler", err, true)
	// }
	return nil
}

func Run(params ...string) { //Kangan only

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

func GetEnvConf() error { // Kangan only

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
