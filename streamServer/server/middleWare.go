//Package server
/*
   @author:xie
   @date:2022/2/5
   @note:
*/
package server

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandle struct {
	r *httprouter.Router
	l *ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandle{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func (m middleWareHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many request")
		return
	}
	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}
