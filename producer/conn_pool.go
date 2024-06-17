package producer

import (
	"sync"
	"time"
)

type pool struct {
	mutex sync.Mutex
	conns []*Conn
}

func (p *pool) Get() *Conn {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	conn := p.conns[0]
	p.conns = p.conns[1:]
	return conn
}

func (p *pool) Put(conn *Conn) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.conns = append(p.conns, conn)
}

func (p *pool) Close() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	conns := p.conns
	p.conns = nil
	for _, conn := range conns {
		conn.Close()
	}
	return nil
}

type Pool struct {
	pool *pool
}

func (p *Pool) Put(body []byte, priority uint32, delay time.Duration, ttr time.Duration) (id uint64, err error) {
	conn := p.pool.Get()

	defer p.pool.Put(conn)

	return conn.Put(body, priority, delay, ttr)
}

func (p *Pool) Close() error {
	return p.pool.Close()
}

func NewPool(endpoint string, tube string, size uint) *Pool {
	if size == 0 {
		size = 1
	}
	conns := make([]*Conn, 0, size)

	for i := uint(0); i < size; i++ {
		conns = append(conns, NewConn(endpoint, tube))
	}

	return &Pool{pool: &pool{conns: conns}}
}
