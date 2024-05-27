package producer

import (
	"errors"
	"github.com/DouFuJuShi/beanstalk-go/job"
	"github.com/DouFuJuShi/beanstalk-go/types"
	"time"
)

type Config struct {
	Endpoint string
	PoolSize int
	TubeName string
}

type Producer struct {
	pc   Config
	pool Pool
}

func (p *Producer) put(body []byte, priority uint32, delay time.Duration, ttr time.Duration) (id uint64, err error) {
	if priority == 0 {
		priority = types.DefaultPriority
	}

	if ttr < time.Second {
		return 0, errors.New("ttr must be greater than 1s")
	}

	if ttr == 0 {
		ttr = time.Minute
	}

	return p.pool.Get().Put(body, priority, delay, ttr)
}

func (p *Producer) Put(job *job.Job) error {
	if job == nil {
		return errors.New("job cannot be nil")
	}

	id, err := p.put(job.Body(), job.Priority(), job.Delay(), job.TTR())
	if err != nil {
		return err
	}

	job.SetID(id)

	return err
}

func NewProducer(pool Pool) (*Producer, error) {
	return &Producer{
		pool: pool,
	}, nil
}
