package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
)

func WithArgHandler(writer http.ResponseWriter, request *http.Request) {
	value := getAuth()
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(writer, string(value))
}

func getAuth() string {
	key := os.Getenv("key")
	pass := os.Getenv("pass")
	client, err := vod.NewClientWithAccessKey("cn-shanghai", key, pass)

	request := vod.CreateCreateUploadVideoRequest()
	request.FileName = "/Users/lizhe/Movies/test.mov"
	request.Scheme = "https"
	request.Title = "TestVedio"

	response, err := client.CreateUploadVideo(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf(response.GetHttpContentString())
	return response.GetHttpContentString()
}
