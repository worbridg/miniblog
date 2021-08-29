package domain

import "html/template"

type FileUtil interface {
	ReadFile(filename string) ([]byte, error)
	ReadTemplate(filename string) (*template.Template, error)
}
