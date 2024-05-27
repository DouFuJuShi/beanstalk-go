package producer

import (
	"github.com/beanstalkd/go-beanstalk"
	"time"
)

type Tube struct {
	endpoint string
	tube     *beanstalk.Tube
}

func (t Tube) Release(tube Tube) {
	// TODO implement me
	panic("implement me")
}

func (t Tube) Get() Tube {
	return t
}

func (t Tube) Close() error {
	return t.tube.Conn.Close()
}

func (t Tube) Reconnect() error {
	conn, err := beanstalk.Dial("tcp", t.endpoint)
	if err != nil {
		return err
	}
	t.tube.Conn = conn
	return nil
}

func (t Tube) Put(body []byte, priority uint32, delay time.Duration, ttr time.Duration) (id uint64, err error) {
	return t.tube.Put(body, priority, delay, ttr)
}
