package claw

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOrder(t *testing.T) {
	order := ""
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	c := New(func(rw http.ResponseWriter, req *http.Request) {
		order += "b"
	})
	m := func(rw http.ResponseWriter, req *http.Request) {
		order += "a"
	}
	h := func(rw http.ResponseWriter, req *http.Request) {}

	c.Use(h).Add(m).ServeHTTP(response, request)

	if order != "ab" {
		t.Fatalf("Wrong execution order")
	}
}

func TestStack(t *testing.T) {
	order := ""
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	c := New()
	s := NewStack(func(rw http.ResponseWriter, req *http.Request) {
		order += "a"
	})
	m := func(rw http.ResponseWriter, req *http.Request) {
		order += "b"
	}
	h := func(rw http.ResponseWriter, req *http.Request) {}

	c.Use(h).Add(m).Stack(s).ServeHTTP(response, request)

	if order != "ab" {
		t.Fatalf("Wrong execution order")
	}
}

func TestSurcharge(t *testing.T) {
	result := ""
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	c := New()
	s := NewStack()
	m := func(rw http.ResponseWriter, req *http.Request) {
		result += "b"
	}
	for i := 0; i < 20; i++ {
		s = append(s, mutate(m))
	}

	h := func(rw http.ResponseWriter, req *http.Request) {
	}

	c.Use(h).Stack(s).ServeHTTP(response, request)

	if len(result) != 20 {
		t.Fatalf("Doesn't run every middleware")
	}
}
