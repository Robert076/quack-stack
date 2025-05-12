package main

import (
	"encoding/json"
	"log"
	"net/http"

	db "github.com/Robert076/quack-stack.git/internal/database"
)

func main() {
	http.HandleFunc("/ducks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid method. Only GET method allowed to this endpoint.", http.StatusMethodNotAllowed)
			log.Print("Invalid method. Only GET method allowed to this endpoint.")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		ducks, err := db.GetDucksFromDatabase()

		if err != nil {
			http.Error(w, "Error retrieving ducks from database", http.StatusInternalServerError)
			log.Print("Error retrieving ducks from database: ", err)
			return
		}

		if err := json.NewEncoder(w).Encode(ducks); err != nil {
			http.Error(w, "Error encoding ducks", http.StatusInternalServerError)
			log.Print("Error encoding ducks: ", err)
			return
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Print("Error starting http server: ", err)
		return
	}
}
