package api

import (
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	res, err := Main(r.Context(), Request{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	log.Printf("Hello Morning Night Dream! %v", res)

	fmt.Fprintf(w, "<h1>Hello Morning Night Dream!</h1>")
}
