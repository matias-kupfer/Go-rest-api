package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/matiascfgm/Go-rest-api/entity"
	"github.com/matiascfgm/Go-rest-api/errors"
	"github.com/matiascfgm/Go-rest-api/service"
	"net/http"
)

type MovieController interface {
	GetMovies(w http.ResponseWriter, r *http.Request)
	GetMovieById(w http.ResponseWriter, r *http.Request)
	CreateMovie(w http.ResponseWriter, r *http.Request)
	DeleteMovie(w http.ResponseWriter, r *http.Request)
	UpdateMovie(w http.ResponseWriter, r *http.Request)
}

type controller struct{}

var (
	movieService service.MovieService
)

func NewMovieController(service service.MovieService) MovieController {
	movieService = service
	return &controller{}
}

// FUNCTIONS
func (*controller) GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	movies, err := movieService.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{"Error getting the movies"})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}

func (*controller) GetMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movie, err := movieService.GetMovieById(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{"Error getting the movie"})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}

func (*controller) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movie, err := movieService.DeleteMovie(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{"Error getting the movie"})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}

func (*controller) CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie entity.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{"Error unmarshal data"})
	}
	err1 := movieService.Validate(&movie)
	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{err1.Error()})
	}
	result, err2 := movieService.Create(&movie)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{"Error saving the movie"})
	}
	json.NewEncoder(w).Encode(result)
}

func (*controller) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie entity.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{"Error unmarshal data"})
	}
	err1 := movieService.Validate(&movie)
	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{err1.Error()})
	}
	result, err2 := movieService.UpdateMovie(&movie)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{"Error saving the movie"})
	}
	json.NewEncoder(w).Encode(result)
}

/*func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.ID == params["id"] {
			var movie entity.Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			break
		}
	}
	json.NewEncoder(w).Encode(&entity.Movie{})
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[:index+1]...)
			break
		}
	}
	json.NewEncoder(w).Encode(&entity.Movie{})
}*/
