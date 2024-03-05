package interfaces

import (
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"
)

type UserService interface {
	NewUser(restModel.UserRestModel) (string, error)
	AuthenticateUser(string, string) (string, *entity.User, error)
	GetUser(string) (*entity.User, error)
}
