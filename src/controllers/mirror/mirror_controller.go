package mirror

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDbMirrorConfig(c *gin.Context) {
	c.String(http.StatusOK, "db-mirror-configuration")
}
