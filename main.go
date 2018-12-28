package main

import (
	"net/http"
	"fmt"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	fmt.Println("开启服务")
	http.ListenAndServe(":12027", nil)
}
