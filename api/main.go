package main

import (
	"net/http"
	"video_server/api/server"
)

func main() {
	r := server.RegisterHandlers()
	err := http.ListenAndServe(":8800", r)
	if err != nil {
		return
	}
}
