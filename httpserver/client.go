package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	fmt.Println("开始请求百度网页")
	//Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36
	resp, err := http.Get("https://www.baidu.com/")
	if err != nil {
		panic(err)
	} else {
		defer resp.Body.Close()
	}


	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Println(s)
}
