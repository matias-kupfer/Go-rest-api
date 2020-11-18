package main

import (
	"fmt"
	"github.com/matiascfgm/Go-rest-api/controller"
	router "github.com/matiascfgm/Go-rest-api/http"
	"github.com/matiascfgm/Go-rest-api/service"
	"net/http"
	"os"
)

var (
	movieService    service.MovieService       = service.NewMovieService(nil)
	movieController controller.MovieController = controller.NewMovieController(movieService)
	r               router.Router              = router.NewMuxRouter()
)

func main() {
	fmt.Println(os.Getenv("$PORT"))
	r.GET("/", helloWorld)
	r.GET("/api/movies", movieController.GetMovies)
	r.GET("/api/movies/{id}", movieController.GetMovieById)
	r.POST("/api/movies", movieController.CreateMovie)
	r.PUT("/api/movies", movieController.UpdateMovie)
	r.DELETE("/api/movies/{id}", movieController.DeleteMovie)

	r.SERVE(os.Getenv("PORT"))
	//r.SERVE(":8000")
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world!</h1>")
}
