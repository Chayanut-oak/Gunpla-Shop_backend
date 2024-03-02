package interfaces

import (
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/restModel"
)

type GunplaService interface {
	GetAllGunplas() ([]*entity.Gunpla, error)
	AddGunpla(restModel.GunplaRestModal) (*restModel.GunplaRestModal, error)
	UpdateGunpla(entity.Gunpla) (*entity.Gunpla, error)
	// UpdateGunplaStock([]valueObject.Product) (string, error)
	DeleteGunpla(string) error
}
