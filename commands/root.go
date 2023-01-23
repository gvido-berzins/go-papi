package commands

import (
	"go-papi/pkg/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var logLevelMap = map[string]zerolog.Level{
	"trace": zerolog.TraceLevel,
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
	"panic": zerolog.PanicLevel,
}

// Execute sets up the CLI interface and executes the user provided command.
func Execute() error {
	rootCmd := &cobra.Command{
		Use: "papi",
	}
	if err := setup(); err != nil {
		return err
	}
	log.Trace().Msg("Setting up commands.")

	cfg := config.New()
	rootCmd.AddCommand(createServerCmd(cfg))
	log.Trace().Msg("rootCmd.Execute() reached")
	return rootCmd.Execute()
}

func setup() error {
	setupEnv()
	setupLogging()
	return nil
}

func setupEnv() error {
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/papi")
	viper.AddConfigPath("~/.config/papi")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.SetEnvPrefix("papi")
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}

func setupLogging() {
	logLevel := viper.GetString("log_level")
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerologLevel, ok := logLevelMap[logLevel]
	if !ok {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Warn().Msg("Failed to change log level, set to Info.")
	} else {
		zerolog.SetGlobalLevel(zerologLevel)
	}
}
