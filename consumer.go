package beanstalk

import (
	"time"
)

type Consumer struct {
}

func (c *Consumer) Watch() *Consumer {
	return c
}

func (c *Consumer) ReserveWithFunc(f execFunc, timeout time.Duration) (*ConsumeResult, error) {
	job := c.Reserve(timeout)
	if job == nil {
		return nil, NoJobError
	}

	return f(job), nil
}

func (c *Consumer) Reserve(timeout time.Duration) *Job {
	return &Job{}
}

func NewConsumer() *Consumer {
	return &Consumer{}
}
