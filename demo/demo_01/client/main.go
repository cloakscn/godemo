package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)

func main() {
	fmt.Println(Add(1, 2))
}

type ResponseData struct {
	Data int `json:"data"`
}

func Add(i int, i2 int) int {
	// protocol http
	request := HttpRequest.NewRequest()
	get, _ := request.Get(fmt.Sprintf("http://localhost:8080/%s?a=%d&b=%d", "add", i, i2))
	body, _ := get.Body()
	//fmt.Println(string(body))

	data := ResponseData{}
	json.Unmarshal(body, &data)
	return data.Data
}
