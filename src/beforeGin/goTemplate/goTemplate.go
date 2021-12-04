/*
	https://www.liwenzhou.com/posts/Go/go_template/
	html/template包实现了数据驱动的模板，用于生成可防止代码注入的安全的HTML内容。
	它提供了和text/template包相同的接口，Go语言中输出HTML的场景都应使用html/template这个包。

	事先定义好的HTML文档文件，模板渲染的作用机制可以简单理解为文本替换操作–使用相应的数据去替换HTML文档中事先准备好的标记。
	很多编程语言的Web框架中都使用各种模板引擎，比如Python语言中Flask框架中使用的jinja2模板引擎。
*/

/*

Go语言内置了文本模板引擎text/template和用于HTML文档的html/template。
	模板文件通常定义为.tmpl和.tpl为后缀（也可以使用其他的后缀），必须使用UTF8编码。
	模板文件中使用{{和}}包裹和标识需要传入的数据。
	传给模板这样的数据就可以通过点号（.）来访问，如果数据是复杂类型的数据，可以通过{ { .FieldName }}来访问它的字段。
	除{{和}}包裹的内容外，其他内容均不做修改原样输出。
*/

/*
Go语言模板引擎的使用可以分为三部分：定义模板文件、解析模板文件和模板渲染.

定义模板文件
	其中，定义模板文件时需要我们按照相关语法规则去编写，后文会详细介绍。

解析模板文件
	上面定义好了模板文件之后，可以使用下面的常用方法去解析模板文件，得到模板对象：

	func (t *Template) Parse(src string) (*Template, error)
	func ParseFiles(filenames ...string) (*Template, error)
	func ParseGlob(pattern string) (*Template, error)
当然，你也可以使用func New(name string) *Template函数创建一个名为name的模板，然后对其调用上面的方法去解析模板字符串或模板文件。

模板渲染
渲染模板简单来说就是使用数据去填充模板，当然实际上可能会复杂很多。

	func (t *Template) Execute(wr io.Writer, data interface{}) error
	func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
*/
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	////htmlStr := "<h1>Hello Golang</h1>"
	//byteStr, _ := ioutil.ReadFile("/home/kali/Code/GOLANG/src/GinL/src/FirstHtml.html")
	//_, _ = fmt.Fprintln(w, string(byteStr)) //向 w中写入Str
	//------------------------
	//解析模板
	//tmplPath:="/home/kali/Code/GOLANG/src/GinL/src/beforeGin/goTemplate/Templates/Login.tmpl" //绝对路径 不适合项目迁移
	tmplPath := "src/beforeGin/goTemplate/Templates/Login.tmpl" // 项目路径，go run均正常

	//tmplPath:="./Templates/Login.tmpl"//go build正常 go run报错(找不到文件)，可以按当前文件的相对路径直接打包
	parseFiles, err := template.ParseFiles(tmplPath) // 项目路径
	if err != nil {
		fmt.Printf("parseFiles, err%v\n", err)
		return
	}
	inputStr := "花たん"

	err1 := parseFiles.Execute(w, inputStr) //插入到模板的数据
	if err1 != nil {
		fmt.Printf("parseFiles.Execute%v\n", err1)
		return
	}
	//渲染模板
}
func doIt() {
	http.HandleFunc("/login", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http serve err:%v\n ", err)
		return
	}
}
func main() {
	doIt()
}
