package vehicle

import "github.com/google/uuid"

type Service interface {
	CreateVehicle(vehicle *Vehicle) (*Vehicle, error)
	GetVehicle(id uuid.UUID) (*Vehicle, error)
	GetVehicles() ([]Vehicle, error)
}

type service struct {
	repository *Repository
}

func (s *service) CreateVehicle(vehicle *Vehicle) (*Vehicle, error) {
	return s.repository.CreateVehicle(vehicle)
}

func (s *service) GetVehicle(id uuid.UUID) (*Vehicle, error) {
	return s.repository.GetVehicle(id)
}

func (s *service) GetVehicles() ([]Vehicle, error) {
	return s.repository.GetVehicles()
}

// static functions

func NewService(repository *Repository) *service {
	return &service{repository: repository}
}
