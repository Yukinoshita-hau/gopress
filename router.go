package gopress

import (
	"context"
	"net/http"
)

type Router struct {
	tree *Tree
}

type route struct {
	methods    []string
	path       string
	handler    http.Handler
	middleware []Middleware
}

var (
	tmpRoute = &route{}
)

func NewRouter() *Router {
	return &Router{
		tree: NewTree(),
	}
}

func (r *Router) Methods(methods ...string) *Router {
	tmpRoute.methods = append(tmpRoute.methods, methods...)
	return r
}

func (r *Router) Handler(path string, handler http.Handler, middlewares ...Middleware) {
	tmpRoute.handler = handler
	tmpRoute.path = path
	tmpRoute.middleware = append(tmpRoute.middleware, middlewares...)
	r.Handle()
}

func (r *Router) Handle() {
	r.tree.Insert(tmpRoute.methods, tmpRoute.path, tmpRoute.handler, tmpRoute.middleware...)
	tmpRoute = &route{}
}

func (r *Router) createMethodHandle(path string, method string, handler http.Handler, middlewares ...Middleware) {
	r.Methods(method)
	r.Handler(path, handler, middlewares...)
	r.Handle()
}

func (r *Router) Get(path string, handler http.Handler, middleware ...Middleware) {
	r.createMethodHandle(path, http.MethodGet, handler, middleware...)
}

func (r *Router) Post(path string, handler http.Handler, middleware ...Middleware) {
	r.createMethodHandle(path, http.MethodPost, handler, middleware...)
}

func (r *Router) Delete(path string, handler http.Handler, middleware ...Middleware) {
	r.createMethodHandle(path, http.MethodDelete, handler, middleware...)
}

func (r *Router) Patch(path string, handler http.Handler, middleware ...Middleware) {
	r.createMethodHandle(path, http.MethodPatch, handler, middleware...)
}

func (r *Router) Put(path string, handler http.Handler, middleware ...Middleware) {
	r.createMethodHandle(path, http.MethodPut, handler, middleware...)
}

func (r *Router) Head(path string, handler http.Handler, middleware ...Middleware) {
	r.createMethodHandle(path, http.MethodHead, handler, middleware...)
}

func (r *Router) Option(path string, handler http.Handler, middleware ...Middleware) {
	r.createMethodHandle(path, http.MethodOptions, handler, middleware...)
}

func (r *Router) Static(pathPrefix, directory string) {
	fs := http.FileServer(http.Dir(directory))
	r.Get(pathPrefix+"/", http.StripPrefix(pathPrefix, fs))
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path
	result, params, err := r.tree.Search(method, path)
	if err != nil {
		errorHandler(w, req, err)
		return
	}

	if result == nil || result.Actions == nil || result.Actions.Handler == nil {
        http.Error(w, "Handler not found", http.StatusNotFound)
        return
    }

	finalHandler := result.Actions.Handler

	if len(params) > 0 {
		finalHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := req.Context()
			for key, value := range params {
				ctx = context.WithValue(ctx, key, value)
			}
			req = req.WithContext(ctx)
			result.Actions.Handler.ServeHTTP(w, req)
		})
	}

	for i := len(result.Middlewares) - 1; i >= 0; i-- {
		finalHandler = result.Middlewares[i](finalHandler)
	}

	finalHandler.ServeHTTP(w, req)
}

func handleErr(err error) (int, []byte) {
	var status int
	var body []byte
	switch err {
	case ErrMethodNotAllowed:
		status = http.StatusMethodNotAllowed
		body = Http405Response
	case ErrNotFound:
		status = http.StatusNotFound
		body = Http404Response
	}
	return status, body
}
