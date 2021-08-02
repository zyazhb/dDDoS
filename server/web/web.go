package web

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

func RunWeb() {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*.html")
	router.StaticFS("/js", http.Dir("web/templates/assets/js"))
	router.StaticFS("/css", http.Dir("web/templates/assets/css"))
	router.StaticFS("/img", http.Dir("web/templates/assets/img"))
	router.StaticFS("/plugins", http.Dir("web/templates/assets/plugins"))
	router.StaticFS("/fonts", http.Dir("web/templates/assets/fonts"))

	// 如果使用 LoadHTMLFiles 的话这么做（需要列举所有需要加载的文件，不如上述 LoadHTMLGlob 模式匹配方便）：
	// router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":        "Main website",
			"Total_stream": "56",
			"time":         nowInfo.Now.String(),
			"reason":       nowInfo.Reason,
		})
	})

	router.GET("/indexb", func(c *gin.Context) {

		c.HTML(http.StatusOK, "indexb.html", gin.H{
			"title":        "Main website2",
			"Total_stream": "44",
			"time":         nowInfo.Now.String(),
			"reason":       nowInfo.Reason,
		})
	})

	router.GET("/devices", func(c *gin.Context) {

		c.HTML(http.StatusOK, "devices.html", gin.H{
			"title":        "Main website2",
			"Total_stream": "31",
			"time":         nowInfo.Now.String(),
			"reason":       nowInfo.Reason,
		})
	})

	router.GET("/controller", func(c *gin.Context) {

		c.HTML(http.StatusOK, "controller.html", gin.H{
			"title":        "Main website2",
			"Total_stream": "31",
			"time":         nowInfo.Now.String(),
			"reason":       nowInfo.Reason,
		})
	})

	router.GET("/traffic", func(c *gin.Context) {

		c.HTML(http.StatusOK, "traffic.html", gin.H{
			"title":        "Main website2",
		})
	})

	router.GET("/situationawareness", func(c *gin.Context) {

		c.HTML(http.StatusOK, "situationawareness.html", gin.H{
			"title":        "Main website2",
		})
	})

	router.GET("/log", func(c *gin.Context) {

		c.HTML(http.StatusOK, "log.html", gin.H{
			"title":        "Main website2",
			
		})

	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":80")
}

func GetMsg() {
	for {
		select {
		case <-time.Tick(1 * time.Second):
			nowInfo = <-capture.UpChan
		}
	}
}
