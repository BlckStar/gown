package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	flag.Parse()
	http.HandleFunc("/home", serveHome)
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		h := hub.newHub()
		h.run()
	})

	http.HandleFunc("/ws/:id", func(w http.ResponseWriter, r *http.Request) {
		serveWs(h, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
