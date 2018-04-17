package main

import (
	"fmt"
	"time"

	"github.com/kumparan/go-lib/logger"
	"github.com/kumparan/kumgo-stack/buff"
	"github.com/kumparan/kumgo-stack/config"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func init() {
	config.GetConf()
	logger.SetupLoggerAuto("", "")
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", config.HostName(), config.Port()), grpc.WithInsecure())
	if err != nil {
		logger.Fatal("did not connect: ", err)
	}
	defer conn.Close()
	c := buff.NewEchoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Echo(ctx, &buff.EchoRequest{Message: "i'm screaming"})
	if err != nil {
		logger.Fatal("could not greet: ", err)
	}
	logger.Info("Echo result: ", r.Message)

}
