package main

import (
	"fmt"
	"net/http"

	"github.com/squiidz/bone"
	"github.com/squiidz/claw"
	"github.com/squiidz/claw/mw"
)

func main() {
	mux := bone.New()
	c := claw.New(mw.Logger)
	stk := claw.NewStack(Middle1, Middle2)

	mux.Handle("/home", c.Use(Home).Stack(stk))

	http.ListenAndServe(":8080", mux)
}

func Home(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Home Handler\n"))
}

func Middle1(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("FROM MIDDLEWARE 1\n")
}

func Middle2(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("FROM MIDDLEWARE 2\n")
}

func Useless(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("I'M A COMPLETLY USELESS MIDDLEWARE\n")
}
