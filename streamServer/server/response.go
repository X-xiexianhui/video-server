//Package server
/*
   @author:xie
   @date:2022/2/4
   @note:消息应答
*/
package server

import (
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, sc int, errMsg string) {
	w.WriteHeader(sc)
	_, err := io.WriteString(w, errMsg)
	if err != nil {
		return
	}
}
