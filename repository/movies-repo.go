package repository

import (
	"github.com/matiascfgm/Go-rest-api/entity"
)

type MovieRepository interface {
	Save(movie *entity.Movie) (*entity.Movie, error)
	FindAll() ([]entity.Movie, error)
	GetMovieById(id string) (*entity.Movie, error)
	DeleteMovie(id string) (*entity.Movie, error)
	UpdateMovie(movie *entity.Movie) (*entity.Movie, error)
}