package producer

import (
	"sync"
	"time"
)

type Pool interface {
	Get() *Tube
	Release(tube Tube)
	Close() error
	Reconnect() error

	Put(body []byte, priority uint32, delay time.Duration, ttr time.Duration) (id uint64, err error)
}

type TubePool struct {
	endpoint string
	pool     []*Tube
	lock     sync.Mutex
}

func (t *TubePool) Release(tube *Tube) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.pool = append(t.pool, tube)
}

func (t *TubePool) Close() error {
	t.lock.Lock()
	defer t.lock.Unlock()
	for _, t := range t.pool {
		_ = t.Close()
	}
	t.pool = nil
	return nil
}

func (t *TubePool) Reconnect() error {
	t.lock.Lock()
	defer t.lock.Unlock()
	for _, t := range t.pool {
		_ = t.Reconnect()
	}
	return nil
}

func (t *TubePool) Get() *Tube {
	t.lock.Lock()
	defer t.lock.Unlock()
	tube := t.pool[0]
	t.pool = t.pool[1:]
	return tube
}

func (t *TubePool) Put(body []byte, priority uint32, delay time.Duration, ttr time.Duration) (id uint64, err error) {
	return t.Get().Put(body, priority, delay, ttr)
}
