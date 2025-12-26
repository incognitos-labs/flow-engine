package main

import (
	"fmt"
	"net/http"
)

func triggerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Workflow triggered")
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/trigger", triggerHandler)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Workflow Engine Server Running!")
}
