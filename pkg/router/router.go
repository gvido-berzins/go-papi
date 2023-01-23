package router

import (
	"net/http"

	"go-papi/pkg/config"
	"go-papi/pkg/router/command"
	"go-papi/pkg/router/file"
	"go-papi/pkg/router/host"

	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"github.com/go-oauth2/oauth2/server"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/store"
)

// StartServer starts the REST API and waits for it to finish.
func StartServer(cfg *config.Conf, addr string) error {

	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(store.NewFileTokenStore("data.db"))
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)

	// Initialize the oauth2 service
	ginserver.InitServer(manager)
	ginserver.SetAllowGetAccessRequest(true)
	ginserver.SetClientInfoHandler(server.ClientFormHandler)

	r := gin.Default()
	r.Any("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello"})
	})
	auth := r.Group("/oauth2")
	auth.GET("/token", ginserver.HandleTokenRequest)

	r.Use(ginserver.HandleTokenVerify())

	file.SetupFileRoutes(r)
	command.SetupCommandRoutes(r, cfg)
	host.SetupHostRoutes(r, cfg)

	if addr == "" {
		addr = ":9001"
	}
	return r.Run(addr)
}
