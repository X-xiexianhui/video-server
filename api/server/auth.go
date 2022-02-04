//Package server
/*
   @author:xie
   @date:2022/1/31
   @note:身份验证
*/
package server

import (
	"net/http"
	"video_server/api/entity"
	"video_server/api/session"
)

var (
	HeaderFieldSession = "X-Session-Id"
	HeaderFieldUname   = "X-User-Name"
)

func ValidateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HeaderFieldSession)
	if len(sid) == 0 {
		return false
	}
	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HeaderFieldUname, uname)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HeaderFieldUname)
	if len(uname) == 0 {
		sendErrorResponse(w, entity.ErrorNotAuthUser)
		return false
	}
	return true
}
