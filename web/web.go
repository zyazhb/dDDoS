package main

import (
	"main/capture"
	"github.com/gin-gonic/gin"
)

func RunWeb() {
	router := gin.Default()
	router.LoadHTMLGlob("../ui/*.html")
  router.LoadHTMLGlob("web/templates/*")
	router.StaticFS("/js", http.Dir("../ui/assets/js"))
	router.StaticFS("/css", http.Dir("../ui/assets/css"))
	router.StaticFS("/img", http.Dir("../ui/assets/img"))
	router.StaticFS("/plugins", http.Dir("../ui/assets/plugins"))
	router.StaticFS("/fonts", http.Dir("../ui/assets/fonts"))
  	
	// 如果使用 LoadHTMLFiles 的话这么做（需要列举所有需要加载的文件，不如上述 LoadHTMLGlob 模式匹配方便）：
	// router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
	  c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
		"Total_stream": "123456789",
	  })
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/message", UpdateInfo)
	router.Run(":80")
}

func UpdateInfo(c *gin.Context) {
	nowInfo := <- capture.UpChan

	c.HTML(200, "info.html", gin.H{
			"time": nowInfo.Now.String(),
			"reason": nowInfo.Reason,
		})
}
