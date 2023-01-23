package commands

import (
	"go-papi/pkg/config"
	"go-papi/pkg/router"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func createServerCmd(cfg *config.Conf) *cobra.Command {
	var addr string
	serverCmd := &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info().Msg("starting the REST API server")
			if err := router.StartServer(cfg, addr); err != nil {
				log.Fatal().Err(err).Msg("Failure when running the server.")
			}
		},
	}
	serverCmd.Flags().StringVar(&addr, "ip-address", ":9001", "ip address of the server")
	return serverCmd
}
