package db

import (
	"database/sql"
	"log"
	"strconv"
)

func parseUserRating(record []string) (UserRating, error) {
	userID, err := strconv.Atoi(record[0])
	if err != nil {
		return UserRating{}, err
	}
	animeID, err := strconv.Atoi(record[1])
	if err != nil {
		return UserRating{}, err
	}
	rating, err := strconv.ParseFloat(record[2], 64)
	if err != nil {
		return UserRating{}, err
	}
	return UserRating{
		UserID:  userID,
		AnimeID: animeID,
		Rating:  rating,
	}, nil
}

func insertUserRating(db *sql.DB, toInsert []UserRating) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(`
		INSERT INTO user_ratings (user_id, anime_id, rating)
		VALUES (?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, row := range toInsert {
		_, err := stmt.Exec(row.UserID, row.AnimeID, row.Rating)
		if err != nil {
			log.Println("insert error:", err)
		}
	}

	return tx.Commit()
}
