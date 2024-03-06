package repository

import (
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"
)

type ForumRepository interface {
	GetAllForums() ([]*entity.Forum, error)
	AddForum(restModel.ForumRestModel) (*restModel.ForumRestModel, error)
	UpdateForum(entity.Forum) (*entity.Forum, error)
	DeleteForum(string) error
	GetAllCommentsInForum(string) ([]*entity.Comment, error)
	GetForum(string) (*entity.Forum, error)
	AddComment(restModel.CommentRestModel) (*restModel.CommentRestModel, error)
	UpdateComment(entity.Comment) (*entity.Comment, error)
	DeleteComment(string) error
}
