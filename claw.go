/**********************************
***  Middleware Chaining in Go  ***
***  Code is under MIT license  ***
***    Code by CodingFerret     ***
*** 	github.com/squiidz      ***
***********************************/

package claw

import (
	"net/http"
)

// CalwHandler only wrap a http.Handler
type ClawHandler struct {
	http.Handler
}

// newHandler Generate a pointer to a ClawHandler
func newHandler(h http.Handler) *ClawHandler {
	return &ClawHandler{h}
}

// ClawFunc redefine http.HandlerFunc
type ClawFunc func(rw http.ResponseWriter, req *http.Request)

// Serve HTTP Request
func (c ClawFunc) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	c(rw, req)
}

// Middleware is the signature of a valid middleware with Claw
type MiddleWare func(http.Handler) http.Handler

// Claw is the array of the Global Middleware
type Claw struct {
	Handlers []MiddleWare
}

// New create a new empty Claw
func New(m ...interface{}) *Claw {
	c := &Claw{}
	if m != nil {
		c.Handlers = toMiddleware(m)
	}
	return c
}

// wrap add some Global middleware to the Claw.Handlers array
func (c *Claw) wrap(m []interface{}) {
	stack := toMiddleware(m)
	for _, s := range stack {
		c.Handlers = append(c.Handlers, s)
	}
}

// Use, merge all the global middleware with the provided http.HandlerFunc
func (c *Claw) Use(h http.HandlerFunc) *ClawHandler {
	if len(c.Handlers) > 0 {
		var stack ClawHandler
		for i, m := range c.Handlers {
			switch i {
			case 0:
				stack = *newHandler(m(h))
			default:
				stack = *newHandler(m(stack))
			}
		}
		return &stack
	}
	return newHandler(ClawFunc(h))
}

// Add some middleware to a particular handler
func (c *ClawHandler) Add(m ...interface{}) *ClawHandler {
	var n http.Handler
	if m != nil {
		stack := toMiddleware(m)
		for i, s := range stack {
			if i == 0 {
				n = s(c)
			} else {
				n = s(n)
			}
		}
	}
	return newHandler(n)
}

// Stack takes a Stack type variable and use it on the ClawHandler who call the function.
func (c *ClawHandler) Stack(stk ...Stack) *ClawHandler {
	var t http.Handler
	for _, s := range stk {
		for i, _ := range s {
			if i == 0 {
				t = s[(len(s)-1)-i](c)
			} else {
				t = s[(len(s)-1)-i](t)
			}
		}
	}

	return newHandler(t)
}
