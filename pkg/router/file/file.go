package file

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type FileRequest struct {
	Path string
}

func SetupFileRoutes(r *gin.Engine) {
	g := r.Group("/file")
	g.GET("/download", func(c *gin.Context) {
		var body FileRequest
		if err := c.BindJSON(&body); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		c.File(body.Path)
	})
	g.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		dst, ok := c.GetPostForm("dst")
		if !ok {
			dst = "uploads/"
		}
		log.Debug().Msgf("Uploading file with size %d", file.Size)
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to upload file to '%s'", dst))
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
}
