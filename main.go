package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hostname", hostnameHandler)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/issue", issueHandler)
	http.HandleFunc("/mem", memHandler)
	http.HandleFunc("/lastlog", lastlogHandler)
	http.HandleFunc("/swap", swapHandler)
	http.HandleFunc("/df", dfHandler)
	http.HandleFunc("/arp", arpHandler)
	http.HandleFunc("/load", loadHandler)
	http.HandleFunc("/netstat", netstatHandler)
	http.HandleFunc("/uptime", uptimeHandler)
	fmt.Println("Serving on port 3000...")
	http.ListenAndServe(":3000", nil)
}

func hostnameHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := hostname()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(hostname))
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	time, err := time()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(time))
}

func issueHandler(w http.ResponseWriter, r *http.Request) {
	distro, kernel, err := issue()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(distro + " " + kernel))
}

func memHandler(w http.ResponseWriter, r *http.Request) {
	mem, err := mem()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(mem))
}

func lastlogHandler(w http.ResponseWriter, r *http.Request) {
	lastlog, err := lastlog()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(lastlog))
}

func swapHandler(w http.ResponseWriter, r *http.Request) {
	swap, err := swap()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(swap))
}

func dfHandler(w http.ResponseWriter, r *http.Request) {
	df, err := df()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(df))
}

func arpHandler(w http.ResponseWriter, r *http.Request) {
	arp, err := arp()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(arp))
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	load, err := load()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(load))
}

func netstatHandler(w http.ResponseWriter, r *http.Request) {
	netstat, err := netstat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(netstat))
}

func uptimeHandler(w http.ResponseWriter, r *http.Request) {
	uptime, err := uptime()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(uptime))
}
