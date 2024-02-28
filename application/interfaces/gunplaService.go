package interfaces

import (
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
)

type GunplaService interface {
	GetAllGunplas() ([]*entity.Gunpla, error)
	AddGunpla(entity.NewGunpla) (*entity.NewGunpla, error)
	UpdateGunpla(entity.Gunpla) (*entity.Gunpla, error)
}
