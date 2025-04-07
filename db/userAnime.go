package db

import (
	"database/sql"
	"log"
	"strconv"
)

func parseUserAnime(row []string) (UserAnime, error) {
	animeID, _ := strconv.Atoi(row[1])
	myScore, _ := strconv.ParseFloat(row[2], 64)
	userID, _ := strconv.Atoi(row[3])
	score, _ := strconv.ParseFloat(row[8], 64)
	scoredBy, _ := strconv.Atoi(row[9])
	rank, _ := strconv.Atoi(row[10])
	popularity, _ := strconv.Atoi(row[11])

	return UserAnime{
		Username:   row[0],
		AnimeID:    animeID,
		MyScore:    myScore,
		UserID:     userID,
		Gender:     row[4],
		Title:      row[5],
		Type:       row[6],
		Source:     row[7],
		Score:      score,
		ScoredBy:   scoredBy,
		Rank:       rank,
		Popularity: popularity,
		Genre:      row[12],
	}, nil
}

func insertUserAnime(db *sql.DB, rows []UserAnime) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(`
		INSERT INTO user_anime (
			username, anime_id, my_score, user_id, gender, title, type,
			source, score, scored_by, rank, popularity, genre
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, row := range rows {
		_, err := stmt.Exec(
			row.Username, row.AnimeID, row.MyScore, row.UserID, row.Gender, row.Title,
			row.Type, row.Source, row.Score, row.ScoredBy, row.Rank, row.Popularity, row.Genre,
		)
		if err != nil {
			log.Println("insert error:", err)
		}
	}

	return tx.Commit()
}
