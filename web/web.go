package main

import (
	"net/http"/* 
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie" */
	"github.com/gin-gonic/gin"
)



func total_healthe() {

}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("../ui/*.html")
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
	router.Run(":8080")
}