package model

type URLStats struct {
	ID       int64  `json:"ID"`
	ShortUrl string `json:"ShortUrl"`
	Visits   int    `json:"Visits"`
	Rank     int    `json:"Rank"`
}
