package producer

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

var ErrClosed = errors.New("conn pool is closed")

type pool struct {
	mutex     sync.Mutex
	_isClosed int32

	conns   []*Conn
	channel chan *Conn
}

func (p *pool) isClosed() bool {
	return atomic.LoadInt32(&p._isClosed) == 1
}

func (p *pool) Get() (*Conn, error) {
	if p.isClosed() {
		return nil, ErrClosed
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: conns为nil 怎么处理
	conn := p.conns[0]
	copy(p.conns, p.conns[1:])
	return conn, nil
}

func (p *pool) Put(conn *Conn) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.isClosed() {
		return
	}
	p.conns = append(p.conns, conn)
}

func (p *pool) Close() error {
	atomic.StoreInt32(&p._isClosed, 1)

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
	conn, err := p.pool.Get()
	if err != nil {
		return 0, err
	}

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
