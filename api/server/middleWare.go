//Package server
/*
   @author:xie
   @date:2022/1/31
   @note:中间层，拦截请求
*/
package server

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandle(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	ValidateUserSession(r)
	m.r.ServeHTTP(w, r)
}
