package types

import (
	"errors"
	"time"
)

const Version = "0.0.1"

type JobCommand int32

const (
	Delete JobCommand = iota + 1
	Bury   JobCommand = iota + 1
)

var (
	// NoJobError no job
	NoJobError = errors.New("no job")

	// JobExistError job exist
	JobExistError = errors.New("job exist")
)

const (
	DefaultPriority = uint32(1024)
	DefaultTTR      = time.Minute
)
