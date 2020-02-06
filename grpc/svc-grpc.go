package grpc

import (
	// "grpc-kangan/routers"
	"bytes"
	"context"
	"net"

	cf "github.com/EricKim65/abango/config"
	e "github.com/EricKim65/abango/etc"
	g "github.com/EricKim65/abango/global"
	grp1 "github.com/EricKim65/abango/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

func GrpcServiveStandBy() {

	lis, err := net.Listen(g.XConfig["gRpcProtocol"], g.XConfig["gRpcAddr"]+":"+g.XConfig["gRpcPort"])
	if err != nil {
		e.MyErr("QERQDFGVX", err, true)
		// log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	grp1.RegisterGrp1Server(grpcServer, NewGrp1Server())
	e.OkLog("gRpc->" + g.XConfig["gRpcProtocol"] + ":" + g.XConfig["gRpcAddr"] + ":" + g.XConfig["gRpcPort"] + " is starting !")

	grpcServer.Serve(lis)

}

type grp1Server struct {
}

func NewGrp1Server() *grp1Server {
	return &grp1Server{}
}

func (s *grp1Server) StdRpc(c context.Context, ask *grp1.StdAsk) (*grp1.StdRet, error) {

	p, _ := peer.FromContext(c)
	ip_no := p.Addr.String()
	// arr_ip := strings.Split(ip_no, ":")

	// arrask := strings.Split(ask.AskMsg, g.XConfig["MsgDelimiter"])
	arrask := bytes.Split(ask.AskMsg, []byte(g.XConfig["MsgDelimiter"]))
	if len(arrask) != 3 {
		return nil, e.MyErr("Count of varialbles in askstr is NOT 3", nil, true)
	}

	AskName := string(arrask[0])
	e.Tp("gRpc request arrived [" + AskName + "] from " + ip_no)
	if err := cf.GetServerVarsInSvc(arrask[1]); err == nil {
	} else {

		e.MyErr("QEWRQCVZVCXER-cf.ServerVars-", err, false)
	}

	var retstr, retsta []byte

	aname := "sss"
	if aname == "setup-" {

	} else { // 처리없이 자료를 되돌려줌
		retsta = []byte("200")
		retstr = arrask[2]
	}

	ret := grp1.StdRet{
		RetSta: []byte(retsta),
		RetMsg: []byte(retstr),
	}

	return &ret, nil

}
