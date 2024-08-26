package gopress

import "net/http"

type RouterGroup struct {
	prefix     string
	Middleware []Middleware
	router     *Router
}

func (r *Router) Group(prefix string, middlewares ...Middleware) *RouterGroup {
	return &RouterGroup{
		prefix:     prefix,
		Middleware: middlewares,
		router:     r,
	}
}

func (g *RouterGroup) Get(path string, handler http.Handler, middleware ...Middleware) {
	fipath := g.prefix + path
	allMidlleware := append(g.Middleware, middleware...)
	g.router.Get(fipath, handler, allMidlleware...)

}

func (g *RouterGroup) Post(path string, handler http.Handler, middleware ...Middleware) {
	fipath := g.prefix + path
	allMidlleware := append(g.Middleware, middleware...)
	g.router.Post(fipath, handler, allMidlleware...)

}

func (g *RouterGroup) Delete(path string, handler http.Handler, middleware ...Middleware) {
	fipath := g.prefix + path
	allMidlleware := append(g.Middleware, middleware...)
	g.router.Delete(fipath, handler, allMidlleware...)

}

func (g *RouterGroup) Patch(path string, handler http.Handler, middleware ...Middleware) {
	fipath := g.prefix + path
	allMidlleware := append(g.Middleware, middleware...)
	g.router.Patch(fipath, handler, allMidlleware...)

}

func (g *RouterGroup) Put(path string, handler http.Handler, middleware ...Middleware) {
	fipath := g.prefix + path
	allMidlleware := append(g.Middleware, middleware...)
	g.router.Put(fipath, handler, allMidlleware...)

}

func (g *RouterGroup) Head(path string, handler http.Handler, middleware ...Middleware) {
	fipath := g.prefix + path
	allMidlleware := append(g.Middleware, middleware...)
	g.router.Head(fipath, handler, allMidlleware...)

}