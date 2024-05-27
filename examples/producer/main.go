package main

import (
	"github.com/DouFuJuShi/beanstalk-go/producer"
)

func main() {
	config := producer.Config{
		Endpoint: "127.0.0.1:11300",
		PoolSize: 10,
		TubeName: "test",
	}

	pool := producer.NewTubePool(config)

	_, err := producer.NewProducer(pool)
	if err != nil {
		panic(err)
	}
}
