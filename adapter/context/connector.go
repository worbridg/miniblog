package context

import (
	"net/http"

	"github.com/worbridg/miniblog/app"
	"github.com/worbridg/miniblog/domain"
)

func ConnectToHTTPHandler(handler func(app.Context), f domain.FileUtil) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := NewHTTPContext(w, r, f, "ui")
		handler(ctx)
	}
}
