package db

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/schollz/progressbar/v3"
)

func LoadData() {
	const animeFile string = "dataset/anime-dataset-2023.csv"
	const userAnimeFile string = "dataset/final_animedataset.csv"
	const userDetailsFile string = "dataset/users-details-2023.csv"
	const userScoresFile string = "dataset/users-score-2023.csv"
	const userRatingsFile string = "dataset/user-filtered.csv"
	db := GetDB()

	fmt.Printf("📄 File: %s\n", filepath.Base(animeFile))
	err := loadConcurrently(db, animeFile, "Processing Anime...", parseAnime, insertAnime)
	if err != nil {
		fmt.Println()
		fmt.Printf("❌ Error loading anime data: %v\n", err)
	} else {
		fmt.Println()
		fmt.Println("✅ Successfully loaded anime data")
	}

	fmt.Printf("📄 File: %s\n", filepath.Base(userAnimeFile))
	err = loadConcurrently(db, userAnimeFile, "Processing User Anime...", parseUserAnime, insertUserAnime)
	if err != nil {
		fmt.Println()
		fmt.Printf("❌ Error loading user anime data: %v\n", err)
	} else {
		fmt.Println()
		fmt.Println("✅ Successfully loaded user anime data")
	}

	fmt.Printf("📄 File: %s\n", filepath.Base(userDetailsFile))
	err = loadConcurrently(db, userDetailsFile, "Processing User Details...", parseUserDetails, insertUserDetails)
	if err != nil {
		fmt.Println()
		fmt.Printf("❌ Error loading user details data: %v\n", err)
	} else {
		fmt.Println()
		fmt.Println("✅ Successfully loaded user details data")
	}

	fmt.Printf("📄 File: %s\n", filepath.Base(userScoresFile))
	err = loadConcurrently(db, userScoresFile, "Processing User Scores...", parseUserScore, insertUserScore)
	if err != nil {
		fmt.Println()
		fmt.Printf("❌ Error loading user score data: %v\n", err)
	} else {
		fmt.Println()
		fmt.Println("✅ Successfully loaded user score data")
	}

	fmt.Printf("📄 File: %s\n", filepath.Base(userRatingsFile))
	err = loadConcurrently(db, userRatingsFile, "Processing User Ratings...", parseUserRating, insertUserRating)
	if err != nil {
		fmt.Println()
		fmt.Printf("❌ Error loading user ratings data: %v\n", err)
	} else {
		fmt.Println()
		fmt.Println("✅ Successfully loaded user ratings data")
	}
}

func loadConcurrently[T any](
	db *sql.DB,
	filePath string,
	desc string,
	parse func(records []string) (T, error),
	insert func(db *sql.DB, toInsert []T) error,
) error {
	const workerCount, batchSize = 10, 1000
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	bar := progressbar.NewOptions(len(records)-1,
		progressbar.OptionSetDescription(desc),
		progressbar.OptionShowCount(),
	)

	rowsChan := make(chan T, 1000)
	var wg sync.WaitGroup

	wg.Add(workerCount)
	for w := range workerCount {
		go func(start int) {
			defer wg.Done()
			for i := start + 1; i < len(records); i += workerCount {
				record := records[i]
				parsed, err := parse(record)
				if err != nil {
					fmt.Printf("error parsing row: [%d]: %v\n", i, err)
					continue
				}
				rowsChan <- parsed
			}
		}(w)
	}

	// Close rowsChan once all parsers finish
	go func() {
		wg.Wait()
		close(rowsChan)
	}()

	// Inserter (can be main goroutine)
	var batch []T
	for row := range rowsChan {
		batch = append(batch, row)
		if len(batch) >= batchSize {
			if err := insert(db, batch); err != nil {
				log.Printf("insert error: %v", err)
			}
			bar.Add(len(batch))
			batch = batch[:0]
		}
	}
	if len(batch) > 0 {
		if err := insert(db, batch); err != nil {
			log.Printf("insert error (final): %v", err)
		}
		bar.Add(len(batch))
	}

	bar.Finish()
	return nil
}
