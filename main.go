package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Go Calculator API!")
	})

	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/sub", subHandler)
	http.HandleFunc("/mul", mulHandler)
	http.HandleFunc("/div", divHandler)

	fmt.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func getParams(r *http.Request) (float64, float64, error) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid parameter a")
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid parameter b")
	}

	return a, b, nil
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := getParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%f", a+b)
}

func subHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := getParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%f", a-b)
}

func mulHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := getParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%f", a*b)
}

func divHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := getParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if b == 0 {
		http.Error(w, "cannot divide by zero", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%f", a/b)
}
