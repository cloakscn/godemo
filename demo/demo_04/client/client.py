import requests

request = {
    "id": 0,
    "params": ["cloaks"],
    "method": "HelloService.Hello"
}

responce = requests.post("http://localhost:8080/jsonRPC", json=request)
print(responce.text)