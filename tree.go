package gopress

import (
	"errors"
	"net/http"
	"strings"
)

var (
	ErrNotFound = errors.New("not Found: 404")
	ErrMethodNotAllowed = errors.New("method Not Allowed: 405")
	Http404Response = []byte("page not found")
	Http405Response = []byte("method not allowed")
)

const (
	pathRoot string = "/"
)

type Tree struct {
	node *node
}

type Middleware func(next http.Handler) http.Handler

type node struct {
	label    string
	actions  map[string]*action
	children map[string]*node  
	middlewares []Middleware
}

type action struct {
	Handler http.Handler
}

type result struct {
	Actions *action
	Middlewares []Middleware
}

func NewResult() *result {
	return &result{}
}

func NewTree() *Tree {
	return &Tree{
		node: &node{
			label: pathRoot,
			actions: make(map[string]*action),
			children: make(map[string]*node),
			middlewares: make([]Middleware, 0),
		},
	}
}

func (t *Tree) Insert(methods []string, path string, handler http.Handler, middlewares ...Middleware) {
    curNode := t.node
    if path == pathRoot {
        curNode.label = path
        for _, method := range methods {
            curNode.actions[method] = &action{
                Handler: handler,
            }
        }
		curNode.middlewares = append(curNode.middlewares, middlewares...)
    }
	ep := explodePath(path)

	for i, p := range ep {
		nextNode, ok := curNode.children[p]
		if ok {
			curNode = nextNode
		}

		if !ok {
			curNode.children[p] = &node{
				label: p,
				actions: make(map[string]*action),
				children: make(map[string]*node),
			}
			curNode = curNode.children[p]
		}

		if i == len(ep) - 1 {
			curNode.label = p
			for _, method := range methods {
				curNode.actions[method] = &action{
					Handler: handler,
				}
			}
			curNode.middlewares = append(curNode.middlewares, middlewares...)
			break
		}
	}
}

func explodePath(path string) []string {
	s := strings.Split(path, "/")
	var r []string
	for _, str := range s {
		if str != "" {
			r  = append(r, str)
		}
	}
	return r
}

func (t *Tree) Search(method string, path string) (*result, error) {
	result := NewResult()
	curNode := t.node
	if path != pathRoot {
		for _, p := range explodePath(path) {
			nextNode, ok := curNode.children[p]
			if !ok {
				if p == curNode.label {
					break
				} else {
					return nil, ErrNotFound
				}
			}
			curNode = nextNode
			continue
		}
	}
	result.Actions = curNode.actions[method]
	result.Middlewares = append(result.Middlewares, curNode.middlewares...)
	if result.Actions == nil {
		return nil, ErrMethodNotAllowed
	}
	return result, nil
}