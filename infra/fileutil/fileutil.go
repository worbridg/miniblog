package fileutil

import (
	"html/template"
	"io/ioutil"
)

type FileUtil struct{}

func NewFileUtil() *FileUtil {
	return &FileUtil{}
}

func (f *FileUtil) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func (f *FileUtil) ReadTemplate(filename string) (*template.Template, error) {
	return template.ParseFiles(filename)
}
