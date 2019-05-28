package source

import (
	"bytes"
	"io/ioutil"
	"os"
)

type Filesystem struct {
	title string
	file  []byte
	info  []os.FileInfo
}

func (f *Filesystem) getFile(dir string) {
	b, err := ioutil.ReadFile(dir)
	if err != nil {
		panic(err)
	}
	f.file = b
}

func (f *Filesystem) getDir(dir string) {
	p, err := ioutil.ReadDir(dir)

	if err != nil {
		panic(err)
	}
	f.info = p
}

func (f *Filesystem) writeFile(fileName string, b bytes.Buffer) {
	err := ioutil.WriteFile(fileName+".html", b.Bytes(), 0644)

	if err != nil {
		panic(err)
	}
}
