package main

import (
	"net/http"
	"video_server/api/server"
)

func main() {
	r := server.RegisterHandlers()
	mh := server.NewMiddleWareHandle(r)
	err := http.ListenAndServe(":8800", mh)
	if err != nil {
		return
	}
}
