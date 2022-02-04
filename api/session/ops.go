//Package session
/*
   @author:xie
   @date:2022/1/31
   @note:
*/
package session

import (
	"log"
	"sync"
	"time"
	"video_server/api/dao"
	"video_server/api/entity"
	"video_server/api/utils"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func NowInMillion() int64 {
	return time.Now().UnixNano() / 100000
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	err := dao.DeleteSession(sid)
	if err != nil {
		return
	}
}
func LoadSessionFromDB() {
	r, err := dao.RetrieveAllSessions()
	if err != nil {
		log.Printf("Load Session From DB Failed")
		return
	}

	r.Range(func(key, value interface{}) bool {
		ss := value.(*entity.Session)
		sessionMap.Store(key, ss)
		return true
	})
}

func GenerateNewSessionId(uname string) string {
	id, _ := utils.NewUUID()
	ctime := NowInMillion()
	ttl := ctime + 30*60*1000 //session过期时间，30min
	ss := &entity.Session{
		Username: uname,
		TTL:      ttl,
	}
	sessionMap.Store(id, ss)
	err := dao.InsertSession(id, ttl, uname)
	if err != nil {
		return "Insert Into Session failed"
	}
	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ctime := NowInMillion()
		if ss.(*entity.Session).TTL < ctime {
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*entity.Session).Username, false
	}
	return "", true
}
