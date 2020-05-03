package worker

import (
	"fmt"
)

// Job holds the attributes needed to perform unit of work.
type Job struct {
	Url string
	Run func(string)
}

// NewWorker creates takes a numeric id and a channel w/ worker pool.
func NewWorker(id int, workerPool chan chan Job) Worker {
	return Worker{
		id:         id,
		jobQueue:   make(chan Job),
		workerPool: workerPool,
		quitChan:   make(chan bool),
	}
}

//Worker define worker struct
type Worker struct {
	id         int
	jobQueue   chan Job
	workerPool chan chan Job
	quitChan   chan bool
}

func (w Worker) start() {
	go func() {
		for {
			// Add my jobQueue to the worker pool.
			w.workerPool <- w.jobQueue

			select {
			case job := <-w.jobQueue:
				// fmt.Printf("worker %d Started \n", w.id)
				job.Run(job.Url)
				// fmt.Printf("worker %d: completed!\n", w.id)

			case <-w.quitChan:
				// We have been asked to stop.
				fmt.Printf("worker %d completed so stopping\n", w.id)
				return
			}
		}
	}()
}

// stop workers
func (w Worker) stop() {
	go func() {
		w.quitChan <- true
	}()
}

// NewDispatcher creates, and returns a new Dispatcher object.
func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	workerPool := make(chan chan Job, maxWorkers)

	return &Dispatcher{
		jobQueue:   jobQueue,
		maxWorkers: maxWorkers,
		workerPool: workerPool,
	}
}

// Dispatcher define dispatcher
type Dispatcher struct {
	workerPool chan chan Job
	maxWorkers int
	jobQueue   chan Job
}

// run starts numbers of worker based on maxWorkers
func (d *Dispatcher) run() {
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(i+1, d.workerPool)
		worker.start()
	}

	go d.dispatch()
}

// dispatch when a job comes in by jobQueue chanel it assign it to a worker
func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.jobQueue:
			// go func() {
			// fetching workerJobQueue
			workerJobQueue := <-d.workerPool

			// adding the job to workerJobQueue
			workerJobQueue <- job
			// }()
		}
	}
}

// Init create job queue and dispatcher and run dispatcher
func Init(maxQueueSize, maxWorkers int) chan Job {
	// Create the job queue.
	jobQueue := make(chan Job, maxQueueSize)

	// Start the dispatcher.
	dispatcher := NewDispatcher(jobQueue, maxWorkers)
	dispatcher.run()

	return jobQueue
}
