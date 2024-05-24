package beanstalk

import "github.com/beanstalkd/go-beanstalk"

type conn interface {
	Get() *beanstalk.Conn
	Put(conn *beanstalk.Conn)
	Close() error
	Reset() error
}

type Conn struct {
	conn *beanstalk.Conn
}

type ConnPool struct {
	conn []*beanstalk.Conn
}
