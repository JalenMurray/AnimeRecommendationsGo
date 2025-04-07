package db

import (
	"database/sql"
	"log"
	"strconv"
)

func parseUserScore(record []string) (UserScore, error) {
	userID, err := strconv.Atoi(record[0])
	if err != nil {
		return UserScore{}, err
	}
	animeID, err := strconv.Atoi(record[2])
	if err != nil {
		return UserScore{}, err
	}
	rating, err := strconv.ParseFloat(record[4], 64)
	if err != nil {
		return UserScore{}, err
	}
	return UserScore{
		UserID:     userID,
		Username:   record[1],
		AnimeID:    animeID,
		AnimeTitle: record[3],
		Rating:     rating,
	}, nil
}

func insertUserScore(db *sql.DB, toInsert []UserScore) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(`
		INSERT INTO user_scores (user_id, username, anime_id, anime_title, rating)
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, row := range toInsert {
		_, err := stmt.Exec(row.UserID, row.Username, row.AnimeID, row.AnimeTitle, row.Rating)
		if err != nil {
			log.Println("insert error:", err)
		}
	}

	return tx.Commit()
}
