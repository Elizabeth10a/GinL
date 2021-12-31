package TltFunc

import "time"

//模板函数， 在模板之前 添加

func UnixToTime(timeStamp int) string {
	unix := time.Unix(int64(timeStamp), 0)
	return unix.Format("2006-01-02 15:04:05")
}
func StringJoin(str, str2 string) string {
	return str + "------" + str2
}
