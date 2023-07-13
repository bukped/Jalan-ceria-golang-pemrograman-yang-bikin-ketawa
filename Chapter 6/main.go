package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Halo, dunia!")
}
func main() {
	// Menetapkan rute dan handler
	http.HandleFunc("/", homeHandler)
	// Memulai server HTTP pada port yang ditentukan
	http.ListenAndServe(":8000", nil)
}
