package main

import (
	"fmt"
	"github.com/DouFuJuShi/beanstalk-go/job"
	"github.com/DouFuJuShi/beanstalk-go/producer"
)

func main() {
	p, err := producer.NewProducer("127.0.0.1:11311", "test", 10)
	err = p.Put(&job.Job{})
	fmt.Println(err)
}
