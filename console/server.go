package console

import (
	"fmt"
	"net"

	"github.com/kumparan/go-lib/logger"
	"github.com/kumparan/kumgo-stack/buff"
	"github.com/kumparan/kumgo-stack/config"
	"github.com/kumparan/kumgo-stack/echosvc"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var runCmd = &cobra.Command{
	Use:   "server",
	Short: "run server",
	Long:  `This subcommand start the server`,
	Run:   run,
}

func init() {
	RootCmd.AddCommand(runCmd)
}

func run(cmd *cobra.Command, args []string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Port()))
	if err != nil {
		logger.Fatal("failed to listen: %v", err)
	}

	logger.Info("Listening on ", config.Port())

	server := grpc.NewServer()
	buff.RegisterEchoServiceServer(server, echosvc.NewServer())

	if err := server.Serve(lis); err != nil {
		logger.Fatal("failed to serve: %v", err)
	}
}
