package main

import (
	"fmt"
	"github.com/matiascfgm/restAPI/controller"
	router "github.com/matiascfgm/restAPI/http"
	"github.com/matiascfgm/restAPI/service"
	"net/http"
)

var (
	movieService    service.MovieService       = service.NewMovieService(nil)
	movieController controller.MovieController = controller.NewMovieController(movieService)
	r               router.Router              = router.NewMuxRouter()
)

func main() {
	r.GET("/", helloWorld)
	//r.HandleFunc("/", helloWorld)
	r.GET("/api/movies", movieController.GetMovies)
	//r.HandleFunc("/api/movies/{id}", getMovie).Methods("GET")
	r.POST("/api/movies", movieController.CreateMovie)
	//r.HandleFunc("/api/movies/{id}", updateMovie).Methods("PUT")
	//r.HandleFunc("/api/movies/{id}", deleteMovie).Methods("DELETE")
	r.SERVE(":8000")
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world!</h1>")
}
