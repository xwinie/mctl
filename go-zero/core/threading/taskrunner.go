package threading

import (
	"github.com/wenj91/mctl/go-zero/core/lang"
	"github.com/wenj91/mctl/go-zero/core/rescue"
)

type TaskRunner struct {
	limitChan chan lang.PlaceholderType
}

func NewTaskRunner(concurrency int) *TaskRunner {
	return &TaskRunner{
		limitChan: make(chan lang.PlaceholderType, concurrency),
	}
}

func (rp *TaskRunner) Schedule(task func()) {
	rp.limitChan <- lang.Placeholder

	go func() {
		defer rescue.Recover(func() {
			<-rp.limitChan
		})

		task()
	}()
}
