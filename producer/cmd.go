package producer

import "time"

type Cmd interface {
	Put(body []byte, priority uint32, delay time.Duration, ttr time.Duration) (id uint64, err error)
	Close() error
}
