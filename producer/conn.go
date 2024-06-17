package producer

import (
	"github.com/beanstalkd/go-beanstalk"
	"time"
)

type Conn struct {
	tube *beanstalk.Tube
}

func NewConn(endpoint, tubeName string) *Conn {
	conn, err := beanstalk.DialTimeout("tcp", endpoint, 10*time.Second)
	if err != nil {
		panic(err)
	}
	tube := beanstalk.NewTube(conn, tubeName)
	return &Conn{
		tube,
	}
}

func (conn *Conn) Put(body []byte, priority uint32, delay time.Duration, ttr time.Duration) (id uint64, err error) {
	return conn.tube.Put(body, priority, delay, ttr)
}

func (conn *Conn) Close() error {
	tube := conn.tube
	conn.tube = nil
	return tube.Conn.Close()
}
