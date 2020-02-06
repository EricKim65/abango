package config

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	e "github.com/EricKim65/abango/etc"
	g "github.com/EricKim65/abango/global"
	_ "github.com/go-sql-driver/mysql"
)

func GetXConfig() error { // Kafka, gRpc, REST 통합 업그레이드

	conf := "conf/"
	RunFilename := conf + "config_select.json"

	run := struct {
		ConfSelect  string
		ConfPostFix string
	}{}

	if file, err := os.Open(RunFilename); err != nil {
		e.MyErr("WERQRRQERQWERFD", nil, true)
		return err
	} else {
		decoder := json.NewDecoder(file)
		if err = decoder.Decode(&run); err != nil {
			e.MyErr("ERTFDFDAFA", err, true)
			return err
		}
	}

	g.XConfig = make(map[string]string) // Just like malloc
	config := []g.ComVar{}

	// var varMap []map[string]interface{}
	filename := conf + run.ConfSelect + run.ConfPostFix
	if file, err := os.Open(filename); err != nil {
		e.MyErr("QERTRRTRRW", err, true)
		return err
	} else {
		decoder := json.NewDecoder(file)
		if err = decoder.Decode(&config); err == nil {
			for _, p := range config {
				g.XConfig[p.Key] = p.Value
			}
		} else {
			e.MyErr("LAAFDFERHWERYTY", err, true)
			return err
		}
	}

	if g.XConfig["ApiType"] == "Kafka" {
		e.Tp("=====" + "Config file prefix: " + run.ConfSelect + "=====" + g.XConfig["ApiType"] + " Connection: " + g.XConfig["KafkaAddr"] + ":" + g.XConfig["KafkaPort"] + "=====")
	} else if g.XConfig["ApiType"] == "gRpc" {
		e.Tp("=====" + "Config file prefix: " + run.ConfSelect + "=====" + g.XConfig["ApiType"] + " Connection: " + g.XConfig["gRpcAddr"] + ":" + g.XConfig["gRpcPort"] + "=====")
	}
	// e.Tp(g.XConfig["AppName"])

	return nil
}

func GetServerVarsInEnd(askname string, unique_id string) (string, error) { // Kafka, gRpc, REST 통합 업그레이드

	// unique_id := e.RandString(20)
	fvars := []g.ComVar{}
	// comarr := make(map[string]string)

	filename := "conf/server-vars.json"
	if file, err := os.Open(filename); err != nil {
		e.MyErr("QERTRRTRRWQWRE", err, true)
		return "", err
	} else {
		decoder := json.NewDecoder(file)
		if err = decoder.Decode(&fvars); err == nil {
			for i := 0; i < len(fvars); i++ {
				if fvars[i].Key == "askname" {
					fvars[i].Value = askname // 유일키
				} else if fvars[i].Key == "unique_id" {
					fvars[i].Value = unique_id
				} else if fvars[i].Key == "server_addr" {
					addrs, _ := net.InterfaceAddrs()
					fvars[i].Value = fmt.Sprintf("%v", addrs[0]) // Server IP
				}
			}
		} else {
			e.MyErr("LAAFDFERHYWE", err, true)
			return "", err
		}
	}

	fstr, _ := json.Marshal(&fvars)
	return string(fstr), nil
}

func GetServerVarsInSvc(t []byte) error { // Kafka, gRpc, REST 통합 업그레이드

	g.ServerVars = make(map[string]string) // 반드시 = 로 할 것
	evars := []g.ComVar{}

	if err := json.Unmarshal(t, &evars); err == nil {
		for _, p := range evars {
			g.ServerVars[p.Key] = p.Value
		}
	} else {
		e.MyErr("QWECVZDFVBXGF", err, true)
		return err
	}

	return nil
}

// func GetMapVars(t []ComVar) (map[string]string, error) { // Kafka, gRpc, REST 통합 업그레이드

// 	comarr := make(map[string]string)
// 	if content, err := ioutil.ReadFile("golangcode.txt"); err == nil {
// 		if err := json.Unmarshal(content, &t); err == nil {
// 			for _, p := range t {
// 				comarr[p.Name] = p.Value
// 			}
// 		} else {
// 			e.MyErr("QWECVZDFVBXGF", err, true)
// 			return nil, err
// 		}
// 	} else {
// 		e.MyErr("QERTRRTRRW", err, true)
// 		return nil, err
// 	}

// 	return comarr, nil
// }
