package db

import (
	"database/sql"
	"strconv"
)

func parseUserDetails(record []string) (UserDetail, error) {
	malID, _ := strconv.Atoi(record[0])
	daysWatched, _ := strconv.ParseFloat(record[6], 64)
	meanScore, _ := strconv.ParseFloat(record[7], 64)
	parseInt := func(s string) int {
		val, _ := strconv.Atoi(s)
		return val
	}
	return UserDetail{
		MalID:           malID,
		Username:        record[1],
		Gender:          record[2],
		Birthday:        record[3],
		Location:        record[4],
		Joined:          record[5],
		DaysWatched:     daysWatched,
		MeanScore:       meanScore,
		Watching:        parseInt(record[8]),
		Completed:       parseInt(record[9]),
		OnHold:          parseInt(record[10]),
		Dropped:         parseInt(record[11]),
		PlanToWatch:     parseInt(record[12]),
		TotalEntries:    parseInt(record[13]),
		Rewatched:       parseInt(record[14]),
		EpisodesWatched: parseInt(record[15]),
	}, nil
}

func insertUserDetails(db *sql.DB, toInsert []UserDetail) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(`
		INSERT INTO user_details (
			mal_id, username, gender, birthday, location, joined,
			days_watched, mean_score, watching, completed, on_hold,
			dropped, plan_to_watch, total_entries, rewatched, episodes_watched
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, r := range toInsert {
		_, err := stmt.Exec(
			r.MalID, r.Username, r.Gender, r.Birthday, r.Location, r.Joined,
			r.DaysWatched, r.MeanScore, r.Watching, r.Completed, r.OnHold,
			r.Dropped, r.PlanToWatch, r.TotalEntries, r.Rewatched, r.EpisodesWatched,
		)
		if err != nil {
			return err
		}
	}
	tx.Commit()
	return nil
}
