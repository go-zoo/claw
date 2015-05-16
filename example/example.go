package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-zoo/claw"
	mw "github.com/go-zoo/claw/middleware"
)

func main() {
	mux := http.NewServeMux()

	logger := mw.NewLogger(os.Stdout, "[Example]", 2)

	c := claw.New(logger)
	stk := claw.NewStack(Middle1, Middle2)

	mux.HandleFunc("/home", Home)

	http.ListenAndServe(":8080", c.Merge(mux).Stack(stk))
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
