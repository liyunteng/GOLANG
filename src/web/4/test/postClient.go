package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// 关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer", err)
		return err
	}

	// 打开文件句柄
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file", err)
		return err
	}
	defer fh.Close()


	// iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		fmt.Println("io Copy", err)
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		fmt.Println("post", err)
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll", err)
		return err
	}

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

func main() {
	target_url := "http://localhost:9090/upload"
	filename := "./postClient.go"
	postFile(filename, target_url)
}
