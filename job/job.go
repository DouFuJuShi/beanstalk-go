package job

import (
	"time"
)

type Job struct {
	id       uint64
	body     []byte
	priority uint32
	delay    time.Duration
	ttr      time.Duration
	reserved bool
}

func (j *Job) ID() uint64 {
	return j.id
}
func (j *Job) SetID(id uint64) {
	j.id = id
}

func (j *Job) Body() []byte {
	return j.body
}

func (j *Job) Priority() uint32 {
	return j.priority
}

func (j *Job) Delay() time.Duration {
	return j.delay
}

func (j *Job) TTR() time.Duration {
	return j.ttr
}

func (j *Job) Reserved() bool {
	return j.reserved
}

func (j *Job) Delete() error {
	return nil
}

func (j *Job) Release() error {
	return nil
}

func (j *Job) Bury() error {
	return nil
}

func (j *Job) Kick() error {
	return nil
}

func (j *Job) Touch() error {
	return nil
}

func (j *Job) Stats() error {
	return nil
}

func NewJob(id uint64, body []byte, priority uint32, delay time.Duration, ttr time.Duration, reserved bool) *Job {
	return &Job{
		id:       id,
		body:     body,
		priority: priority,
		delay:    delay,
		ttr:      ttr,
		reserved: reserved,
	}
}
