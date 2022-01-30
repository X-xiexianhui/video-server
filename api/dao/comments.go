// Package dao
/*
   @author:xie
   @date:2022/1/30
   @note:评论功能
*/
package dao

import (
	"database/sql"
	"video_server/api/entity"
	"video_server/api/utils"
)

func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}
	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id,video_id,author_id,content) VALUES(?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}
	defer func(stmtIns *sql.Stmt) {
		_ = stmtIns.Close()
	}(stmtIns)
	return nil
}

func ListComments(vid string, from, to int) ([]*entity.Comment, error) {
	stmtOut, err := dbConn.Prepare(`select comments.id,users.login_name,comments.content from comments
	inner 	join users on comments.author_id=users.id
	where comments.video_id=? and comments.time>from_unixtime(?) and comments.time<=from_unixtime(?)`)
	var res []*entity.Comment
	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}
		c := &entity.Comment{
			Id:      id,
			VideoId: vid,
			Author:  name,
			Content: content,
		}
		res = append(res, c)
	}
	defer func(stmtOut *sql.Stmt) {
		_ = stmtOut.Close()
	}(stmtOut)
	return res, nil
}
