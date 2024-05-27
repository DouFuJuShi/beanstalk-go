package producer

import "time"

type Pool interface {
	Release(tube Tube)
	Close() error
	Reconnect() error

	Get() Tube
	Put(body []byte, priority uint32, delay time.Duration, ttr time.Duration) (id uint64, err error)
}

type TubePool struct {
	tubes []Tube
}

func (t TubePool) Release(tube Tube) {
	// TODO implement me
	panic("implement me")
}

func (t TubePool) Close() error {
	// TODO implement me
	panic("implement me")
}

func (t TubePool) Reconnect() error {
	// TODO implement me
	panic("implement me")
}

func (t TubePool) Get() Tube {
	// TODO implement me
	panic("implement me")
}

func (t TubePool) Put(body []byte, priority uint32, delay time.Duration, ttr time.Duration) (id uint64, err error) {
	// TODO implement me
	panic("implement me")
}
