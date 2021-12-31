package Route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
import "GinL/src/Gin/model"

func GetJson(engine *gin.Engine) {
	engine.GET("/json", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"name": "as",
			"age":  12,
		})
	})
	engine.GET("/json1", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"name": "a2s",
			"age":  1234,
		})
	})

	engine.GET("/json2", func(context *gin.Context) {
		user := &model.User{Name: "As", Addr: "鞍山市"}
		context.JSON(200, &user)
	})

	engine.GET("/jsonp", func(context *gin.Context) {
		user := &model.User{Name: "As", Addr: "鞍山市"}

		//可以读取url中的参数
		context.JSONP(200, &user)
	})

}
func GetXml(engine *gin.Engine) {
	engine.GET("/xml", func(context *gin.Context) {

		context.XML(http.StatusOK, gin.H{
			"user": &model.User{Name: "Ad", Addr: "DDs"},
		})
	})
}
func GetHtml(engine *gin.Engine) {

	//加载模板目录中的文件 ，文件全称
	engine.GET("/back/goods", func(context *gin.Context) {
		//src/Gin/templates/goods.html
		context.HTML(http.StatusOK, "back/goods.html", gin.H{
			"user":    &model.User{Name: "back goods数据", Addr: "DDs"},
			"names":   []string{"a", "b", "c"},
			"nilList": []string{},
			"date":    int(time.Now().Unix()), //当前时间的时间戳
		})
	})
	engine.GET("/font/news", func(context *gin.Context) {
		//src/Gin/templates/goods.html
		context.HTML(http.StatusOK, "font/news.html", gin.H{
			"user": &model.User{Name: "font news数据", Addr: "DDs"},
		})
	})
	engine.GET("/", func(context *gin.Context) {
		//src/Gin/templates/goods.html
		context.HTML(http.StatusOK, "index.html", gin.H{
			"user": &model.User{Name: "font news数据", Addr: "DDs"},
		})
	})
	engine.GET("/pub/head", func(context *gin.Context) {
		//src/Gin/templates/goods.html
		context.HTML(http.StatusOK, "pub/head.html", gin.H{
			"us":   &model.User{Name: "font news数据", Addr: "阿斯爱上的顿"},
			"mode": "爱上是否",
			"date": 1640869714,
		})
	})

}
