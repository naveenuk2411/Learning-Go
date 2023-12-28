package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func sayHi(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request in sayHi handler\n")
	io.WriteString(w, "Hi from server")
}

func sayYowaiMo(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request in sayYowaiMo handler\n")
	io.WriteString(w, "Yowai mo!!")
}

type myCustomType struct{}

func (h myCustomType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Yowai Mo! from custom")
}

func gojo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Yowai Mo!!")
}

const keyServerAddress = "serverAddr"

func getGojo(w http.ResponseWriter, r *http.Request) {
	requestContext := r.Context()
	fmt.Printf("Got request from the server %s\n", requestContext.Value(keyServerAddress))
	fmt.Fprintf(w, "Gojo Satoru!")
}

func getSukuna(w http.ResponseWriter, r *http.Request) {
	requestContext := r.Context()
	fmt.Printf("Got request from the server %s\n", requestContext.Value(keyServerAddress))
	fmt.Fprintf(w, "Sukuna!")
}

func main() {
	// http.HandleFunc("/", sayHi)
	// http.HandleFunc("/gojo", http.HandlerFunc(gojo))
	// http.HandleFunc("/sukuna", gojo)
	serveMux := http.NewServeMux()
	// serveMux.Handle("/gojo", http.HandlerFunc(gojo))
	// serveMux.HandleFunc("/sukuna", gojo)
	// serveMux.Handle("/custom", myCustomType{})
	serveMux.HandleFunc("/gojo", getGojo)
	serveMux.HandleFunc("/sukuna", getSukuna)

	ctx, cancel := context.WithCancel(context.Background())

	server1 := &http.Server{
		Addr:    ":8080",
		Handler: serveMux,
		// BaseContext is a way to change parts of the context.Context that handler functions
		// receive when they call the Context method of *http.Request.
		// In the case of your program, you’re adding the address the server is listening on (l.Addr().String()) to the context
		// with the key serverAddr, which will then be printed to the handler function’s output.
		BaseContext: func(listener net.Listener) context.Context {
			ctx := context.WithValue(ctx, keyServerAddress, listener.Addr().String())
			return ctx
		},
	}

	server2 := &http.Server{
		Addr:    ":8081",
		Handler: serveMux,
		BaseContext: func(listener net.Listener) context.Context {
			ctx := context.WithValue(ctx, keyServerAddress, listener.Addr().String())
			return ctx
		},
	}
	go func() {
		defer cancel()
		// ListenAndServer() is a blocking function call, until the server is shut down by the user or crashes for somer reason/error.
		// Hence the goroutine will continue to work, until the server crashes.
		// If the server crashes or is stopped, we cancel the context so that the context passed to the request for the handlers registered
		// to the router of this server also get cancelled.
		err := server1.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				fmt.Println("Server has been shut down")
			} else {
				fmt.Println("Failed to bring http server up and running on port 8080")
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
	}()

	go func() {
		defer cancel()
		// ListenAndServer() is a blocking function call, until the server is shut down by the user or crashes for somer reason/error.
		// Hence the goroutine will continue to work, until the server crashes.
		// If the server crashes or is stopped, we cancel the context so that the context passed to the request for the handlers registered
		// to the router of this server also get cancelled.
		err := server2.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				fmt.Println("Server has been shut down")
			} else {
				fmt.Println("Failed to bring http server up and running on port 8081")
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
	}()

	// Do not allow the main routine to complete it's execution instead we can make it dependent on the our context's done channel.
	// Hence if any of the server dies, then ctx done's channel will be closed, and hence our main routine will terminate.
	// Hence with our code, if any of the server dies, our program / main routine will exit.
	<-ctx.Done()
}
