package main

import (
	"fmt"
	"log"
	"net/http"
)
type Person struct {
    name    string
    address string
}

var people []Person

func formHandler (w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	person := Person{name: name, address: address}

	people = append(people, person)

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
	fmt.Fprintf(w, "hello!")
}

func peopleHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/people"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}

	for i := 0; i < len(people); i++ {
        fmt.Println(people[i].name + " " + people[i].address + "\n")
		fmt.Fprintf(w, people[i].name + " " + people[i].address + "\n")
    }
}


func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/people", peopleHandler)

	fmt.Printf("Starting server on port 3001\n")
	if err:= http.ListenAndServe(":3001", nil); err != nil {
		log.Fatal(err);
	}
}