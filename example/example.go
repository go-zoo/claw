package main

import (
	"fmt"
	"net/http"

	"github.com/squiidz/claw"
	"github.com/squiidz/claw/mw"
)

func main() {
	c := claw.New(mw.Logger)
	sch := claw.NewSchema(Middle1, Middle2)

	http.Handle("/home", c.Use(Home).Schema(sch).Add(Useless))

	http.ListenAndServe(":8080", nil)
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
