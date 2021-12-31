package TltFunc

import (
	"html/template"
)

//type FuncMap map[string]interface{}

type FM struct {
}

func AllMethodMap() template.FuncMap {
	t := template.FuncMap{}
	t["UnixToTime"] = UnixToTime
	t["StringJoin"] = StringJoin
	return t
}
