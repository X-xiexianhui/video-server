package server

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"video_server/api/dao"
	"video_server/api/entity"
	"video_server/api/session"
	"video_server/api/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	uBody := &entity.UserCredential{}

	if err := json.Unmarshal(res, uBody); err != nil {
		sendErrorResponse(w, entity.ErrorRequestBodyParseFailed)
		return
	}

	if err := dao.AddUserCredential(uBody.Username, uBody.Pwd); err != nil {
		sendErrorResponse(w, entity.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(uBody.Username)
	su := &entity.SignedUp{Success: true, SessionId: id}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, entity.ErrorSessionFaults)
		return
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	log.Printf("%s", res)
	uBody := &entity.UserCredential{}
	if err := json.Unmarshal(res, uBody); err != nil {
		log.Printf("%s", err)
		//io.WriteString(w, "wrong")
		sendErrorResponse(w, entity.ErrorRequestBodyParseFailed)
		return
	}

	// Validate the request body
	uname := p.ByName("username")
	log.Printf("Login url name: %s", uname)
	log.Printf("Login body name: %s", uBody.Username)
	if uname != uBody.Username {
		sendErrorResponse(w, entity.ErrorNotAuthUser)
		return
	}

	log.Printf("%s", uBody.Username)
	pwd, err := dao.GetUserCredential(uBody.Username)
	log.Printf("Login pwd: %s", pwd)
	log.Printf("Login body pwd: %s", uBody.Pwd)
	if err != nil || len(pwd) == 0 || pwd != uBody.Pwd {
		sendErrorResponse(w, entity.ErrorNotAuthUser)
		return
	}

	id := session.GenerateNewSessionId(uBody.Username)
	si := &entity.SignedIn{Success: true, SessionId: id}
	if resp, err := json.Marshal(si); err != nil {
		sendErrorResponse(w, entity.ErrorSessionFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

func GetUserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !ValidateUser(w, r) {
		log.Printf("Unathorized user \n")
		return
	}

	uname := p.ByName("username")
	u, err := dao.GetUser(uname)
	if err != nil {
		log.Printf("Error in GetUserInfo: %s", err)
		sendErrorResponse(w, entity.ErrorDBError)
		return
	}

	ui := &entity.UserInfo{Id: u.Id}
	if resp, err := json.Marshal(ui); err != nil {
		sendErrorResponse(w, entity.ErrorSessionFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

func AddNewVideo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !ValidateUser(w, r) {
		log.Printf("Unaythorized user\n")
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	nvBody := &entity.NewVideo{}
	if err := json.Unmarshal(res, nvBody); err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, entity.ErrorRequestBodyParseFailed)
		return
	}

	vi, err := dao.AddNewVideo(nvBody.AuthorId, nvBody.Name)
	if err != nil {
		log.Printf("Error in AddNewVideo:%s", err)
		sendErrorResponse(w, entity.ErrorDBError)
		return
	}
	if resp, err := json.Marshal(vi); err != nil {
		sendErrorResponse(w, entity.ErrorSessionFaults)
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}

func ListAllVideos(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !ValidateUser(w, r) {
		log.Printf("Unaythorized user\n")
		return
	}
	uname := p.ByName("username")
	vs, err := dao.ListVideoInfo(uname, 0, utils.GetCurrentTimestampSec())
	if err != nil {
		log.Printf("Error in ListVideo:%s", err)
		sendErrorResponse(w, entity.ErrorDBError)
		return
	}
	vsi := &entity.VideosInfo{Videos: vs}
	if resp, err := json.Marshal(vsi); err != nil {
		sendErrorResponse(w, entity.ErrorSessionFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

func DeleteVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !ValidateUser(w, r) {
		log.Printf("Unauthorized user\n")
		return
	}
	vid := p.ByName("vid-id")
	err := dao.DeleteVideoInfo(vid)
	if err != nil {
		log.Printf("Error in Delete Video:%s", err)
		sendErrorResponse(w, entity.ErrorDBError)
		return
	}
	go utils.SendDeleteVideoRequest(vid)
	sendNormalResponse(w, "", 204)
}

func PostComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !ValidateUser(w, r) {
		log.Printf("Unauthorized user\n")
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	cBody := &entity.NewComment{}
	if err := json.Unmarshal(reqBody, cBody); err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, entity.ErrorRequestBodyParseFailed)
		return
	}
	vid := p.ByName("vid-id")
	if err := dao.AddNewComments(vid, cBody.AuthorId, cBody.Content); err != nil {
		log.Printf("Error in post comment")
		sendErrorResponse(w, entity.ErrorDBError)
	} else {
		sendNormalResponse(w, "ok", 201)
	}
}

func ShowComments(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !ValidateUser(w, r) {
		log.Printf("Show comments failed\n")
		return
	}
	vid := p.ByName("vid-id")
	cm, err := dao.ListComments(vid, 0, utils.GetCurrentTimestampSec())
	if err != nil {
		log.Printf("Error in list comments")
		sendErrorResponse(w, entity.ErrorDBError)
		return
	}
	cms := &entity.Comments{Comments: cm}
	if resp, err := json.Marshal(cms); err != nil {
		sendErrorResponse(w, entity.ErrorSessionFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}
