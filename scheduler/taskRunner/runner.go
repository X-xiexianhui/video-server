//Package taskRunner
/*
   @author:xie
   @date:2022/2/5
   @note:
*/
package taskRunner

type Runner struct {
	Controller controlChan
	Error      controlChan
	Data       dataChan
	dataSize   int
	longLived  bool
	Dispatcher fn
	Executor   fn
}

func NewRunner(size int, longLived bool, d fn, e fn) *Runner {
	return &Runner{
		Controller: make(chan string, 1),
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, size),
		dataSize:   size,
		longLived:  longLived,
		Dispatcher: d,
		Executor:   e,
	}
}

func (r *Runner) startDispatch() {
	defer func() {
		if !r.longLived {
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()
	for {
		select {
		case c := <-r.Controller:
			if c == ReadyToDispatch {
				err := r.Dispatcher(r.Data)
				if err != nil {
					r.Error <- Close
				} else {
					r.Controller <- ReadyToExecute
				}
			}
			if c == ReadyToExecute {
				err := r.Executor(r.Data)
				if err != nil {
					r.Error <- Close
				} else {
					r.Controller <- ReadyToDispatch
				}
			}
		case e := <-r.Error:
			if e == Close {
				return
			}
		default:

		}
	}
}

func (r *Runner) StartAll() {
	r.Controller <- ReadyToDispatch
	r.startDispatch()
}
