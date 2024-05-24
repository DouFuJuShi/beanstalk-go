package beanstalk

type Job struct {
	id       uint64
	priority uint32
	reserved bool
	body     []byte
}

func (j *Job) ID() uint64 {
	return j.id
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

func NewJob(id uint64, priority uint32, reserved bool) *Job {
	return &Job{
		id:       id,
		priority: priority,
		reserved: reserved,
	}
}
