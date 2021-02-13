package types

type ChatMessage struct {
	ID      string `json:"id"`
	Author  string `json:"author"`
	Content string `json:"content"`
	Stamp   int64  `json:"stamp"`
}
