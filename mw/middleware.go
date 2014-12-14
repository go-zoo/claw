/**********************************
***  Middleware Chaining in Go  ***
***  Code is under MIT license  ***
***    Code by CodingFerret     ***
*** 	github.com/squiidz      ***
***********************************/

package mw

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"runtime/debug"
)

const (
	DELETE = "41m"
	GET    = "42m"
	POST   = "44m"
)

var (
	LOG = log.New(os.Stdout, "||CLAW|| ", 2)
)

// Very simple Console Logger
func Logger(next http.Handler) http.Handler {
	p := runtime.GOOS
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			if p != "windows" {
				output(GET, req)
			} else {
				LOG.Printf("[%s] %s %s", req.Method, req.RemoteAddr, req.RequestURI)
			}
		case "POST":
			if p != "windows" {
				output(POST, req)
			} else {
				LOG.Printf("[%s] %s %s", req.Method, req.RemoteAddr, req.RequestURI)
			}
		case "DELETE":
			if p != "windows" {
				output(DELETE, req)
			} else {
				LOG.Printf("[%s] %s %s", req.Method, req.RemoteAddr, req.RequestURI)
			}
		}
		next.ServeHTTP(rw, req)
	})
}

// Set the color
func output(meth string, req *http.Request) {
	LOG.Printf("\x1b[%s[%s]\x1b[0m %s %s", meth, req.Method, req.RemoteAddr, req.RequestURI)
}

// Recovery Middleware
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				stack := debug.Stack()
				log.Printf("PANIC: %s\n%s", err, stack)

			}
		}()
		next.ServeHTTP(rw, req)
	})
}
