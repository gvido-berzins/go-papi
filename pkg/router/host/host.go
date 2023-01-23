package host

import (
	"go-papi/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func SetupHostRoutes(r *gin.Engine, cfg *config.Conf) {
	g := r.Group("/host")
	g.GET("/drives", createDrivesHandler(cfg))
	g.GET("/ram", func(c *gin.Context) {
		log.Debug().Msg("called ram")
	})
	g.GET("/resources", func(c *gin.Context) {
		log.Debug().Msg("called resources")
	})
}

func createDrivesHandler(cfg *config.Conf) func(*gin.Context) {
	return func(c *gin.Context) {
		log.Debug().Msg("called drives")
	}
}
