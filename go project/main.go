package main

import (
	"fmt"

	"github.com/vatsal278/go-project/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()

	fmt.Println("Welcome")
	log.Fatal(http.ListenAndServe(":8000", r))
}
