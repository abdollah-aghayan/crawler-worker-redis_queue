package consumer

import (
	"crawler/queue"
	"crawler/utils/domh"
	"crawler/worker"
	"fmt"
)

// Consumer listen to queue and set new job on receive new job
func Consumer(consume chan worker.Job, queue queue.Queue) {
	runner := func(url string) {
		t, err := domh.GetTitle(url)
		if err != nil {
			fmt.Println(err)
			fmt.Println("title get error")
			return
		}
		fmt.Println(t, " >>>> ", url, "\n")
	}

	for {
		// TODO move 5 to config
		url, err := queue.BPop(5)
		if err != nil {
			continue
		}
		job := worker.Job{
			Url: url,
			Run: runner,
		}

		consume <- job
	}
}
