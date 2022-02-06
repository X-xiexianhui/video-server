package main

import (
	"net/http"
	"video_server/api/server"
	"video_server/api/session"
)

func main() {
	session.LoadSessionFromDB()
	r := server.RegisterHandlers()
	mh := server.NewMiddleWareHandle(r)
	err := http.ListenAndServe(":8000", mh)
	if err != nil {
		return
	}
}
