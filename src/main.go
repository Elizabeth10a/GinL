package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default() //返回默认路由引擎

	//指定用户使用GET请求访问， /ping 时，执行 函数内容
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello golang gin",
		})
	})

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 但是，这个规则既能匹配/user/john/格式也能匹配/user/john/send这种格式
	// 如果没有其他路由器匹配/user/john，它将重定向到/user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	router.Run(":9090") // 监听并在 0.0.0.0:9090 上启动服务,默认8000
}
