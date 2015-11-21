package main

import "net/http"

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/test", TestHandler)

	http.ListenAndServe(":8000", r)
}
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Do whatever you'd like here - e.g.
	w.Write([]byte("Hello World"))
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}
