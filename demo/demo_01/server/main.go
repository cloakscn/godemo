package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	// http://localhost:8080/add?a=1&b=2
	// 返回的序列化数据 json {"data":3}
	// callID: request.URL.Path
	// data protocol: url
	// net protocol: http
	// encoding/decoding json
	http.HandleFunc("/add", func(writer http.ResponseWriter, request *http.Request) {
		_ = request.ParseForm()
		fmt.Println("path: ", request.URL.Path)
		a, _ := strconv.Atoi(request.Form["a"][0])
		b, _ := strconv.Atoi(request.Form["b"][0])
		// 设置 http 头文件
		writer.Header().Set("Content-Type", "application/json")
		// 序列化
		marshal, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		_, _ = writer.Write(marshal)
	})

	http.ListenAndServe(":8080", nil)
}
