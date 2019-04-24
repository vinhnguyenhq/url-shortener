package main

import (
	"log"
	"net/http"
)

func StartServer(m MappingValues, port string) {
	redirect := func(w http.ResponseWriter, r *http.Request, url string) {
		http.Redirect(w, r, url, http.StatusSeeOther)
	}

	for path, url := range m {
		func(path, url string) {
			http.HandleFunc("/"+path, func(w http.ResponseWriter, r *http.Request) {
				redirect(w, r, url)
			})
		}(path, url)
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
