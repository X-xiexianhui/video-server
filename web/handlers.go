//Package main
/*
   @author:xie
   @date:2022/2/6
   @note:请求处理
*/
package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type HomePage struct {
	Name string
}

type UserPage struct {
	Name string
}

func homeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cname, err1 := r.Cookie("username")
	sid, err2 := r.Cookie("session")
	if err1 != nil || err2 != nil {
		page := &HomePage{Name: "xie"}
		t, err := template.ParseFiles("./templates/home.html")
		if err != nil {
			log.Printf("Parsing template error:%s", err)
			return
		}
		t.Execute(w, page)
		return
	}
	if len(cname.Value) != 0 && len(sid.Value) != 0 {
		http.Redirect(w, r, "/userhome", http.StatusFound)
		return
	}

}

func userHomeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cname, err1 := r.Cookie("username")
	_, err2 := r.Cookie("session")
	if err2 != nil || err1 != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	}
	fName := r.FormValue("username")
	var page *UserPage
	if len(cname.Value) != 0 {
		page = &UserPage{Name: cname.Value}
	} else if len(fName) != 0 {
		page = &UserPage{Name: fName}
	}
	t, err := template.ParseFiles("./templates/userhome.html")
	if err != nil {
		log.Printf("Parsing userhome error:%s", err)
		return
	}
	t.Execute(w, page)
	return
}

func apiHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != http.MethodPost {
		re, _ := json.Marshal(ErrorRequestNotRecognized)
		io.WriteString(w, string(re))
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	apiBody := &ApiBody{}
	if err := json.Unmarshal(res, apiBody); err != nil {
		re, _ := json.Marshal(ErrorRequestBodyParseFailed)
		io.WriteString(w, string(re))
		return
	}

	request(apiBody, w, r)
	defer r.Body.Close()
}

func proxyHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u, _ := url.Parse("http://127.0.0.1:9000/")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}

func loadProxyHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u, _ := url.Parse("http://127.0.0.1:9000/")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}
