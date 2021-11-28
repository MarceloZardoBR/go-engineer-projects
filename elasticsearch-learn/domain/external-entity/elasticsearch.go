package externalentity

import "time"

type ElasticSearchBody struct {
	Query Query `json:"query"`
}

// Query holds entity values
type Query struct {
	MultiMatch MultiMatch `json:"multi_match"`
}

// MultiMatch holds entity values
type MultiMatch struct {
	Query  string   `json:"query"`
	Fields []string `json:"fields"`
}

// ElasticSearchResponse holds body response struct
type ElasticSearchResponse struct {
	Hits Hits `json:"hits"`
}

type Hits struct {
	HitsValues []HitsValues `json:"hits"`
}

type HitsValues struct {
	Source Source `json:"_source"`
}

type Source struct {
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
