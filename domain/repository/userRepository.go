package repository

import (
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"
)

type UserRepository interface {
	NewUser(restModel.UserRestModel) (string, error)
	AuthenticateUser(string, string) (*entity.User, error)
	GetUserByEmail(string) (*entity.User, error)
}
