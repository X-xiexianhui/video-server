//Package web
/*
   @author:xie
   @date:2022/2/6
   @note:前端模块入口
*/
package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// RegisterHandler 注册路由
func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", homeHandler)
	router.POST("/", homeHandler)
	router.GET("/userhome", userHomeHandler)
	router.POST("/userhome", userHomeHandler)
	router.POST("/api", apiHandler)
	router.ServeFiles("/statics/*filepath", http.Dir("./template"))
	router.POST("/upload/:vid-vid", proxyHandler)
	return router
}

func main() {
	r := RegisterHandler()
	http.ListenAndServe(":8080", r)
}
