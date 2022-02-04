package server

import (
	"encoding/json"
	"io"
	"net/http"
	"video_server/api/entity"
)

func sendErrorResponse(w http.ResponseWriter, errResp entity.ErrorResponse) {
	w.WriteHeader(errResp.HttpSC)
	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
}
func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
