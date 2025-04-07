package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/JalenMurray/AnimeRecommendationsGo/db"
)

func GetAnimeByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not ALlowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/anime/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid anime ID", http.StatusBadRequest)
		return
	}

	anime, err := db.GetAnimeByID(id)
	if err != nil {
		http.Error(w, "Anime not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(anime)
}

func GetAnimeByQuery(w http.ResponseWriter, r *http.Request) {
	database := db.GetDB()
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	q := r.URL.Query()
	query := "SELECT * FROM anime WHERE 1=1"
	args := []interface{}{}

	if name := q.Get("name"); name != "" {
		query += " AND name LIKE ?"
		args = append(args, "%"+name+"%")
	}

	if genreList := q.Get("genres"); genreList != "" {
		genres := strings.Split(genreList, ",")
		for _, g := range genres {
			query += " AND genres LIKE ?"
			args = append(args, "%"+g+"%")
		}
	}

	addNumericFilter := func(field string) {
		if val := q.Get(field); val != "" {
			query += " AND " + field + " = ?"
			args = append(args, val)
		}
		if val := q.Get(field + "_gt"); val != "" {
			query += " AND " + field + " > ?"
			args = append(args, val)
		}
		if val := q.Get(field + "_lt"); val != "" {
			query += " AND " + field + " < ?"
			args = append(args, val)
		}
	}

	addNumericFilter("score")
	addNumericFilter("episodes")
	addNumericFilter("popularity")

	if typ := q.Get("type"); typ != "" {
		query += " AND type = ?"
		args = append(args, typ)
	}

	rows, err := database.Query(query, args...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []db.Anime
	for rows.Next() {
		var a db.Anime
		err := rows.Scan(
			&a.AnimeID, &a.Name, &a.EnglishName, &a.OtherName, &a.Score,
			&a.Genres, &a.Synopsis, &a.Type, &a.Episodes, &a.Aired,
			&a.Premiered, &a.Status, &a.Producers, &a.Licensors,
			&a.Studios, &a.Source, &a.Duration, &a.Rating, &a.Rank,
			&a.Popularity, &a.Favorites, &a.ScoredBy, &a.Members, &a.ImageURL,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		results = append(results, a)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
