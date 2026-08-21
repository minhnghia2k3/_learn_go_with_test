package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Chris")
}

func main() {
	// Greet(os.Stdout, "Chris")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(GreetHandler)))
}
