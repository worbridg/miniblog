package main

import (
	"log"
	"net/http"

	"github.com/worbridg/miniblog/adapter/context"
	"github.com/worbridg/miniblog/app"
	"github.com/worbridg/miniblog/infra/blogdb"
	"github.com/worbridg/miniblog/infra/fileutil"
)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(
			"%s %s %s [remote_addr:%s]",
			r.Method,
			r.URL,
			r.Proto,
			r.RemoteAddr,
		)
		next.ServeHTTP(w, r)
	})
}

func main() {
	c := app.NewBlogController(blogdb.NewBlogDB())
	handler := http.NewServeMux()
	f := fileutil.NewFileUtil()
	handler.HandleFunc("/", context.ConnectToHTTPHandler(c.RootHandler, f))
	handler.HandleFunc("/blog", context.ConnectToHTTPHandler(c.BlogHandler, f))
	handler.HandleFunc("/post", context.ConnectToHTTPHandler(c.PostHandler, f))
	handler.HandleFunc("/edit", context.ConnectToHTTPHandler(c.EditHandler, f))
	handler.HandleFunc("/delete", context.ConnectToHTTPHandler(c.DeleteHandler, f))

	server := &http.Server{
		Addr:    ":9090",
		Handler: logger(handler),
	}
	server.ListenAndServe()
}
