package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

// This is used to understand pn what are the different ways clients can interact with our http server via requests.
func gojo(w http.ResponseWriter, r *http.Request) {
	offsetProvided := r.URL.Query().Has("offset")
	offset := r.URL.Query().Get("offset")
	limitProvided := r.URL.Query().Has("limit")
	limit := r.URL.Query().Get("limit")
	fmt.Fprintf(w, "Gojo Satoru received query params offset (%t) %s, and limit (%t) %s ", offsetProvided, offset, limitProvided, limit)
	fmt.Fprintf(w, "Gojo Satoru!")
}

func sukuna(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading the request body")
		return
	}
	fmt.Fprintf(w, "Sukuna! %s", requestBody)
}

func nanami(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	if name == "" {
		name = "Daijobu, we did not receive the name"
	}
	fmt.Fprintf(w, "Received the form value of name as %s", name)
}

func itadori(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	if name == "" {
		w.Header().Set("x-missing-field", "name")
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Received the form value of name as %s", name)
}

func main() {
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/gojo", gojo)
	serveMux.HandleFunc("/sukuna", sukuna)
	serveMux.HandleFunc("/nanami", nanami)
	serveMux.HandleFunc("/itadori", itadori)

	err := http.ListenAndServe(":8080", serveMux)
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server was closed gracefully")
		} else {
			fmt.Println("Unable to spin up a server at port 8080 or the server has probably crashed")
		}
	}
}
