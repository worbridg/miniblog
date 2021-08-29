package context

import (
	"errors"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/worbridg/miniblog/domain"
)

type HTTPContext struct {
	w    http.ResponseWriter
	r    *http.Request
	f    domain.FileUtil
	path string
}

func NewHTTPContext(
	w http.ResponseWriter,
	r *http.Request,
	f domain.FileUtil,
	path string,
) *HTTPContext {
	return &HTTPContext{
		w:    w,
		r:    r,
		f:    f,
		path: path,
	}
}

func (c *HTTPContext) Header(key, value string) {
	c.w.Header().Set(key, value)
}

func (c *HTTPContext) Status(code int) {
	c.w.WriteHeader(code)
}

func (c *HTTPContext) Template(t *template.Template, data interface{}) {
	c.Header("Content-Type", "text/html; charset=UTF-8")
	c.Status(http.StatusOK)
	t.Execute(c.w, data)
}

func (c *HTTPContext) IndexPage(blogs map[int]domain.Blog) {
	tpl, err := c.f.ReadTemplate(filepath.Join(c.path, "index.html.tpl"))
	if err != nil {
		c.ErrorPage(errors.New("missing index template"))
		return
	}
	c.Template(tpl, struct{ Blogs map[int]domain.Blog }{Blogs: blogs})
}

func (c *HTTPContext) ErrorPage(err error) {
	c.Header("Content-Type", "text/plain; charset=UTF-8")
	c.Status(http.StatusInternalServerError)
	c.w.Write([]byte(err.Error()))
}

func (c *HTTPContext) HTML(file []byte) {
	c.Header("Content-Type", "text/html; charset=UTF-8")
	c.Status(http.StatusOK)
	c.w.Write(file)
}

func (c *HTTPContext) File(filename string) {
	file, err := c.f.ReadFile(filepath.Join(c.path, filename))
	if err != nil {
		c.ErrorPage(errors.New("missing " + filename))
	} else {
		c.HTML(file)
	}
}

func (c *HTTPContext) BlogPage(blog domain.Blog) {
	tpl, err := c.f.ReadTemplate(filepath.Join(c.path, "blog.html.tpl"))
	if err != nil {
		c.ErrorPage(errors.New("missing blog template"))
		return
	}
	c.Template(tpl, blog)
}

func (c *HTTPContext) BlogID() (int, error) {
	return strconv.Atoi(c.r.URL.Query().Get("id"))
}

func (c *HTTPContext) Method() string {
	return c.r.Method
}

func (c *HTTPContext) CreateBlog() (domain.Blog, error) {
	if err := c.r.ParseForm(); err != nil {
		return domain.Blog{}, errors.New("posted data error")
	}
	return domain.NewBlog(
		c.r.Form.Get("blog-title"),
		c.r.Form.Get("blog-content"),
	)
}

func (c *HTTPContext) ResponseWriter() http.ResponseWriter {
	return c.w
}

func (c *HTTPContext) Request() *http.Request {
	return c.r
}

func (c *HTTPContext) Redirect(url string) {
	http.Redirect(c.w, c.r, url, http.StatusMovedPermanently)

}

func (c *HTTPContext) EditPage(id int, blog domain.Blog) {
	tpl, err := c.f.ReadTemplate(filepath.Join(c.path, "edit.html.tpl"))
	if err != nil {
		c.ErrorPage(errors.New("missing edit template"))
		return
	}
	c.Template(tpl, struct {
		ID   int
		Blog domain.Blog
	}{id, blog})
}
