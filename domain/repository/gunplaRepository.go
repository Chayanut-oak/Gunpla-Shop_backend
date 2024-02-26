package repository

import (
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
)

type GunplaRepository interface {
	GetAllGunplas() ([]*entity.Gunpla, error)
}
