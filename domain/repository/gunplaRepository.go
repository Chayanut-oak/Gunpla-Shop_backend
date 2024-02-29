package repository

import (
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"
)

type GunplaRepository interface {
	GetAllGunplas() ([]*entity.Gunpla, error)
	AddGunpla(restModel.GunplaRestModal) (*restModel.GunplaRestModal, error)
	UpdateGunpla(entity.Gunpla) (*entity.Gunpla, error)
	UpdateGunplaStock(restModel.OrderRestModal) (string, error)
	DeleteGunpla(string) error
}
