package compress

import (
	"archive/zip"
	"io/ioutil"
	"os"
)

type zipPackage struct {
	files []zipFile
}

type zipFile struct {
	name     string
	filename string
}

func NewZip() *zipPackage {
	return &zipPackage{}
}

func (z *zipPackage) AddFile(name string, f string) error {
	if _, err := os.Stat(f); err != nil {
		return err
	}

	z.files = append(z.files, zipFile{name: name, filename: f})
	return nil
}

func (z *zipPackage) WriteToFile(filename string) error {
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	return z.Write(out)
}

func (z *zipPackage) Write(out *os.File) error {
	zipWriter := zip.NewWriter(out)
	defer zipWriter.Close()

	for _, f := range z.files {
		fileWriter, err := zipWriter.Create(f.name)
		if err != nil {
			return err
		}

		in, err := ioutil.ReadFile(f.filename)
		if err != nil {
			return err
		}

		if _, err = fileWriter.Write(in); err != nil {
			return err
		}
	}

	return nil
}
