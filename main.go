package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hostname", systemHandler(hostname))
	http.HandleFunc("/time", systemHandler(time))
	http.HandleFunc("/issue", issueHandler)
	http.HandleFunc("/mem", systemHandler(mem))
	http.HandleFunc("/lastlog", systemHandler(lastlog))
	http.HandleFunc("/swap", systemHandler(swap))
	http.HandleFunc("/df", systemHandler(df))
	http.HandleFunc("/arp", systemHandler(arp))
	http.HandleFunc("/load", systemHandler(load))
	http.HandleFunc("/netstat", systemHandler(netstat))
	http.HandleFunc("/uptime", systemHandler(uptime))
	fmt.Println("Serving on port 3000...")
	http.ListenAndServe(":3000", nil)
}

type system func() (string, error)

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

func issueHandler(w http.ResponseWriter, r *http.Request) {
	distro, kernel, err := issue()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(distro + " " + kernel))
}
