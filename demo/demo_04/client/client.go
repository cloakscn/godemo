package main

import (
	"encoding/json"

	"github.com/kirinlabs/HttpRequest"
)

type ResponseData struct {
	Data int `json:"data"`
}

func main() {

	marshal, _ := json.Marshal(struct {
		id     int
		params []string
		method string
	}{
		id:     0,
		params: []string{"cloaks"},
		method: "HelloService.Hello",
	})
	// protocol http
	request := HttpRequest.NewRequest()
	get, _ := request.Post("http://localhost:8080/jsonRPC", marshal)
	body, _ := get.Body()
	//fmt.Println(string(body))

	data := ResponseData{}
	json.Unmarshal(body, &data)
}

// json {"method": "HelloService.Hello", "params": ["cloaks"], "id": 0}
