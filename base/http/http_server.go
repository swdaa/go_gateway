package main

import "net/http"

// 1. Register routes and set routing rules
// 2. Start monitoring and provide services
func main() {
	//set path
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world"))
	})

	http.ListenAndServe("127.0.0.1:8084", nil)
}
