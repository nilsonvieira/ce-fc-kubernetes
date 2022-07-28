package main

import "fmt"
import "net/http"
import "os"
import "io/ioutil"
import "log"

func main() {
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello i am  %s. I am %s", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/go/src/goapp/family.txt")
	if err != nil {
		log.Fatalf("Error reading file ", err)
	}
	fmt.Fprintf(w, "My Family: %s", string(data))
}