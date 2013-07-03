package main

import (
	"log"
	"net/http"
	"strconv"
)

func UVHandler(store *UVStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		qs := r.URL.Query()

		start, err := strconv.ParseUint(qs.Get("start"), 10, 64)
		if err != nil {
			log.Println(err)
			http.Error(w, "Invalid start parameter", http.StatusBadRequest)
			return
		}

		end, err := strconv.ParseUint(qs.Get("end"), 10, 64)
		if err != nil {
			log.Println(err)
			http.Error(w, "Invalid end parameter", http.StatusBadRequest)
			return
		}

		count := store.Count(start, end)
		asciiCount := []byte(strconv.FormatUint(count, 10))

		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", strconv.Itoa(len(asciiCount)))
		w.Write(asciiCount)
	}
}
