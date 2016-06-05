package main

import (
	"fmt"
	"net/url"
	"text/template"
	"os"
)

func main() {
	v := url.Values{}
	v.Set("name", "liyunteng")
	v.Add("friend", "dengxiao")
	v.Add("friend", "lianghengyuan")
	v.Add("friend", "shanglijun")

	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("friend"))
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])

	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}{{end}}!`)
	err = t.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>")
	if err != nil {
		fmt.Println(err)
	}
}
