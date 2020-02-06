package grpc

import (
	"context"
	"time"

	grp1 "github.com/EricKim65/abango/protos"

	e "github.com/EricKim65/abango/etc"
	g "github.com/EricKim65/abango/global"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func GrpcRequest(askstr string, dummy string) (string, string, error) {

	dial, err := grpc.Dial(
		g.XConfig["gRpcAddr"]+":"+g.XConfig["gRpcPort"],
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                time.Millisecond,
			Timeout:             time.Millisecond,
			PermitWithoutStream: true,
		}),
	)

	if err != nil {
		e.MyErr("WOEIURQPWERQ", nil, true)
	}
	defer dial.Close()
	c := grp1.NewGrp1Client(dial)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// e.Tp(askstr)
	if r, err := c.StdRpc(ctx, &grp1.StdAsk{AskMsg: []byte(askstr)}); err == nil {
		return string(r.RetMsg), string(r.RetSta), nil
	} else {
		return "", "", e.MyErr("WERVCZWERTGS", err, true)
	}

}
