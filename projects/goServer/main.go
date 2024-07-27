package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse form() err: %v", err);
		return;
	}
	fmt.Fprintf(w, "post request successful");
	name := r.FormValue("name");
	address := r.FormValue(("address"));
	fmt.Fprintf(w, "Name is %v", name);
	fmt.Fprintf(w, "Address is %v", address);
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	fmt.Println("hello")

	fileServer := http.FileServer(http.Dir("./static"));
	http.Handle("/", fileServer);
	http.HandleFunc("/form", formHandler);
	http.HandleFunc("/hello", helloHandler);

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}