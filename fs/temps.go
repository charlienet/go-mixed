package fs

import (
	"io/ioutil"
	"os"
)

func TempFileWithText(text string) (*os.File, error) {
	return TempFile([]byte(text))
}

func TempFilenameWithText(text string) (string, error) {
	tmpfile, err := TempFileWithText(text)
	if err != nil {
		return "", err
	}

	filename := tmpfile.Name()
	if err = tmpfile.Close(); err != nil {
		return "", err
	}

	return filename, nil
}

func TempFile(data []byte) (*os.File, error) {
	tmpfile, err := ioutil.TempFile(os.TempDir(), "tmp*")
	if err != nil {
		return nil, err
	}

	if err := ioutil.WriteFile(tmpfile.Name(), data, os.ModeTemporary); err != nil {
		return nil, err
	}

	return tmpfile, nil
}
