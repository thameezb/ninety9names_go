package repository

import (
	"time"

	"github.com/jmoiron/sqlx"

	//pg lib
	_ "github.com/lib/pq"
	"github.com/thameezb/ninety9names/src/models"
)

type Interface interface {
	InsertName(name models.Name) error
	GetName(id string) (models.Name, error)

	GetAll() ([]models.Name, error)
	DeleteAll() error
}

type repoImpl struct {
	db *sqlx.DB
}

func New(dbURL string) Interface {
	db := sqlx.MustConnect("postgres", dbURL)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(5 * time.Minute)
	return &repoImpl{db}
}

func (r *repoImpl) InsertName(name models.Name) error {
	_, err := r.db.NamedExec(`
	INSERT INTO names 
		VALUES (:id, :arabic, :transliteration, :meaning_shaykh, :meaning_general, :explanation)
	`, &name)
	return err
}

func (r *repoImpl) GetName(id string) (models.Name, error) {
	var name models.Name
	err := r.db.Get(&name, `SELECT * FROM names WHERE id = $1`, id)
	return name, err
}

func (r *repoImpl) GetAll() ([]models.Name, error) {
	var names []models.Name
	err := r.db.Select(&names, `SELECT * FROM names`)
	return names, err
}

func (r *repoImpl) DeleteAll() error {
	_, err := r.db.Exec(`DELETE FROM names WHERE 1=1`)
	return err
}
