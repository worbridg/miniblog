package context

import (
	"errors"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/worbridg/miniblog/app/mock"
	"github.com/worbridg/miniblog/domain"
)

func TestHTTPContext_IndexPage(t *testing.T) {
	type fields struct {
		w    http.ResponseWriter
		r    *http.Request
		f    func(*testing.T) (*mock.MockFileUtil, func())
		path string
	}
	type args struct {
		blogs map[int]domain.Blog
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wanted string
	}{
		{
			name: "render",
			fields: fields{
				w: httptest.NewRecorder(),
				f: func(t *testing.T) (*mock.MockFileUtil, func()) {
					ctrl := gomock.NewController(t)
					f := mock.NewMockFileUtil(ctrl)
					f.EXPECT().
						ReadTemplate("ui/index.html.tpl").
						Return(template.New("test").Parse(`{{ range $id, $blog := .Blogs }}{{.Title}}{{ end }}`))
					return f, func() { ctrl.Finish() }
				},
				path: "ui",
			},
			args: args{
				blogs: map[int]domain.Blog{
					1: {Title: "title"},
				},
			},
			wanted: "title",
		},
		{
			name: "error",
			fields: fields{
				w: httptest.NewRecorder(),
				f: func(t *testing.T) (*mock.MockFileUtil, func()) {
					ctrl := gomock.NewController(t)
					f := mock.NewMockFileUtil(ctrl)
					f.EXPECT().
						ReadTemplate("ui/index.html.tpl").
						Return(nil, errors.New("error"))
					return f, func() { ctrl.Finish() }
				},
				path: "ui",
			},
			args: args{
				blogs: map[int]domain.Blog{},
			},
			wanted: "missing index template",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, finish := tt.fields.f(t)
			defer finish()

			c := &HTTPContext{
				w:    tt.fields.w,
				r:    tt.fields.r,
				f:    f,
				path: tt.fields.path,
			}
			c.IndexPage(tt.args.blogs)
			res := c.w.(*httptest.ResponseRecorder)
			if res.Body.String() != tt.wanted {
				t.Errorf("got: %v, want: %v", res.Body, tt.wanted)
			}
		})
	}
}

func TestHTTPContext_BlogID(t *testing.T) {
	type fields struct {
		w    http.ResponseWriter
		r    *http.Request
		f    domain.FileUtil
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name: "id=1",
			fields: fields{
				r: func() *http.Request {
					r, _ := http.NewRequest(http.MethodGet, "/?id=1", nil)
					return r
				}(),
			},
			want: 1,
		},
		{
			name: "id=a",
			fields: fields{
				r: func() *http.Request {
					r, _ := http.NewRequest(http.MethodGet, "/?id=a", nil)
					return r
				}(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &HTTPContext{
				w:    tt.fields.w,
				r:    tt.fields.r,
				f:    tt.fields.f,
				path: tt.fields.path,
			}
			got, err := c.BlogID()
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPContext.BlogID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HTTPContext.BlogID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPContext_ErrorPage(t *testing.T) {
	type fields struct {
		w    http.ResponseWriter
		r    *http.Request
		f    domain.FileUtil
		path string
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "error",
			fields: fields{
				w: httptest.NewRecorder(),
			},
			args: args{
				err: errors.New("error message"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &HTTPContext{
				w:    tt.fields.w,
				r:    tt.fields.r,
				f:    tt.fields.f,
				path: tt.fields.path,
			}
			c.ErrorPage(tt.args.err)
		})
	}
}

func TestHTTPContext_File(t *testing.T) {
	type fields struct {
		w    http.ResponseWriter
		r    *http.Request
		f    func(t *testing.T) (*mock.MockFileUtil, func())
		path string
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wanted string
	}{
		{
			name: "render",
			fields: fields{
				w: httptest.NewRecorder(),
				f: func(t *testing.T) (*mock.MockFileUtil, func()) {
					ctrl := gomock.NewController(t)
					f := mock.NewMockFileUtil(ctrl)
					f.EXPECT().
						ReadFile("ui/index.html.tpl").
						Return([]byte(`index`), nil)
					return f, func() { ctrl.Finish() }
				},
				path: "ui",
			},
			args: args{
				filename: "index.html.tpl",
			},
			wanted: "index",
		},
		{
			name: "error",
			fields: fields{
				w: httptest.NewRecorder(),
				f: func(t *testing.T) (*mock.MockFileUtil, func()) {
					ctrl := gomock.NewController(t)
					f := mock.NewMockFileUtil(ctrl)
					f.EXPECT().
						ReadFile("ui/index.html.tpl").
						Return(nil, errors.New(`error`))
					return f, func() { ctrl.Finish() }
				},
				path: "ui",
			},
			args: args{
				filename: "index.html.tpl",
			},
			wanted: "missing index.html.tpl",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, finish := tt.fields.f(t)
			defer finish()
			c := &HTTPContext{
				w:    tt.fields.w,
				r:    tt.fields.r,
				f:    f,
				path: tt.fields.path,
			}
			c.File(tt.args.filename)
		})
	}
}
func TestHTTPContext_BlogPage(t *testing.T) {
	type fields struct {
		w    http.ResponseWriter
		r    *http.Request
		f    func(t *testing.T) (*mock.MockFileUtil, func())
		path string
	}
	type args struct {
		blog domain.Blog
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wanted string
	}{
		{
			name: "render",
			fields: fields{
				w: httptest.NewRecorder(),
				f: func(t *testing.T) (*mock.MockFileUtil, func()) {
					ctrl := gomock.NewController(t)
					f := mock.NewMockFileUtil(ctrl)
					f.EXPECT().
						ReadTemplate("ui/blog.html.tpl").
						Return(template.New("test").Parse(`blog`))
					return f, func() { ctrl.Finish() }
				},
				path: "ui",
			},
			args: args{
				blog: domain.Blog{},
			},
			wanted: "blog",
		},
		{
			name: "error",
			fields: fields{
				w: httptest.NewRecorder(),
				f: func(t *testing.T) (*mock.MockFileUtil, func()) {
					ctrl := gomock.NewController(t)
					f := mock.NewMockFileUtil(ctrl)
					f.EXPECT().
						ReadTemplate("ui/blog.html.tpl").
						Return(nil, errors.New("error"))
					return f, func() { ctrl.Finish() }
				},
				path: "ui",
			},
			args: args{
				blog: domain.Blog{},
			},
			wanted: "missing blog template",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, finish := tt.fields.f(t)
			defer finish()
			c := &HTTPContext{
				w:    tt.fields.w,
				r:    tt.fields.r,
				f:    f,
				path: tt.fields.path,
			}
			c.BlogPage(tt.args.blog)
		})
	}
}
