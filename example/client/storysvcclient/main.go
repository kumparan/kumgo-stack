package main

import (
	"fmt"
	"time"

	"github.com/kumparan/go-lib/logger"
	"github.com/kumparan/kumgo-stack/buff"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func init() {
	logger.SetupLoggerAuto("", "")
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("localhost:9001"), grpc.WithInsecure())
	if err != nil {
		logger.Fatal("did not connect: ", err)
	}
	defer conn.Close()
	c := buff.NewStoryServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetStories(ctx, &buff.Empty{})
	if err != nil {
		logger.Fatal("could not greet: ", err)
	}

	logger.Info("Echo result: ", r.Stories)
}
