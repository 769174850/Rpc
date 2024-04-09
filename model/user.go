package model

type User struct {
	ID       int64  `json:"ID"`
	Username string `json:"UserName"`
	Password string `json:"PassWord"`
}
