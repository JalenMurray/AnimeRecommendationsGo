package db

import (
	"database/sql"
	"log"
	"strconv"
	"strings"
)

func parseAnime(row []string) (AnimeRow, error) {
	animeID, _ := strconv.Atoi(row[0])
	score, _ := strconv.ParseFloat(row[4], 64)
	episodes, _ := strconv.Atoi(strings.Split(row[8], ".")[0]) // "26.0" → 26
	rank, _ := strconv.Atoi(row[18])
	popularity, _ := strconv.Atoi(row[19])
	favorites, _ := strconv.Atoi(row[20])
	scoredBy, _ := strconv.Atoi(row[21])
	members, _ := strconv.Atoi(row[22])

	return AnimeRow{
		AnimeID:     animeID,
		Name:        row[1],
		EnglishName: row[2],
		OtherName:   row[3],
		Score:       score,
		Genres:      row[5],
		Synopsis:    row[6],
		Type:        row[7],
		Episodes:    episodes,
		Aired:       row[9],
		Premiered:   row[10],
		Status:      row[11],
		Producers:   row[12],
		Licensors:   row[13],
		Studios:     row[14],
		Source:      row[15],
		Duration:    row[16],
		Rating:      row[17],
		Rank:        rank,
		Popularity:  popularity,
		Favorites:   favorites,
		ScoredBy:    scoredBy,
		Members:     members,
		ImageURL:    row[23],
	}, nil
}

func insertAnime(db *sql.DB, rows []AnimeRow) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(`
		INSERT INTO anime (
			anime_id, name, english_name, other_name, score, genres, synopsis,
			type, episodes, aired, premiered, status, producers, licensors, studios,
			source, duration, rating, rank, popularity, favorites, scored_by,
			members, image_url
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, row := range rows {
		_, err := stmt.Exec(
			row.AnimeID, row.Name, row.EnglishName, row.OtherName, row.Score, row.Genres, row.Synopsis,
			row.Type, row.Episodes, row.Aired, row.Premiered, row.Status, row.Producers, row.Licensors,
			row.Studios, row.Source, row.Duration, row.Rating, row.Rank, row.Popularity, row.Favorites,
			row.ScoredBy, row.Members, row.ImageURL,
		)
		if err != nil {
			log.Println("insert error:", err)
		}
	}

	return tx.Commit()
}

func GetAnimeByID(id int) (*Anime, error) {
	db := GetDB()
	query := `SELECT anime_id, name, english_name, other_name, score, genres, synopsis, type, episodes, aired, premiered, status, producers, licensors, studios, source, duration, rating, rank, popularity, favorites, scored_by, members, image_url FROM anime WHERE anime_id = ?`
	row := db.QueryRow(query, id)

	var a Anime
	err := row.Scan(
		&a.AnimeID, &a.Name, &a.EnglishName, &a.OtherName, &a.Score,
		&a.Genres, &a.Synopsis, &a.Type, &a.Episodes, &a.Aired,
		&a.Premiered, &a.Status, &a.Producers, &a.Licensors,
		&a.Studios, &a.Source, &a.Duration, &a.Rating, &a.Rank,
		&a.Popularity, &a.Favorites, &a.ScoredBy, &a.Members, &a.ImageURL,
	)
	if err != nil {
		return nil, err
	}
	return &a, nil
}
