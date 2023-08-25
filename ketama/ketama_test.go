package ketama_test

import (
	"testing"

	"github.com/charlienet/go-mixed/ketama"
)

func TestNew(t *testing.T) {
	k := ketama.New()

	t.Logf("%+v", k)
	k.Synchronize()

	// k.Lock()

	t.Logf("%+v", k)

	t.Log(k.IsEmpty())
	t.Logf("%+v", k)
	t.Log(k.IsEmpty())
	t.Logf("%+v", k)
	k.Synchronize()

	t.Log(k.IsEmpty())
	t.Logf("%+v", k)
	t.Log(k.IsEmpty())
	t.Logf("%+v", k)
	t.Log(k.IsEmpty())

	t.Logf("%+v", k)
}
