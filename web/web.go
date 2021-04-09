package web

import (
	"main/capture"

	"github.com/gin-gonic/gin"
)

func RunWeb() {
	r := gin.Default()

	r.LoadHTMLGlob("web/templates/*")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/message", UpdateInfo)
	r.Run(":80")
}

func UpdateInfo(c *gin.Context) {
	nowInfo := <- capture.UpChan

	c.HTML(200, "info.html", gin.H{
			"time": nowInfo.Now.String(),
			"reason": nowInfo.Reason,
		})
}
