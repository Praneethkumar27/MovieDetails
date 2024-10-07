package database

import "myproject/Desktop/GoCompleteProject/domain"

type MovieMap struct {
	mapsMovieDb map[string]domain.Movie
}

func NewMovieMap() *MovieMap {
	movieMap := make(map[string]domain.Movie)
	return &MovieMap{
		mapsMovieDb: movieMap,
	}
}
func (m *MovieMap) Get() []domain.Movie {
	var movies []domain.Movie
	for _, movie := range m.mapsMovieDb {
		movies = movie.append(movies, movie)
	}

}
func (m *MovieMap) GetId(id string) domain.Movie {
	return m.mapsMovieDb[id]

}
func (m *MovieMap) CreateMovie(movie *domain.Movie) error {
	m.mapsMovieDb[movie.ID] = *movie
	return nil

}
