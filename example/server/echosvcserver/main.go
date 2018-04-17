package main

import (
	"fmt"
	"net"

	"github.com/kumparan/go-lib/logger"
	"github.com/kumparan/kumgo-stack/example/server/echosvcserver/buff"
	"github.com/kumparan/kumgo-stack/example/server/echosvcserver/echosvc"
	"google.golang.org/grpc"
)

func init() {
	logger.SetupLoggerAuto("", "")
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "9001"))
	if err != nil {
		logger.Fatal("failed to listen: %v", err)
	}

	logger.Info("Listening on 9001...")

	server := grpc.NewServer()
	buff.RegisterEchoServiceServer(server, echosvc.NewServer())

	if err := server.Serve(lis); err != nil {
		logger.Fatal("failed to serve: %v", err)
	}
}
