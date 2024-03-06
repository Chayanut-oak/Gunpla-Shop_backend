package entity

type Comment struct {
	ForumId   string `json:"forumId"`
	CommentId string `json:"commentId"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	CreatedAt string `json:"createdAt"`
}
