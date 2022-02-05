//Package server
/*
   @author:xie
   @date:2022/2/5
   @note:
*/
package server

import "github.com/julienschmidt/httprouter"

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)

	return router
}
