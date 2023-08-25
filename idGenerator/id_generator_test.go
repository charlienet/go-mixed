package idgenerator_test

import (
	"testing"

	idgenerator "github.com/charlienet/go-mixed/idGenerator"
)

func TestNextId(t *testing.T) {
	generator := idgenerator.New(idgenerator.Config{})

	for i := 0; i < 10; i++ {
		t.Log(generator.Next().Id())
	}
}

func TestBatch(t *testing.T) {
	generator := idgenerator.New(idgenerator.Config{})

	b := generator.Batch(20)
	for _, i := range b {
		t.Log(i)
	}
}

func TestNext(t *testing.T) {
	generator := idgenerator.New(idgenerator.Config{})

	for i := 0; i < 10; i++ {
		t.Log(generator.Next())
	}
}
