/**********************************
***  Middleware Chaining in Go  ***
***  Code is under MIT license  ***
***    Code by CodingFerret     ***
*** 	github.com/squiidz      ***
***********************************/

package claw

import (
	"fmt"
	"net/http"
	"reflect"
)

type ClawHandler struct {
	http.Handler
}

func NewHandler(h http.Handler) *ClawHandler {
	return &ClawHandler{h}
}

// PodFunc redefine http.HandlerFunc
type ClawFunc func(rw http.ResponseWriter, req *http.Request)

func (c ClawFunc) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	c(rw, req)
}

// Middleware is the signature of a valid middleware with Pod
type MiddleWare func(http.Handler) http.Handler

// Pod is the array of the Global Middleware
type Claw struct {
	Handlers []MiddleWare
}

// NewPod create a new empty Pod
func New(m ...interface{}) *Claw {
	c := &Claw{}
	if m != nil {
		c.wrap(m)
	}
	return c
}

func (c *Claw) Glob(ms []MiddleWare) {
	for _, s := range ms {
		c.Handlers = append(c.Handlers, s)
	}
}

// wrap add some Global middleware to the Pod.Handlers array
func (c *Claw) wrap(m []interface{}) {
	stack := toMiddleware(m)
	for _, s := range stack {
		c.Handlers = append(c.Handlers, s)
	}
}

// Use, merge all the global middleware with the provided http.HandlerFunc
func (c *Claw) Use(h http.HandlerFunc) *ClawHandler {
	if len(p.Handlers) > 0 {
		var stack ClawHandler
		for i, m := range c.Handlers {
			switch i {
			case 0:
				stack = m(h)
			default:
				stack = m(stack)
			}
		}
		return NewHandler(stack)
	}
	return NewHandler(ClawFunc(h))
}

// Add some middleware to a particular handler
func (c *ClawHandler) Add(m ...interface{}) http.Handler {
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
	return n
}

// Schema takes a Schema type variable and use it on the PodFunc who call the function.
func (c *ClawHandler) Schema(sc ...*Schema) *ClawHandler {
	var t http.Handler
	for _, s := range sc {
		for _, m := range *s {
			t = m(c)
		}
	}

	return NewHandler(t)
}

// Mutate generate a valid handler with a provided http.HandlerFunc
func mutate(h http.HandlerFunc) MiddleWare {
	return func(next http.Handler) http.Handler {
		return PodFunc(func(rw http.ResponseWriter, req *http.Request) {
			h(rw, req)
			next.ServeHTTP(rw, req)
		})
	}
}

// Get the interface type and transform to MiddleWare type.
func toMiddleware(m []interface{}) []MiddleWare {
	var stack []MiddleWare
	if len(m) > 0 {
		for _, f := range m {
			switch v := f.(type) {
			case func(http.ResponseWriter, *http.Request):
				stack = append(stack, mutate(http.HandlerFunc(v)))
			case func(http.Handler) http.Handler:
				stack = append(stack, v)
			default:
				fmt.Println("[x] [", reflect.TypeOf(v), "] is not a valid MiddleWare Type.")
			}
		}
	}
	return stack
}
