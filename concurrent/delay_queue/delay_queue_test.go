package delayqueue_test

import (
	"testing"
	"time"

	delayqueue "github.com/charlienet/go-mixed/concurrent/delay_queue"
)

type delayTask struct {
	message string
	delay   time.Time
}

func (t delayTask) Delay() time.Time {
	return t.delay
}

func TestDelayQueue(t *testing.T) {
	queue := delayqueue.New[delayTask]()
	queue.Push(delayTask{})
}

func TestDelayedFunc(t *testing.T) {
	q := delayqueue.New[delayTask]()
	q.Push(delayTask{})
}

func TestDelayedChannel(t *testing.T) {
	q := delayqueue.New[delayTask]()
	c := q.Channel(10)

	q.Push(delayTask{message: "abc", delay: time.Now().Add(time.Second)})
	q.Push(delayTask{message: "abcaaa", delay: time.Now().Add(time.Second * 3)})

	for {
		if q.IsEmpty() {
			t.Log("队列为空,退出")
			break
		}

		select {
		case task := <-c:
			t.Log(task)
		case <-time.After(time.Second * 2):
		}
	}
}
