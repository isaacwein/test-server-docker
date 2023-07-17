package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

func ToJson(d any) string {
	buf := bytes.Buffer{}
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.Encode(d)
	return buf.String()
}
func ToJsonI(d any) string {
	buf := bytes.Buffer{}
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	enc.Encode(d)

	return buf.String()
}

func main() {
	var host string
	var port string
	flag.StringVar(&host, "host", "127.0.0.1", "host to listen on")
	flag.StringVar(&port, "port", "8080", "port to listen on")
	flag.Parse()

	http.HandleFunc("/", Handler)
	fmt.Printf("Server is running on port %s:%s\n", host, port)
	http.ListenAndServe(":"+port, nil)

}

func Handler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World!\n\n")
	fmt.Fprintf(w, "Request URL: %s\n\n", ToJsonI(r.URL))
	fmt.Fprintf(w, "Remote URL: %s\n\n", r.RemoteAddr)
	fmt.Fprintf(w, "Request Method: %s\n\n", r.Method)
	fmt.Fprintf(w, "Request Proto: %s\n\n", r.Proto)

	fmt.Fprintf(w, "Request Header: %s\n\n", ToJsonI(r.Header))
	reqBody, _ := io.ReadAll(r.Body)
	fmt.Fprintf(w, "Request Body: %s\n\n", string(reqBody))
}
