package middleware

import (
	"compress/gzip"
	"net/http"
)

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
