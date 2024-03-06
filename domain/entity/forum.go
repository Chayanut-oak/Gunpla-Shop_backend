package entity

type Forum struct {
	ForumId   string `json:"forumId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	CreatedAt string `json:"createdAt"`
}
