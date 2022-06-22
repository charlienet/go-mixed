package fs

import (
	"io"
	"net/http"
	"os"
)

func GetMimeByFileName(filename string) string {
	f, err := os.Open(filename)
	if err != nil {

	}

	defer f.Close()
	return GetMimeByStream(f)
}

func GetMimeByStream(fp io.Reader) string {
	buf := make([]byte, 32)
	fp.Read(buf)

	return http.DetectContentType(buf)
}

func GetMimeByBytes(b []byte) string {
	return http.DetectContentType(b)
}
