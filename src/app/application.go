package app

import (
	"github.com/gin-gonic/gin"
	"simpleDbMirror/utils/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("applications listening on port 4000...")
	router.Run(":4000")
}
