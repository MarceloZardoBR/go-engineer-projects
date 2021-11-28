package viewmodels

import "time"

type Titles struct {
	Titles []Title `json:"titles"`
}

// Title holds movie struct
type Title struct {
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	Duration    string    `json:"duration"`
	ListedIn    string    `json:"listed_in"`
	Cast        string    `json:"cast"`
	DataAdded   time.Time `json:"data_added"`
	Rating      string    `json:"rating"`
	ReleaseYear int64     `json:"release_year"`
	Director    string    `json:"director"`
	ShowID      string    `json:"show_id"`
	Country     string    `json:"country"`
}
