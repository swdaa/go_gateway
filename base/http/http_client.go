package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// create client
// post request
// process response
// close connection
func main() {
	client := &http.Client{}
	resp, err := client.Get("http://127.0.0.1:8084/hello")
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	bds, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bds))
}
