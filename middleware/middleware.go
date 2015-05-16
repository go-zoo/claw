/**********************************
***  Middleware Chaining in Go  ***
***  Code is under MIT license  ***
***    Code by CodingFerret     ***
*** 	github.com/squiidz      ***
***********************************/

package middleware

import (
	"compress/gzip"
	"io"
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
	logger = log.New(os.Stdout, "||CLAW|| ", 2)
)

func NewLogger(out io.Writer, prefix string, flag int) func(http.Handler) http.Handler {
	logg := log.New(out, prefix+" ", flag)
	return func(next http.Handler) http.Handler {
		p := runtime.GOOS
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			switch req.Method {
			case "GET":
				if p != "windows" {
					output(logg, GET, req)
				} else {
					logg.Printf("[%s] %s %s", req.Method, req.RemoteAddr, req.RequestURI)
				}
			case "POST":
				if p != "windows" {
					output(logg, POST, req)
				} else {
					logg.Printf("[%s] %s %s", req.Method, req.RemoteAddr, req.RequestURI)
				}
			case "DELETE":
				if p != "windows" {
					output(logg, DELETE, req)
				} else {
					logg.Printf("[%s] %s %s", req.Method, req.RemoteAddr, req.RequestURI)
				}
			}
			next.ServeHTTP(rw, req)
		})
	}
}

// Very simple Console Logger
func Logger(next http.Handler) http.Handler {
	p := runtime.GOOS
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			if p != "windows" {
				output(logger, GET, req)
			} else {
				logger.Printf("[%s] %s %s", req.Method, req.RemoteAddr, req.RequestURI)
			}
		case "POST":
			if p != "windows" {
				output(logger, POST, req)
			} else {
				logger.Printf("[%s] %s %s", req.Method, req.RemoteAddr, req.RequestURI)
			}
		case "DELETE":
			if p != "windows" {
				output(logger, DELETE, req)
			} else {
				logger.Printf("[%s] %s %s", req.Method, req.RemoteAddr, req.RequestURI)
			}
		}
		next.ServeHTTP(rw, req)
	})
}

// Set the color
func output(log *log.Logger, meth string, req *http.Request) {
	log.Printf("\x1b[%s[%s]\x1b[0m %s %s", meth, req.Method, req.RemoteAddr, req.RequestURI)
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

type zipResponse struct {
	io.Writer
	http.ResponseWriter
}

func (z zipResponse) Write(b []byte) (int, error) {
	if z.Header().Get("Content-Type") == "" {
		z.Header().Set("Content-Type", http.DetectContentType(b))
	}
	return z.Writer.Write(b)
}

// Compressing Middleware
func Zipper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Encoding", "gzip")

		crw := gzip.NewWriter(rw)
		defer crw.Close()

		zrw := zipResponse{Writer: crw, ResponseWriter: rw}
		next.ServeHTTP(zrw, req)
	})
}
