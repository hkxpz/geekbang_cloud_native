package main

import (
	"log"
	"net/http"
	"net/http/pprof"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexFunc)
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/debug/pprof", pprof.Index)
	mux.HandleFunc("/debug/Profile", pprof.Profile)
	mux.HandleFunc("/debug/Symbol", pprof.Symbol)
	mux.HandleFunc("/debug/Trace", pprof.Trace)
	if err := http.ListenAndServe(":80", mux); err != nil {
		return
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("server alive")); err != nil {
	}
	return
}

func indexFunc(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	for k, v := range r.Header {
		header.Set(k, v[0])
	}
	header.Set("Version", os.Getenv("VERSION"))
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("hello world")); err != nil {
	}

	log.Printf("IP:%s, HTTP Response Code:%d\n", r.RemoteAddr, http.StatusOK)

	return
}
