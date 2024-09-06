package gopress

import (
	"errors"
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

type Middleware func(next Handler) Handler

type node struct {
	label    	string
	isParam  	bool
	paramName 	string
	actions  	map[string]*action
	children 	map[string]*node  
	middlewares []Middleware
}

type action struct {
	Handler Handler
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

func (t *Tree) Insert(methods []string, path string, handler Handler, middlewares ...Middleware) {
    curNode := t.node
	if curNode == nil {
		panic("Tree node is nil")
	}
	
    if path == pathRoot {
        curNode.label = path
		if curNode.actions == nil {
			curNode.actions = make(map[string]*action)
		}
        for _, method := range methods {
            curNode.actions[method] = &action{
                Handler: handler,
            }
        }
		curNode.middlewares = append(curNode.middlewares, middlewares...)
		return
    }
	ep := explodePath(path)

	for i, p := range ep {
		isParam := strings.HasPrefix(p, ":")
		paramName := ""
		if isParam {
			paramName = p[1:]
			p = ":"
		}
		
		nextNode, ok := curNode.children[p]		
		if !ok {
			if curNode.children == nil {
				curNode.children = make(map[string]*node)
			}
			curNode.children[p] = &node{
				label: p,
				isParam: isParam,
				paramName: paramName,
				actions: make(map[string]*action),
				children: make(map[string]*node),
				middlewares: make([]Middleware, 0),
			}
			curNode = curNode.children[p]
		} else {
			curNode = nextNode
		}

		if i == len(ep) - 1 {
			curNode.label = p
			if curNode.actions == nil {
				curNode.actions = make(map[string]*action)
			}
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

func (t *Tree) Search(method string, path string) (*result, map[string]string, error) {
	result := NewResult()
	curNode := t.node
	params := make(map[string]string)

	if path != pathRoot {
		for _, p := range explodePath(path) {
			nextNode, ok := curNode.children[p]
			if !ok {
				for _, childNode := range curNode.children {
					if childNode.isParam {
						nextNode = childNode
						params[childNode.paramName] = p
						break
					}
				}
				if nextNode == nil {
					return nil, nil, ErrNotFound
				}
			}
			curNode = nextNode
		}
	}
	result.Actions = curNode.actions[method]
	result.Middlewares = append(result.Middlewares, curNode.middlewares...)
	if result.Actions == nil {
		return nil, nil, ErrMethodNotAllowed
	}
	return result, params, nil
}