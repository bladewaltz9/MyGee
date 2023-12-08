package Gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node // root node of every request method
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *router) parsePattern(pattern string) []string {
	values := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, value := range values {
		if value != "" {
			parts = append(parts, value)
			if value[0] == '*' { // Only one * is allowed
				break
			}
		}
	}
	return parts
}

func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	parts := r.parsePattern(pattern)

	key := method + "-" + pattern
	// if there isn't the mothod, add a root node
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRouter(method string, path string) (*node, map[string]string) {
	pathParts := r.parsePattern(path)
	params := make(map[string]string)

	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	node := root.search(pathParts, 0)

	if node != nil {
		parts := r.parsePattern(node.pattern)
		for idx, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = pathParts[idx]
			} else if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(pathParts[idx:], "/")
				break
			}
		}
		return node, params
	}

	return nil, nil
}

func (r *router) handle(c *Context) {
	node, params := r.getRouter(c.Method, c.Path)

	if node != nil {
		c.Params = params
		key := c.Method + "-" + node.pattern
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 Not Found: %s\n", c.Path)
		})
	}
	c.Next()
}
