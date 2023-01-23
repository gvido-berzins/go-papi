package main

import (
	"go-papi/commands"

	"github.com/rs/zerolog/log"
)

func main() {
	if err := commands.Execute(); err != nil {
		log.Fatal().Err(err).Msg("Failed to execute the root command.")
	}
}
