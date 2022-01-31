//Package dao
/*
   @author:xie
   @date:2022/1/31
   @note:用户session相关功能
*/
package dao

import (
	"database/sql"
	"log"
	"strconv"
	"sync"
	"video_server/api/entity"
)

func InsertSession(sid string, ttl int64, uname string) error {
	ttlStr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("insert into sessions(session_id, TTL, login_name) VALUES (?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sid, ttlStr, uname)
	if err != nil {
		return err
	}
	defer func(stmtIns *sql.Stmt) {
		_ = stmtIns.Close()
	}(stmtIns)
	return nil
}

func RetrieveSession(sid string) (*entity.Session, error) {
	ss := &entity.Session{}
	stmtOut, err := dbConn.Prepare("select TTL,login_name from sessions where session_id=?")
	if err != nil {
		return nil, err
	}
	var (
		ttl   string
		uname string
	)
	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res, err := strconv.ParseInt(ttl, 10, 64); err != nil {
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil, err
	}
	defer func(stmtOut *sql.Stmt) {
		_ = stmtOut.Close()
	}(stmtOut)
	return ss, nil
}

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	for rows.Next() {
		var id string
		var ttlStr string
		var loginName string
		if err := rows.Scan(&id, &ttlStr, &loginName); err != nil {
			log.Printf("retrive sessions error: %s", err)
			break
		}

		if ttl, err1 := strconv.ParseInt(ttlStr, 10, 64); err1 == nil {
			ss := &entity.Session{Username: loginName, TTL: ttl}
			m.Store(id, ss)
			log.Printf(" session id: %s, ttl: %d", id, ss.TTL)
		}

	}

	return m, nil
}

func DeleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare("delete from sessions where session_id=?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	if _, err := stmtOut.Query(sid); err != nil {
		return err
	}
	return nil
}
