package mirror

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpleDbMirror/services"
	"time"
)

func Config(c *gin.Context) {
	resp, _:= services.MirrorService.GetConfig()
	c.String(http.StatusOK, *resp)
}

func Activate(c *gin.Context) {
	resp, _ := services.MirrorService.Activate()
	c.String(http.StatusOK, *resp)
}

func Deactivate(c *gin.Context) {
	resp, _ := services.MirrorService.Deactivate()
	c.String(http.StatusOK, *resp)
}

func GetLogs(c *gin.Context) {
	strDate := c.Param("DATE")
	date, _ := time.Parse("01/02/2006", strDate)
	resp, _ := services.MirrorService.GetLogs(date)
	c.String(http.StatusOK, *resp)
}

func Execute(c *gin.Context) {
	resp, _ := services.MirrorService.Execute()
	c.String(http.StatusOK, *resp)
}

func Status(c *gin.Context) {
	resp, _ := services.MirrorService.Status()
	c.String(http.StatusOK, *resp)
}
