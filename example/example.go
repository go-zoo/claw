package main

import (
	"fmt"
	"net/http"

	"github.com/squiidz/claw"
)

func main() {
	c := claw.New(Middle1)

	http.Handle("/home", c.Use(Home))

	http.ListenAndServe(":9000", nil)
}

func Home(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Home Handler\n"))
}

func Middle1(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("FROM MIDDLEWARE 1\n")
}
