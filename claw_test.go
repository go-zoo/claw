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
		order += "2"
	})
	m := func(rw http.ResponseWriter, req *http.Request) {
		order += "1"
	}
	h := func(rw http.ResponseWriter, req *http.Request) {}

	c.Use(h).Add(m).ServeHTTP(response, request)

	if order != "12" {
		t.Fatalf("Wrong execution order")
	}
}

func TestSchema(t *testing.T) {
	order := ""
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	c := New()
	s := NewSchema(func(rw http.ResponseWriter, req *http.Request) {
		order += "1"
	})
	m := func(rw http.ResponseWriter, req *http.Request) {
		order += "2"
	}
	h := func(rw http.ResponseWriter, req *http.Request) {}

	c.Use(h).Add(m).Schema(s).ServeHTTP(response, request)

	if order != "12" {
		t.Fatalf("Wrong execution order")
	}
}
