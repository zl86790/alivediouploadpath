package main

import (
	"net/http"
	handler "vedioUploadPath/handler"
)

func main() {
	http.HandleFunc("/", handler.WithArgHandler)
	http.ListenAndServe(":8000", nil)
}
