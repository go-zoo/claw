package claw

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test the ns/op
func BenchmarkPlainMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(Bench))

	for n := 0; n < b.N; n++ {
		mux.ServeHTTP(response, request)
	}
}

func BenchmarkClawMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	mux := http.NewServeMux()
	clw := New(Useless)

	mux.Handle("/", clw.Use(Bench))

	for n := 0; n < b.N; n++ {
		mux.ServeHTTP(response, request)
	}
}

func BenchmarkClawAddMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	mux := http.NewServeMux()
	clw := New()

	mux.Handle("/", clw.Use(Bench).Add(Useless))

	for n := 0; n < b.N; n++ {
		mux.ServeHTTP(response, request)
	}
}

func BenchmarkClawStackMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	mux := http.NewServeMux()
	clw := New()
	stk := NewStack(Useless)

	mux.Handle("/", clw.Use(Bench).Stack(stk))

	for n := 0; n < b.N; n++ {
		mux.ServeHTTP(response, request)
	}
}

func BenchmarkClawFullMux(b *testing.B) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	mux := http.NewServeMux()
	clw := New()
	stk := NewStack(Useless)

	mux.Handle("/", clw.Use(Bench).Stack(stk).Add(Useless))

	for n := 0; n < b.N; n++ {
		mux.ServeHTTP(response, request)
	}
}

func Bench(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("b"))
}

// Useless Middlware
func Useless(rw http.ResponseWriter, req *http.Request) {
}
