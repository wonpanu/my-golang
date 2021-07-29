package usecase

import (
	"github.com/wonpanu/my-golang/pkg/entity"
	"github.com/wonpanu/my-golang/pkg/repo"
)

type IVaccine interface {
	GetAll() ([]entity.Vaccine, error)
	GetByID(ID string) (entity.Vaccine, error)
	Create(vc entity.Vaccine) (entity.Vaccine, error)
	Update(ID string, vc entity.Vaccine) (entity.Vaccine, error)
	Delete(ID string) error
}

type Vaccine struct {
	Repo repo.IVaccine
}

func (uc Vaccine) Create(vc entity.Vaccine) (entity.Vaccine, error) {
	response, err := uc.Repo.Create(vc)
	return response, err
}

func (uc Vaccine) GetAll() ([]entity.Vaccine, error) {
	return uc.Repo.GetAll()
}

func (uc Vaccine) GetByID(ID string) (entity.Vaccine, error) {
	response, err := uc.Repo.GetByID(ID)
	return response, err
}

func (uc Vaccine) Update(ID string, vc entity.Vaccine) (entity.Vaccine, error) {
	response, err := uc.Repo.Update(ID, vc)
	return response, err
}

func (uc Vaccine) Delete(ID string) error {
	return uc.Repo.Delete(ID)
}

func NewVaccineUsecase(vaccineRepo repo.IVaccine) Vaccine {
	return Vaccine{
		Repo: vaccineRepo,
	}
}
