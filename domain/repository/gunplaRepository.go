package repository

import (
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"
)

type GunplaRepository interface {
	GetAllGunplas() ([]*entity.Gunpla, error)
	AddGunpla(restModel.GunplaRestModel) (*restModel.GunplaRestModel, error)
	UpdateGunpla(entity.Gunpla) (*entity.Gunpla, error)
	// UpdateGunplaStock([]valueObject.Product) (string, error)
	DeleteGunpla(string) error
}
