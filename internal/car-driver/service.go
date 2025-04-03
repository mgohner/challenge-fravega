package carDriver

import "github.com/google/uuid"

type Service interface {
	CreateDriver(driver *Driver) (*Driver, error)
	GetDriver(id uuid.UUID) (*Driver, error)
	GetDrivers() ([]Driver, error)
}

type service struct {
	repository *Repository
}

func (s *service) CreateDriver(driver *Driver) (*Driver, error) {
	return s.repository.CreateDriver(driver)
}

func (s *service) GetDriver(id uuid.UUID) (*Driver, error) {
	return s.repository.GetDriver(id)
}

func (s *service) GetDrivers() ([]Driver, error) {
	return s.repository.GetDrivers()
}

// static functions

func NewService(repository *Repository) *service {
	return &service{repository: repository}
}
