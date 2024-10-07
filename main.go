package main

import (
	"fmt"
	useCases "myproject/Desktop/GoCompleteProject/businessLayer"
	"myproject/Desktop/GoCompleteProject/domain"
)

func main() {
	mServices := useCases.NewMovieServices()
	movie := &domain.Movie{
		Title:       "Avatar",
		Description: "Directed by james Bond",
		Year:        2008,
	}
	err := mServices.CreateMovie(movie)
	if err != nil {
		fmt.Println("error creating on the movie", err)
	}
	fmt.Println("movie created sucessfully", movie)
	getmovie, err := mServices.GetMovie(movie.ID)
	if err != nil {
		fmt.Errorf("the error getting movie ", err)

	}
	fmt.Println(getmovie)
}
