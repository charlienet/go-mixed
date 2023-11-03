package delayqueue

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/charlienet/go-mixed/calendar"
	"github.com/charlienet/go-mixed/rand"
)

type delayTask struct {
	Message string
	At      time.Time
}

func (t delayTask) Delay() time.Time {
	return t.At
}

func (t delayTask) execute() {
	println(t.Message)
}

// var _ encoding.BinaryMarshaler = new(myStruct)
// var _ encoding.BinaryUnmarshaler = new(myStruct)

func (t delayTask) BinaryUnmarshaler(data []byte, v any) {
	json.Unmarshal(data, v)
}

func (t delayTask) MarshalBinary() (data []byte, err error) {
	return json.Marshal(t)
}

func TestMemStore(t *testing.T) {
	s := newMemStore[delayTask]()

	for i := 0; i < 10; i++ {
		s.Push(
			context.Background(),
			delayTask{
				Message: "tesss",
				At:      time.Now().Add(-time.Minute * time.Duration(rand.Intn(20))),
			})
	}

	t.Log("count:", s.Len())

	v, exists := s.Peek()
	t.Logf("Peek %v:%v %v", exists, v.Message, calendar.Create(v.Delay()).ToDateTimeString())

	for i := 0; i < 10; i++ {
		v, _ := s.Pop()
		t.Logf("POP:%v %v", v.Message, calendar.Create(v.Delay()).ToDateTimeString())
	}

	v, exists = s.Peek()
	t.Logf("Peek %v:%v %v", exists, v.Message, calendar.Create(v.At).ToDateTimeString())
}

func TestMemPush(t *testing.T) {
	s := newMemStore[delayTask]()

	for i := 0; i < 10; i++ {
		s.Push(
			context.Background(),
			delayTask{
				Message: fmt.Sprintf("abc:%d", i),
				At:      time.Now().Add(time.Second * time.Duration(rand.IntRange(5, 30))),
			})
	}

	now := time.Now()

	delay, _ := s.Pop()
	after := delay.Delay().Sub(now)

	t.Log("after:", calendar.String(now), calendar.String(delay.Delay()), after)
}

func TestExecute(t *testing.T) {
	s := newMemStore[delayTask]()

	s.Push(context.Background(),
		delayTask{
			Message: "这是消息",
			At:      time.Now().Add(time.Second * 2),
		})

	s.Push(context.Background(),
		delayTask{
			Message: "这是消息",
			At:      time.Now().Add(time.Second * 4),
		})

	t.Log("start:", calendar.String(time.Now()))

	for {
		if s.IsEmpty() {
			break
		}

		task, _ := s.Pop()

		for {
			if task.Delay().Before(time.Now()) {
				task.execute()
				t.Log("end:", calendar.String(time.Now()))
				break
			}

			time.Sleep(time.Millisecond * 20)
		}
	}

}
