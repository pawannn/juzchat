package models

type Message struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Text      string `json:"text"`
	Timestamp int64  `json:"timestamp"`
	UserID    string `json:"userId"`
}
