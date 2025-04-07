package db

import (
	"database/sql"
	"log"
	"sync"

	_ "modernc.org/sqlite"
)

var (
	instance *sql.DB
	once     sync.Once
)

func GetDB() *sql.DB {
	once.Do(func() {
		var err error
		instance, err = sql.Open("sqlite", "./db/anime.db")
		if err != nil {
			log.Fatalf("failed to open SQLite DB: %v", err)
		}
	})
	return instance
}

func InitSchema() {
	db := GetDB()
	schema := `
	CREATE TABLE IF NOT EXISTS anime (
		anime_id INTEGER PRIMARY KEY,
		name TEXT,
		english_name TEXT,
		other_name TEXT,
		score REAL,
		genres TEXT,
		synopsis TEXT,
		type TEXT,
		episodes INTEGER,
		aired TEXT,
		premiered TEXT,
		status TEXT,
		producers TEXT,
		licensors TEXT,
		studios TEXT,
		source TEXT,
		duration TEXT,
		rating TEXT,
		rank INTEGER,
		popularity INTEGER,
		favorites INTEGER,
		scored_by INTEGER,
		members INTEGER,
		image_url TEXT
	);

	CREATE TABLE IF NOT EXISTS user_anime (
		username TEXT,
		anime_id INTEGER,
		my_score REAL,
		user_id INTEGER,
		gender TEXT,
		title TEXT,
		type TEXT,
		source TEXT,
		score REAL,
		scored_by INTEGER,
		rank INTEGER,
		popularity INTEGER,
		genre TEXT
	);

	CREATE TABLE IF NOT EXISTS user_ratings (
		user_id INTEGER,
		anime_id INTEGER,
		rating REAL
	);

	CREATE TABLE IF NOT EXISTS user_details (
		mal_id INTEGER PRIMARY KEY,
		username TEXT,
		gender TEXT,
		birthday TEXT,
		location TEXT,
		joined TEXT,
		days_watched REAL,
		mean_score REAL,
		watching INTEGER,
		completed INTEGER,
		on_hold INTEGER,
		dropped INTEGER,
		plan_to_watch INTEGER,
		total_entries INTEGER,
		rewatched INTEGER,
		episodes_watched INTEGER
	);

	CREATE TABLE IF NOT EXISTS user_scores (
		user_id INTEGER,
		username TEXT,
		anime_id INTEGER,
		anime_title TEXT,
		rating REAL
	);
	`

	if _, err := db.Exec(schema); err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}
}
