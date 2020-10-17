package main

import (
	"flygo/fly"
	"fmt"
	"net/http"
)

func main()  {
	// 设置运行时最大核数
	//num := runtime.NumCPU()
	//fmt.Println(num)
	//runtime.GOMAXPROCS(num)
	r := fly.Default()
	fmt.Println("new")
	r.GET("/api_test", func(w http.ResponseWriter, req *http.Request)  {
		fmt.Fprintf(w, "api test URL.Path = %q\n", req.URL.Path)
	})
	r.Run(":8011")

}
