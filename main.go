package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/JalenMurray/AnimeRecommendationsGo/db"
	"github.com/JalenMurray/AnimeRecommendationsGo/handlers"
)

func main() {
	db.InitSchema()

	if len(os.Args) >= 2 {
		command := os.Args[1]
		switch command {
		case "load-data":
			db.LoadData()
			fmt.Println("\n\n✅ Successfully loaded all data!")
			os.Exit(0)
		default:
			fmt.Printf("❌ Unknown command: %s\n\n", command)
			os.Exit(1)
		}
	}

	http.HandleFunc("/anime/", handlers.GetAnimeByID)
	http.HandleFunc("/anime", handlers.GetAnimeByQuery)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
