//Package dao
/*
   @author:xie
   @date:2022/2/5
   @note:写入待删除视频记录
*/
package dao

import "log"

func AddVideoDeletionRecord(vid string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO video_del_rec (video_id) VALUES(?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeletionRecord error: %v", err)
		return err
	}

	defer stmtIns.Close()
	return nil
}
