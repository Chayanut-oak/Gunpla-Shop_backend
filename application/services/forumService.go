package services

import (
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/repository"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"
)

type ForumService struct {
	forumService repository.ForumRepository
}

func CreateForumService(forumService repository.ForumRepository) *ForumService {
	return &ForumService{forumService: forumService}
}

func (s *ForumService) GetAllForums() ([]*entity.Forum, error) {
	return s.forumService.GetAllForums()
}
func (s *ForumService) AddForum(forum restModel.ForumRestModel) (*restModel.ForumRestModel, error) {
	return s.forumService.AddForum(forum)
}
func (s *ForumService) UpdateForum(forum entity.Forum) (*entity.Forum, error) {
	return s.forumService.UpdateForum(forum)
}
func (s *ForumService) DeleteForum(forumId string) error {
	return s.forumService.DeleteForum(forumId)
}
func (s *ForumService) GetAllCommentsInForum(forumId string) ([]*entity.Comment, error) {
	return s.forumService.GetAllCommentsInForum(forumId)
}
func (s *ForumService) GetForum(forumId string) (*entity.Forum, error) {
	return s.forumService.GetForum(forumId)
}
func (s *ForumService) AddComment(comment restModel.CommentRestModel) (*restModel.CommentRestModel, error) {
	return s.forumService.AddComment(comment)
}
func (s *ForumService) UpdateComment(comment entity.Comment) (*entity.Comment, error) {
	return s.forumService.UpdateComment(comment)
}
func (s *ForumService) DeleteComment(commentId string) error {
	return s.forumService.DeleteComment(commentId)
}
