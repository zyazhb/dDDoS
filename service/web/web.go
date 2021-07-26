package web

import (
	"main/capture"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	nowInfo capture.UpInfo
)

func RunWeb() {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*.html")
	router.StaticFS("/js", http.Dir("web/templates/assets/js"))
	router.StaticFS("/css", http.Dir("web/templates/assets/css"))
	router.StaticFS("/img", http.Dir("web/templates/assets/img"))
	router.StaticFS("/plugins", http.Dir("web/templates/assets/plugins"))
	router.StaticFS("/fonts", http.Dir("web/templates/assets/fonts"))

	go GetMsg()

	// 如果使用 LoadHTMLFiles 的话这么做（需要列举所有需要加载的文件，不如上述 LoadHTMLGlob 模式匹配方便）：
	// router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
		"Total_stream": "5",
		"time": nowInfo.Now.String(),
		"reason": nowInfo.Reason,
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
		case <- time.Tick(1 * time.Second):
			nowInfo = <- capture.UpChan
		}
	}
}
