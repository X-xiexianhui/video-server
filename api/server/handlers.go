package server

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, err := io.WriteString(w, "Create User Handler")
	if err != nil {
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	_, err := io.WriteString(w, uname)
	if err != nil {
		return
	}
}
