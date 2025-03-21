package main // package declaration
import (
	"fmt"      // import statement
	"net/http" // import statement
)

func handler(w http.ResponseWriter, r *http.Request) { // function declaration	
	fmt.Fprintf(w, "Hello, World!") // function call
}
func main() { 
	http.HandleFunc("/", handler) // function call
	http.ListenAndServe(":3000", nil) // function call
}