package main

import (
	"github.com/gin-gonic/gin"
)

/*
func (group *RouterGroup) StaticFile(relativePath, filepath string) IRoutes
func (group *RouterGroup) StaticFS(relativePath string, fs http.FileSystem) IRoutes
*/

// UseRoute 使用路由
type UseRoute struct {
	Port, StatusCode int

	Path, Msg string

	ginEngine  *gin.Engine
	ginContext *gin.Context
}

func (r *UseRoute) GetGET() {
	r.ginEngine = gin.Default()

	r.ginEngine.GET(r.Path, func(c *gin.Context) {
		c.String(r.StatusCode, r.Msg)
	})

	err := r.ginEngine.Run(string(rune(r.Port)))

	if err != nil {
		return
	}

	//type HandlerFunc func(*Context)

}

//GetJSON 发送json数据到指定路径

func (r UseRoute) GetJSON(strMap map[string]string) {
	r.ginEngine = gin.Default()
	r.ginEngine.GET("/ping", func(c *gin.Context) {
		c.JSON(r.StatusCode, strMap)
	})
	r.ginEngine.Run() // 在 0.0.0.0:8080 上监听并服务
	//strconv.Itoa(r.Port)
}
