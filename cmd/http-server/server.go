package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/echo", echoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echoHandler(w http.ResponseWriter, req *http.Request) {
	writeRequestInfo(w, req)

	defer req.Body.Close()
	b, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintln(w, "An error occurred")
		return
	}

	fmt.Fprintln(w, string(b))
}

func writeRequestInfo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%s %s\n\n", req.Method, req.URL)
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%s: %s\n", name, h)
		}
	}
	fmt.Fprintln(w, "")
}
