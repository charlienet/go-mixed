package panic

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestPanic(t *testing.T) {
	defer func() {
		t.Log("捕捉异常")
		if e := recover(); e != nil {
			if err, ok := e.(error); ok {
				t.Log(err.Error())
			}
			t.Log("格式化:", e)
		}
	}()

	g := NewPanicGroup(10)
	g.Go(func() {
		panic("1243")
	})

	if err := g.Wait(context.Background()); err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("这条消息可打印")
}
