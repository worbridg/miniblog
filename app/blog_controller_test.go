package app

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/worbridg/miniblog/app/mock"
	"github.com/worbridg/miniblog/domain"
)

type expects struct {
	db  func(*mock.MockBlogDB)
	ctx func(*mock.MockContext)
}

func TestBlogController_RootHandler(t *testing.T) {
	tests := []struct {
		name    string
		expects expects
		path    string
	}{
		{
			name: "render",
			expects: expects{
				db: func(db *mock.MockBlogDB) {
					db.EXPECT().
						Blogs().Return(map[int]domain.Blog{})
				},
				ctx: func(ctx *mock.MockContext) {
					r, _ := http.NewRequest(http.MethodGet, "/", nil)
					ctx.EXPECT().Request().Return(r)
					ctx.EXPECT().IndexPage(map[int]domain.Blog{})
				},
			},
			path: "ui",
		},
		{
			name: "error",
			expects: expects{
				ctx: func(ctx *mock.MockContext) {
					r, _ := http.NewRequest(http.MethodGet, "/notfound", nil)
					ctx.EXPECT().Request().Return(r)
					ctx.EXPECT().ErrorPage(ErrNoPage)
				},
			},
			path: "ui",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, ctx, finish := newBlogControllerAndContextWithMock(
				t,
				tt.expects,
				tt.path,
			)
			defer finish()
			c.RootHandler(ctx)
		})
	}
}

func TestBlogController_BlogHandler(t *testing.T) {
	tests := []struct {
		name    string
		expects expects
		path    string
	}{
		{
			name: "render",
			expects: expects{
				db: func(db *mock.MockBlogDB) {
					db.EXPECT().Get(1).Return(domain.Blog{Title: "title"}, true)
				},
				ctx: func(ctx *mock.MockContext) {
					ctx.EXPECT().BlogID().Return(1, nil)
					ctx.EXPECT().BlogPage(domain.Blog{Title: "title"})
				},
			},
			path: "ui",
		},
		{
			name: "invalid",
			expects: expects{
				ctx: func(ctx *mock.MockContext) {
					ctx.EXPECT().BlogID().Return(0, errors.New("invalid id"))
					ctx.EXPECT().ErrorPage(errors.New("invalid id"))
				},
			},
			path: "ui",
		},
		{
			name: "error",
			expects: expects{
				db: func(db *mock.MockBlogDB) {
					db.EXPECT().Get(1).Return(domain.Blog{}, false)
				},
				ctx: func(ctx *mock.MockContext) {
					ctx.EXPECT().BlogID().Return(1, nil)
					ctx.EXPECT().ErrorPage(errors.New("the blog isn't posted"))
				},
			},
			path: "ui",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, ctx, finish := newBlogControllerAndContextWithMock(t, tt.expects, tt.path)
			defer finish()
			c.BlogHandler(ctx)
		})
	}
}

func TestBlogController_postHandler(t *testing.T) {
	tests := []struct {
		name    string
		expects expects
		path    string
	}{
		{
			name: "post",
			expects: expects{
				db: func(db *mock.MockBlogDB) {
					db.EXPECT().Store(domain.Blog{})
				},
				ctx: func(ctx *mock.MockContext) {
					ctx.EXPECT().
						CreateBlog().Return(domain.Blog{}, nil)
					ctx.EXPECT().Redirect("/")
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, ctx, finish := newBlogControllerAndContextWithMock(t, tt.expects, tt.path)
			defer finish()
			c.postHandler(ctx)
		})
	}

}

func TestBlogController_postFormHandler(t *testing.T) {
	tests := []struct {
		name    string
		expects expects
		path    string
	}{
		{
			name: "get",
			expects: expects{
				ctx: func(ctx *mock.MockContext) {
					ctx.EXPECT().File("post.html")
				},
			},
			path: "ui",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, ctx, finish := newBlogControllerAndContextWithMock(t, tt.expects, tt.path)
			defer finish()
			c.postFormHandler(ctx)
		})
	}
}

func newBlogControllerAndContextWithMock(
	t *testing.T,
	expects expects,
	path string,
) (*BlogController, Context, func()) {
	c, finish := newBlogControllerWithMock(t, expects.db)

	ctrl := gomock.NewController(t)
	ctx := mock.NewMockContext(ctrl)
	if expects.ctx != nil {
		expects.ctx(ctx)
	}
	return c, ctx, func() {
		finish()
		ctrl.Finish()
	}
}

func newBlogControllerWithMock(
	t *testing.T,
	db func(*mock.MockBlogDB),
) (*BlogController, func()) {
	dbCtrl := gomock.NewController(t)
	fileCtrl := gomock.NewController(t)

	c := &BlogController{
		db: mock.NewMockBlogDB(dbCtrl),
	}

	if db != nil {
		db(c.db.(*mock.MockBlogDB))
	}

	return c, func() {
		dbCtrl.Finish()
		fileCtrl.Finish()
	}
}
