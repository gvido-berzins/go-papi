package command

import (
	"go-papi/pkg/config"

	"github.com/gin-gonic/gin"
)

// SetupCommandRoutes is used for setting up "/command" resources
func SetupCommandRoutes(r *gin.Engine, cfg *config.Conf) {
	g := r.Group("/command", gin.BasicAuth(gin.Accounts{"admin": "admin"}))
	g.POST("/shell", createShellHandler(cfg))
	g.GET("/shell/:sessionId", createShellSessionGetHandler(cfg))
}
