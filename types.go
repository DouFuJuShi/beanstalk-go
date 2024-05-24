package beanstalk

import "errors"

type JobCommand int32

const (
	Delete JobCommand = iota + 1
	Bury   JobCommand = iota + 1
)

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

type execFunc func(job *Job) *ConsumeResult

var (
	// NoJobError no job
	NoJobError = errors.New("no job")
)
