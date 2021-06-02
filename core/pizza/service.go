package pizza

import (
	"database/sql"
	"fmt"
	"log"
)

type UseCase interface {
	GetAll() ([]*Pizza, error)
	GetByName(name string) (Pizza, error)
	Store(o *Pizza) error
	Update(ID int64, o *Pizza) error
	Remove(ID int64) error
}

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}
func (s *Service) GetAll() ([]*Pizza, error) {
	var result []*Pizza
	rows, err := s.DB.Query("SELECT * FROM pizza")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p Pizza
		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Ingredients)
		if err != nil {
			return nil, err
		}
		result = append(result, &p)
	}
	return result, nil
}

func (s *Service) GetByName(name string) (Pizza, error) {
	var pizza Pizza
	sqlStatement := `SELECT * FROM pizza where name = $1`
	err := s.DB.QueryRow(sqlStatement, name).Scan(&pizza.ID, &pizza.Name, &pizza.Price, &pizza.Ingredients)
	if err != nil {
		log.Fatal(err)
	}
	return pizza, nil
}

func (s *Service) Store(o *Pizza) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into pizza(name, price, ingredients) values ($1,$2,$3)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(o.Name, o.Price, o.Ingredients)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *Service) Update(ID int64, o *Pizza) error {
	if ID == 0 {
		return fmt.Errorf("invalid ID")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("update pizza set name=$1, price=$2, ingredients=$3 where id=$4")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(o.Name, o.Price, o.Ingredients, ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *Service) Remove(ID int64) error {
	if ID == 0 {
		return fmt.Errorf("invalid ID")
	}
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from pizza where id=?", ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
