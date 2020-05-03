package queue

import (
	"github.com/gomodule/redigo/redis"
)

var queue *Queue

// Queue define queue struct
type Queue struct {
	pool *redis.Pool
	key  string
}

// GetQueue return queue
func GetQueue() *Queue {
	if queue == nil {
		panic("You have not initioated a queue")
	}

	return queue
}

// New create a queue in the given connection string redis server
func New(queueName string, redisConnStr string) Queue {
	pool := newPool(redisConnStr)
	queue = &Queue{
		pool: pool,
		key:  queueName,
	}

	return *queue
}

// create new pool of connections
func newPool(connStr string) *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", connStr)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// Push an element onto the queue
func (q *Queue) Push(el string) (int64, error) {
	return redis.Int64(q.do("LPUSH", q.key, el))
}

//BPop Block and Pop an element off the queue
// Use timeout_secs = 0 to block indefinitely
// On timeout, this DOES return an error because redigo does.
func (q *Queue) BPop(timeout_secs int) (string, error) {
	reply, err := q.do("BRPOP", q.key, timeout_secs)

	if err != nil && reply == nil {
		return "", err
	}

	// convert reply to slice of string
	res, err := redis.Strings(reply, err)
	if err != nil {
		return "", err
	}

	return res[1], nil
}

// get a connection from pool and run the command and close the connection
func (q *Queue) do(cmd string, args ...interface{}) (interface{}, error) {
	conn := q.pool.Get()
	defer conn.Close()
	return conn.Do(cmd, args...)
}
