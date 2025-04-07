package db

type AnimeRow struct {
	AnimeID     int
	Name        string
	EnglishName string
	OtherName   string
	Score       float64
	Genres      string
	Synopsis    string
	Type        string
	Episodes    int
	Aired       string
	Premiered   string
	Status      string
	Producers   string
	Licensors   string
	Studios     string
	Source      string
	Duration    string
	Rating      string
	Rank        int
	Popularity  int
	Favorites   int
	ScoredBy    int
	Members     int
	ImageURL    string
}

type UserAnime struct {
	Username   string
	AnimeID    int
	MyScore    float64
	UserID     int
	Gender     string
	Title      string
	Type       string
	Source     string
	Score      float64
	ScoredBy   int
	Rank       int
	Popularity int
	Genre      string
}

type UserDetail struct {
	MalID           int
	Username        string
	Gender          string
	Birthday        string
	Location        string
	Joined          string
	DaysWatched     float64
	MeanScore       float64
	Watching        int
	Completed       int
	OnHold          int
	Dropped         int
	PlanToWatch     int
	TotalEntries    int
	Rewatched       int
	EpisodesWatched int
}

type UserScore struct {
	UserID     int
	Username   string
	AnimeID    int
	AnimeTitle string
	Rating     float64
}

type UserRating struct {
	UserID  int
	AnimeID int
	Rating  float64
}

type Anime struct {
	AnimeID     int     `json:"anime_id"`
	Name        string  `json:"name"`
	EnglishName string  `json:"english_name"`
	OtherName   string  `json:"other_name"`
	Score       float64 `json:"score"`
	Genres      string  `json:"genres"`
	Synopsis    string  `json:"synopsis"`
	Type        string  `json:"type"`
	Episodes    int     `json:"episodes"`
	Aired       string  `json:"aired"`
	Premiered   string  `json:"premiered"`
	Status      string  `json:"status"`
	Producers   string  `json:"producers"`
	Licensors   string  `json:"licensors"`
	Studios     string  `json:"studios"`
	Source      string  `json:"source"`
	Duration    string  `json:"duration"`
	Rating      string  `json:"rating"`
	Rank        int     `json:"rank"`
	Popularity  int     `json:"popularity"`
	Favorites   int     `json:"favorites"`
	ScoredBy    int     `json:"scored_by"`
	Members     int     `json:"members"`
	ImageURL    string  `json:"image_url"`
}
