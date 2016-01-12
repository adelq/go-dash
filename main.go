package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hostname", systemHandler(hostname))
	http.HandleFunc("/time", systemHandler(time))
	http.HandleFunc("/issue", systemStructHandler(issue))
	http.HandleFunc("/mem", systemHandler(mem))
	http.HandleFunc("/lastlog", systemHandler(lastlog))
	http.HandleFunc("/swap", systemHandler(swap))
	http.HandleFunc("/df", systemStructHandler(df))
	http.HandleFunc("/arp", systemStructHandler(arp))
	http.HandleFunc("/load", systemHandler(load))
	http.HandleFunc("/netstat", systemHandler(netstat))
	http.HandleFunc("/uptime", systemHandler(uptime))
	fmt.Println("Serving on port 3000...")
	http.ListenAndServe(":3000", nil)
}

type system func() (string, error)
type systemStruct interface{}
type systemStructFn func() (systemStruct, error)

func systemHandler(fn system) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := fn()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(result))
	}
}

func systemStructHandler(fn systemStructFn) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := fn()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		buf, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(buf)
	}
}
