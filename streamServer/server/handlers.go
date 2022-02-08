//Package server
/*
   @author:xie
   @date:2022/2/4
   @note:视频流业务处理
*/
package server

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"video_server/streamServer/config"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := config.VideoDir + vid
	video, err := os.Open(vl)
	if err != nil {
		log.Printf("Open video failed")
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
	video.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, config.MaxUploadSize)
	if err := r.ParseMultipartForm(config.MaxUploadSize); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too large")
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error:%v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	fileName := p.ByName("vid-id")
	err = ioutil.WriteFile(config.VideoDir+fileName, data, 0666)
	if err != nil {
		log.Printf("Write file failed")
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, err = io.WriteString(w, "uploaded successfully")
	if err != nil {
		return
	}
}
