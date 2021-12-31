package main

import (
	"GinL/src/Gin/Route"
	"GinL/src/Gin/TltFunc"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	//加载模板函数
	engine.SetFuncMap(TltFunc.AllMethodMap())
	//加载模板 多目录操作-
	engine.LoadHTMLGlob("src/Gin/templates/**/*")

	//配置静态web目录（配置文件）
	engine.Static("/static", "src/Gin/static")

	//engine.LoadHTMLFiles()
	Route.GetXml(engine)
	Route.GetJson(engine)
	Route.GetHtml(engine)
	engine.Run()

}
