package main

import (
	"net/http"
	"video_server/api/server"
)

func main() {
	r := server.RegisterHandlers()
	http.ListenAndServe(":8800", r)
}
