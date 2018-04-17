package hellosvc

/*

IMPORTANT!
This is only example. This example assumes that there is another service out there called ms-hello, which repo
is at github.com/kumparan/ms-hello/buff.

This package later can be used by another package.

============================================================================================================================

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	buff "github.com/kumparan/ms-hello/buff"
	"github.com/kumparan/kumgo-stack/config"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

)

var (
	Client buff.HelloServiceClient
)

func init() {
	config.GetConf()
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func New() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(config.HelloServiceTarget()), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	Client = buff.NewHelloServiceClient(conn)
}

func (c *Client) Hello(s string) string {
	r, err := c.Hello(context.Context, &buff.Hello{Message: s})
	if err != nil {
		log.Fatal("could not say hello: %v", err)
	}
	log.Info("Greeting: ", r.Message)
	return ""
}

*/
