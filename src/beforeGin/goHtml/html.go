package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//htmlStr := "<h1>Hello Golang</h1>"
	byteStr, _ := ioutil.ReadFile("/home/kali/Code/GOLANG/src/GinL/src/FirstHtml.html")
	_, _ = fmt.Fprintln(w, string(byteStr)) //向 w中写入Str

}
func doIt() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http serve err:%v\n ", err)

		return
	}
}
func main() {
	doIt()
}
