package http

import "net/http"

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
