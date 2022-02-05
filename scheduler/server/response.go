//Package server
/*
   @author:xie
   @date:2022/2/5
   @note:
*/
package server

import (
	"io"
	"net/http"
)

func sendResponse(w http.ResponseWriter, sc int, resp string) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
