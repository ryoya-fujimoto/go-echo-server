package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	var bind = flag.String("bind", "localhost:8080", "Bind address")
	flag.Parse()

	fmt.Printf("Echo server listening on: %s.\n", *bind)
	err := http.ListenAndServe(*bind, http.HandlerFunc(handler))
	if err != nil {
		panic(err)
	}
}

func handler(wr http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	fmt.Printf("%s | %s %s | time: %s\n", req.RemoteAddr, req.Method, req.URL, time.Now().String())

	wr.Header().Add("Content-Type", "text/plain")
	wr.WriteHeader(200)

	headerTxt := ""
	for key, values := range req.Header {
		for _, value := range values {
			headerTxt += fmt.Sprintf("%s: %s\n", key, value)
		}
	}
	fmt.Fprintln(wr, headerTxt)

	io.Copy(wr, req.Body)

	fmt.Fprintln(wr, "")
}
