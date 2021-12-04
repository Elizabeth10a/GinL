package main

func main() {
	//r := gin.Default()
	route := UseRoute{

		Port:       8080,
		StatusCode: 200,
		Path:       "/ping",
		Msg:        "okkkkk",
	}
	sendJson(&route)
}

func sendJson(uR *UseRoute) {
	m := make(map[string]string, 10)
	m["ok"] = "ok"
	m["o1k"] = "o23k"
	m["o3k"] = "3ok"

	uR.GetJSON(m)

}
