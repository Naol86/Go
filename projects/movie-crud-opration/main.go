package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type response struct {
	Count int `json:"count"`
	Data []Movie `json:"data"`
}

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	var data response
	data.Count = len(movies)
	data.Data = movies
	json.NewEncoder(w).Encode(data)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	// json.NewEncoder(w).Encode(movies)
	var data response
	data.Count = len(movies)
	data.Data = movies
	json.NewEncoder(w).Encode(data)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			// json.NewEncoder(w).Encode(item)
			var data response
			data.Count = 1
			data.Data = append(data.Data, item)
			json.NewEncoder(w).Encode(data)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	// json.NewEncoder(w).Encode(movie)
	var data response
	data.Count = 1
	data.Data = append(data.Data, movie)
	json.NewEncoder(w).Encode(data)
}


func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			// json.NewEncoder(w).Encode(movie)
			var data response
			data.Count = len(movies)
			data.Data = movies
			json.NewEncoder(w).Encode(data)
			return 
		}
	}
}


func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID: "1",
		Isbn: "123",
		Title: "movie one",
		Director: &Director{
			FirstName: "naol",
			LastName: "kasinet",
		},
	})
	movies = append(movies, Movie{
		ID: "2",
		Isbn: "323",
		Title: "movie two",
		Director: &Director{
			FirstName: "john",
			LastName: "smith",
		},
	})
	movies = append(movies, Movie{
		ID: "3",
		Isbn: "523",
		Title: "movie three",
		Director: &Director{
			FirstName: "abel",
			LastName: "charles",
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}