//Package dao
/*
 @author:xie
   @date:2022/1/30
   @note:视频功能
*/
package dao

import (
	"database/sql"
	"log"
	"time"
	"video_server/api/entity"
	"video_server/api/utils"
)

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

func GetVideoInfo(vid string) (*entity.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id,name,display_ctime FROM video_info WHERE id=?")
	var (
		aid  int
		dct  string
		name string
	)
	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	defer func(stmtOut *sql.Stmt) {
		_ = stmtOut.Close()
	}(stmtOut)
	res := &entity.VideoInfo{
		Id:           vid,
		AuthorId:     aid,
		Name:         name,
		DisplayCtime: dct,
	}
	return res, err
}

func ListVideoInfo(uname string, from, to int) ([]*entity.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare(`SELECT video_info.id, video_info.author_id, video_info.name, video_info.display_ctime FROM video_info 
		WHERE video_server.users.login_name = ? AND video_info.create_time > FROM_UNIXTIME(?) AND video_info.create_time <= FROM_UNIXTIME(?) 
		ORDER BY video_info.create_time DESC`)
	var res []*entity.VideoInfo

	if err != nil {
		return res, err
	}
	row, err := stmtOut.Query(uname, from, to)
	if err != nil {
		log.Printf("%s", err)
	}
	for row.Next() {
		var id, name, ctime string
		var aid int
		if err := row.Scan(&id, &aid, &name, &ctime); err != nil {
			return res, err
		}
		vi := &entity.VideoInfo{
			Id:           id,
			AuthorId:     aid,
			Name:         name,
			DisplayCtime: ctime,
		}
		res = append(res, vi)
	}
	defer func(stmtOut *sql.Stmt) {
		_ = stmtOut.Close()
	}(stmtOut)
	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}
	defer func(stmtDel *sql.Stmt) {
		_ = stmtDel.Close()
	}(stmtDel)
	return nil
}
