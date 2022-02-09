//Package server
/*
   @author:xie
   @date:2022/2/4
   @note:注册路由
*/
package server

import "github.com/julienschmidt/httprouter"

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid-id", streamHandler)
	router.POST("/upload/:vid-id", uploadHandler)
	return router
}
