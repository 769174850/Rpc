package model

import "time"

type Url struct {
	ID        int64     `json:"ID"`
	ShortUrl  string    `json:"ShortUrl"`
	LongUrl   string    `json:"LongUrl"`
	Visits    int       `json:"Visits"`
	UserID    int64     `json:"UserID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
