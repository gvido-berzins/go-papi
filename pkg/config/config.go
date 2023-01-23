package config

import (
	"go-papi/pkg/service"
	"go-papi/pkg/storage"
	"runtime"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Conf represents an universal config that contains a database handler and a services.
type Conf struct {
	OSService *service.OSService
	DB        *gorm.DB
}

// New creates a new configuration instance.
func New() *Conf {
	log.Trace().Msg("creating a new config")
	db := storage.SetupDB()
	var s service.OSService
	if runtime.GOOS == "windows" {
		s = service.WindowsService{}
	} else {
		s = service.LinuxService{}
	}
	log.Trace().Msg("new config created")
	return &Conf{OSService: &s, DB: db}
}
