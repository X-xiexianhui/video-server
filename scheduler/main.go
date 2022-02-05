//Package scheduler
/*
   @author:xie
   @date:2022/2/5
   @note:
*/
package main

import (
	"net/http"
	"video_server/scheduler/server"
	"video_server/scheduler/taskRunner"
)

func main() {
	go taskRunner.Start()
	r := server.RegisterHandlers()
	http.ListenAndServe(":9001", r)
}
