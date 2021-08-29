package app

import (
	"html/template"
	"net/http"

	"github.com/worbridg/miniblog/domain"
)

type Context interface {
	Header(key, value string)
	Status(code int)
	Template(t *template.Template, data interface{})
	IndexPage(blogs map[int]domain.Blog)
	ErrorPage(err error)
	HTML(file []byte)
	File(filename string)
	BlogPage(blog domain.Blog)
	BlogID() (int, error)
	Method() string
	CreateBlog() (domain.Blog, error)
	ResponseWriter() http.ResponseWriter
	Request() *http.Request
	Redirect(url string)
	EditPage(id int, blog domain.Blog)
}
