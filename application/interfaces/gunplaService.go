package interfaces

import (
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
)

type GunplaService interface {
	GetAllGunplas() ([]*entity.Gunpla, error)
}
