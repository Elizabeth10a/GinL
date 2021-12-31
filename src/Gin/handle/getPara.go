package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func handle(engine *gin.Engine) {
	//http://192.168.20.129:8080/hello?name=ok
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		path := context.FullPath()
		fmt.Println(path)

		query := context.DefaultQuery("name", "默认值")
		fmt.Println(query)

		context.Writer.Write([]byte("hello," + query)) //输出到页面

	})

	engine.Handle("POST", "/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		fmt.Println(context.PostForm("name")) //从post 得到 name的值
		fmt.Println(context.PostFormMap("name"))
		for i, i2 := range context.PostFormArray("hobby") {
			fmt.Println(i, ":", i2)
		}
		for k, v := range context.PostFormMap("hobby") {
			fmt.Println(k, "：", v)

		}

		context.Writer.Write([]byte("登录成功")) //输出到页面

	})

}
