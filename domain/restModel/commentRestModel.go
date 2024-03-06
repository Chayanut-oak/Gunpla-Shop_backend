package restModel

type CommentRestModel struct {
	ForumId string `json:"forumId"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
