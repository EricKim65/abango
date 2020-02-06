package grpc

// import (
// 	"context"
// 	"strings"

// 	cf "github.com/EricKim65/abango/config"
// 	grp1 "github.com/EricKim65/abango/protos"

// 	e "github.com/EricKim65/abango/etc"
// 	"google.golang.org/grpc/peer"
// )

// type grp1Server struct {
// }

// func NewGrp1Server() *grp1Server {
// 	return &grp1Server{}
// }

// func (s *grp1Server) StdRpc(c context.Context, ask *grp1.StdAsk) (*grp1.StdRet, error) {

// 	p, _ := peer.FromContext(c)
// 	ip_no := p.Addr.String()
// 	// arr_ip := strings.Split(ip_no, ":")

// 	// if arr_ip[0] == "115.91.250.74" || arr_ip[0] == "115.91.250.77" || arr_ip[0] == "211.178.235.179" {
// 	// 	ret := grp1.StdRet{
// 	// 		RetSta: "200",
// 	// 		RetMsg: "invalid-ip",
// 	// 	}
// 	// 	return &ret, nil
// 	// }

// 	arrask := strings.Split(ask.AskMsg, cf.XConfig["MsgDelimiter"])
// 	if len(arrask) != 3 {
// 		return nil, e.MyErr("Count of varialbles in askstr is NOT 3", nil, true)
// 	}
// 	AskName := arrask[0]
// 	e.Tp("gRpc request arrived [" + AskName + "] from " + ip_no)

// 	retsta := ""
// 	retstr := ""

// 	aname := "sss"
// 	if aname == "setup-" {

// 	} else { // 처리없이 자료를 되돌려줌
// 		retsta = "200"
// 		retstr = arrask[2]
// 	}

// 	ret := grp1.StdRet{
// 		RetSta: retsta,
// 		RetMsg: retstr,
// 	}

// 	return &ret, nil
// }
