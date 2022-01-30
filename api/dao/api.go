package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"video_server/api/entity"
	"video_server/api/utils"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name,pwd) VALUES (?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	err = stmtIns.Close()
	if err != nil {
		return err
	}
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name=?")
	if err != nil {
		log.Panicf("%s", err)
	}

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil {
		return "", err
	}

	return "", nil
}

func DeleteUserCredential(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name=? AND pwd=?")
	if err != nil {
		log.Panicf("%s", err)
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	err = stmtDel.Close()
	if err != nil {
		return err
	}
	return nil
}
func AddNewVideo(aid int, name string) (*entity.VideoInfo, error) {
	// create uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbConn.Prepare("INSERT INTO video_info (id, author_id, name, display_ctime) VALUES(?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	res := &entity.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}

	defer func(stmtIns *sql.Stmt) {
		_ = stmtIns.Close()
	}(stmtIns)
	return res, nil
}
