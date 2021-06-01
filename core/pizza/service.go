package pizza

import (
	"errors"
	"gorm.io/gorm"
)

type UseCase interface {
	GetAll() ([]*Pizza, error)
	GetByName(name string) (*Pizza, error)
	Store(o *Pizza) error
	Update(o *Pizza) error
	Remove(ID int64) error
}

type Service struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
func (s *Service) GetAll() ([]*Pizza, error) {
	var pizza []*Pizza
	s.DB.Table("pizza").Take(&pizza)

	return pizza, nil
}

func (s *Service) GetByName(name string) (*Pizza, error) {
	var pizza *Pizza

	s.DB.Table("pizza").First(&pizza, "name = ?", name)

	return pizza, nil
}

func (s *Service) Store(o *Pizza) error {
	pizza := Pizza{Name: o.Name, Ingredients: o.Ingredients, Price: o.Price}
	result := s.DB.Table("pizza").Create(&pizza) // pass pointer of data to Create
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Service) Update(o *Pizza) error {
	panic("implement me")
}

func (s *Service) Remove(ID int64) error {
	if ID == 0 {
		return errors.New("ID é requerido")
	}
	s.DB.Table("pizza").Delete(&Pizza{}, ID)
	return nil
}
