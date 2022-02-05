//Package taskRunner
/*
   @author:xie
   @date:2022/2/5
   @note:
*/
package taskRunner

const (
	ReadyToDispatch = "d"
	ReadyToExecute  = "e"
	Close           = "c"
	VideoPath       = "./videos/"
)

type controlChan chan string
type dataChan chan interface{}
type fn func(dc dataChan) error
