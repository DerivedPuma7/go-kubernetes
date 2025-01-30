package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/configmap", ConfigMap)
	http.ListenAndServe(":3000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	
	fmt.Fprintf(w, "Hello! I'm %s, I'm %s", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("/app/myFamily/family.txt")
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}
	fmt.Fprintf(w, "My family: %s", string(data))
}