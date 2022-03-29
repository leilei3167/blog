package main

import (
	"blog/pkg/setting"
	"blog/router"
	"fmt"
	"net/http"
)

func main() {
	r := router.InitRouters()
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.HTTP_PORT),
		Handler: r,
	}

	s.ListenAndServe()

}
