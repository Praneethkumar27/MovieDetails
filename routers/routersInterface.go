package routers
type routers interface{

	pong(c *gin.Context)
	CreateMovie(c *gin.Context)
	GetMovieByIdByParam(c *gin.Context)
	GetMovieByIdByQuery(c *gin.Context)
	DeleteMovie(c *gin.Context)
	GetMovies(c *gin.Context)

}