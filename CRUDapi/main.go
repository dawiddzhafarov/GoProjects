package main

// CRUD API that allows to create new movie objects, delete, update and get them.
// Used with gorilla/mux package and net.http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Movie struct as an object
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isdn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Director struct as an object
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Movies slice acting as a database
var movies []Movie

// Main function handling requests and starting the server
func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "3241321", Title: "SpiderMan", Director: &Director{Firstname: "Quentin", Lastname: "Tarantino"}})
	movies = append(movies, Movie{ID: "2", Isbn: "334571", Title: "Batman", Director: &Director{Firstname: "Lorenzo", Lastname: "Quartrao"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// getMovies function is responsible for getting ale the movies 'objects'
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// deleteMovie function deletes the movie which id was sent in the request body
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// getMovie function that gets the movie by id which was sent in a request body
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// createMovie function creates new movie objects and adds it to the movies database, based on request body parameters
// Use PostMan to create correct request
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// updateMovie function "updates" the movie by deleting old object and creating new one with assigning old id
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
