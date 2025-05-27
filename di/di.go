package di

import (
	"fmt"
	"io"
	// "log"
	// "net/http"
	// "os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// Print to stdout.
// func main() {
// 	Greet(os.Stdout, "Elodie")
// }

// MyGreeterHandler says Hello, world over HTTP.
// func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
// 	Greet(w, "world")
// }

// func main() {
// 	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
// }
