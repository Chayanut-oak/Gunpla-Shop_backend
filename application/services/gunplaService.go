package services

import (
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/entity"
	"github.com/Chayanut-oak/Gunpla-Shop_backend/domain/repository"
)

type GunplaService struct {
	gunplaRepository repository.GunplaRepository
}

func CreateGunplaService(gunplaRepository repository.GunplaRepository) *GunplaService {
	return &GunplaService{gunplaRepository: gunplaRepository}
}

func (s *GunplaService) GetAllGunplas() ([]*entity.Gunpla, error) {
	return s.gunplaRepository.GetAllGunplas()
}
