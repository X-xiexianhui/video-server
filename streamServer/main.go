//Package streamServer
/*
   @author:xie
   @date:2022/2/4
   @note:视频流模块
*/
package main

import (
	"net/http"
	"video_server/streamServer/server"
)

func main() {
	r := server.RegisterHandlers()
	mh := server.NewMiddleWareHandler(r, 2)
	http.ListenAndServe(":9000", mh)
}
