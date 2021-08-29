package app

import (
	"errors"
	"net/http"

	"github.com/worbridg/miniblog/domain"
)

var (
	ErrInvalidID      = errors.New("invalid id")
	ErrBlogIsntPosted = errors.New("the blog isn't posted")
	ErrNoPage         = errors.New("no page")
)

type BlogController struct {
	db domain.BlogDB
}

func NewBlogController(db domain.BlogDB) *BlogController {
	return &BlogController{
		db: db,
	}
}

func (c *BlogController) RootHandler(ctx Context) {
	if ctx.Request().URL.Path == "/" {
		ctx.IndexPage(c.db.Blogs())
	} else {
		ctx.ErrorPage(ErrNoPage)
	}
}

func (c *BlogController) BlogHandler(ctx Context) {
	id, err := ctx.BlogID()
	if err != nil {
		ctx.ErrorPage(ErrInvalidID)
		return
	}
	blog, ok := c.db.Get(id)
	if !ok {
		ctx.ErrorPage(ErrBlogIsntPosted)
		return
	}
	ctx.BlogPage(blog)
}

func (c *BlogController) PostHandler(ctx Context) {
	if ctx.Method() == http.MethodPost {
		c.postHandler(ctx)
	} else {
		c.postFormHandler(ctx)
	}
}

func (c *BlogController) postHandler(ctx Context) {
	blog, err := ctx.CreateBlog()
	if err != nil {
		ctx.ErrorPage(err)
		return
	}
	c.db.Store(blog)
	ctx.Redirect("/")
}

func (c *BlogController) postFormHandler(ctx Context) {
	ctx.File("post.html")
}

func (c *BlogController) EditHandler(ctx Context) {
	if ctx.Method() == http.MethodPost {
		c.editHandler(ctx)
	} else {
		c.editFormHandler(ctx)
	}
}

func (c *BlogController) editHandler(ctx Context) {
	id, err := ctx.BlogID()
	if err != nil {
		ctx.ErrorPage(ErrInvalidID)
		return
	}
	if ok := c.db.Exist(id); !ok {
		ctx.ErrorPage(ErrBlogIsntPosted)
		return
	}
	blog, err := ctx.CreateBlog()
	if err != nil {
		ctx.ErrorPage(err)
		return
	}
	c.db.Update(id, blog)
	ctx.Redirect("/")
}

func (c *BlogController) editFormHandler(ctx Context) {
	id, err := ctx.BlogID()
	if err != nil {
		ctx.ErrorPage(ErrInvalidID)
		return
	}
	blog, ok := c.db.Get(id)
	if !ok {
		ctx.ErrorPage(ErrBlogIsntPosted)
		return
	}
	ctx.EditPage(id, blog)
}

func (c *BlogController) DeleteHandler(ctx Context) {
	id, err := ctx.BlogID()
	if err != nil {
		ctx.ErrorPage(ErrInvalidID)
		return
	}
	if ok := c.db.Exist(id); !ok {
		ctx.ErrorPage(ErrBlogIsntPosted)
		return
	}
	c.db.Delete(id)
	ctx.Redirect("/")
}
