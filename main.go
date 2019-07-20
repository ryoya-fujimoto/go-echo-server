package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Echo server listening on port %s.\n", port)
	err := http.ListenAndServe(":"+port, http.HandlerFunc(handler))
	if err != nil {
		panic(err)
	}
}

func handler(wr http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	fmt.Printf("%s | %s %s | time: %s\n", req.RemoteAddr, req.Method, req.URL, time.Now().String())

	for key, values := range req.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}

	buf := &bytes.Buffer{}
	buf.ReadFrom(req.Body)
	fmt.Println(string(buf.Bytes()))

	serveHTTP(wr, req)
}

func serveHTTP(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Add("Content-Type", "text/plain")
	wr.WriteHeader(200)

	fmt.Fprintln(wr, "")
	io.Copy(wr, req.Body)
}
