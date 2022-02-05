//Package taskRunner
/*
   @author:xie
   @date:2022/2/5
   @note:
*/
package taskRunner

import (
	"errors"
	"log"
	"os"
	"sync"
	"video_server/scheduler/dao"
)

func deleteVideo(vid string) error {
	err := os.Remove(VideoPath + vid)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}
func VideoClearDispatcher(dc dataChan) error {
	res, err := dao.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("Video Clear dispatcher error:%v", err)
		return err
	}
	if len(res) == 0 {
		return errors.New("all tasks finished")
	}
	for _, id := range res {
		dc <- id
	}
	return nil
}

func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error
forLoop:
	for {
		select {
		case vid := <-dc:
			go func(id interface{}) {
				if err := deleteVideo(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
				if err := dao.DelVideoDeletionRecord(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
			}(vid)
		default:
			break forLoop
		}
	}
	errMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err != nil {
			return false
		}
		return true
	})

	return err
}
