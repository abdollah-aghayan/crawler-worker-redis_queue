package main

import (
	"crawler/config"
	"crawler/consumer"
	"crawler/queue"
	"crawler/router/httprouter"
	"crawler/worker"
	"flag"
	"fmt"
)

func main() {

	// load config
	config.Init()

	var (
		maxWorkers   = flag.Int("max_workers", 5, "The number of workers to start")
		maxQueueSize = flag.Int("max_queue_size", 100, "The size of job queue")
		rport        = flag.Int("rport", 6379, "The server port")
	)
	flag.Parse()

	// Init Job queue
	jobQueue := worker.Init(*maxQueueSize, *maxWorkers)
	fmt.Println(*maxWorkers)

	// create queue in redis
	connStr := fmt.Sprintf("%s:%d", "localhost", *rport)
	q := queue.New("urls", connStr)

	// lunch consumer to listen ti=o queue and set new job on receive new job
	go consumer.Consumer(jobQueue, q)

	// router
	httprouter.Run(config.HTTPPort)

}
