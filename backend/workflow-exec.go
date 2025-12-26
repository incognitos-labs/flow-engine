package main

import (
	"fmt"
	"net/http"
)

func triggerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Workflow triggered")
	w.Write([]byte("OK"))
	dummyAction()
}

func dummyAction() {
	fmt.Println("Action executed")
}

func main() {
	http.HandleFunc("/trigger", triggerHandler)

	fmt.Println("Workflow Engine Server Running on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
