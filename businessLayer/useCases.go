package useCases

import (
	"myproject/Desktop/GoCompleteProject/domain"
	database "myproject/Desktop/GoCompleteProject/infrastructure"

	"github.com/google/uuid"
)

type MovieServices interface {
	GetMovie(id string) (*domain.Movie, error) // get movie foocus on the return values of the struct
	GetMovies() ([]domain.Movie, error)
	CreateMovie(movie *domain.Movie) error // create movie focus on the struct so we are passing the struct as an argument
	UpdateMovie(movie domain.Movie) error
	DeleteMovie(id string) error
}
type movieServices struct {
	db database.MovieMap
}

func NewMovieServices() MovieServices {
	mapsDb := database.NewMovieMap()

	return &movieServices{db: *mapsDb}

}
func (m movieServices) GetMovie(id string) (*domain.Movie, error) {
	movie := m.db.GetId(id)
	return &movie, nil
}
func (m movieServices) GetMovies() ([]domain.Movie, error) {
	return m.db.Get(), nil
}
func (m movieServices) CreateMovie(movie *domain.Movie) error {
	id := uuid.New()
	movie.ID = id.String()
	err := m.db.CreateMovie(movie)

	return err
}
func (m movieServices) UpdateMovie(movie domain.Movie) error {
	return nil
}
func (m movieServices) DeleteMovie(id string) {

}
