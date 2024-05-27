package types

import (
	"github.com/DouFuJuShi/beanstalk-go/job"
)

type ExecFunc func(job *job.Job) *ConsumeResult

type ConsumeResult struct {
	nextCommand JobCommand
	err         error
}

func (cr ConsumeResult) NextCommand() JobCommand {
	return cr.nextCommand
}

func (cr ConsumeResult) Err() error {
	return cr.err
}

func NewConsumeResult(nextCommand JobCommand, err error) ConsumeResult {
	return ConsumeResult{nextCommand, err}
}
