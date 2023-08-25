package filestore

import (
	"errors"
	"io"
	"os"
)

func New() {
	f, err := os.Open("")

	_ = err
	_ = f

}

func Store(name string, reader io.Reader) error {

	return errors.New("abc")

}
