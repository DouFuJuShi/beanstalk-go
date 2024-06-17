package consumer

import (
	"github.com/DouFuJuShi/beanstalk-go/job"
	"github.com/DouFuJuShi/beanstalk-go/types"
	"time"
)

type Consumer struct {
}

func (c *Consumer) Watch() *Consumer {
	return c
}

func (c *Consumer) ReserveWithFunc(f types.ExecFunc, timeout time.Duration) (*types.ConsumeResult, error) {
	reservedJob := c.Reserve(timeout)
	if reservedJob == nil {
		return nil, types.NoJobError
	}

	return f(reservedJob), nil
}

func (c *Consumer) Reserve(timeout time.Duration) *job.Job {
	return &job.Job{}
}

func NewConsumer() *Consumer {
	return &Consumer{}
}
