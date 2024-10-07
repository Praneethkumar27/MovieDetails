package routers

import (
	useCases "myproject/Desktop/GoCompleteProject/businessLayer"
	"myproject/Desktop/GoCompleteProject/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

var movieServices = useCases.NewMovieServices()

func main() {
	r := gin.Default()
	r.GET("get_movie_by_id", getMoviebyid)
	r.POST("/create", createMovie)
	r.GET("/get_Movie_details", getMovieDetails)

	r.Run("8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
func getMoviebyid(c *gin.Context) {
	id := c.Query("id")

	movie, err := movieServices.GetMovie(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, movie)

}
func createMovie(c *gin.Context) {
	var movie domain.Movie
	err := c.BindJSON(&movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	err = movieServices.CreateMovie(&movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, movie)
}
func getMovieDetails(c *gin.Context) {
	id := c.Query("id")
	movies, err := movieServices.GetMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, movies)
}
