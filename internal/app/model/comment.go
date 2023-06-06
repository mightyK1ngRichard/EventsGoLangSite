package model

type Comment struct {
	ID          string `json:"id"`
	CommentText string `json:"comment_text"`
	CommentDate string `json:"comment_date"`
	EventId     string `json:"event_id"`
	UserId      string `json:"user_id"`
}
