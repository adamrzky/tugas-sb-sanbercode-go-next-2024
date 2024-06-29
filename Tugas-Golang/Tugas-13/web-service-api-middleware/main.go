package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Movie struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

var AllMovies = []Movie{
	{1, "Spider-Man", 2002},
	{2, "John Wick", 2014},
	{3, "Avengers", 2018},
	{4, "Logan", 2017},
}

// GetMovies
func getMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataMovies, err := json.Marshal(AllMovies)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataMovies)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}

// PostMovie
func PostMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Mov Movie
	Mov.ID = len(AllMovies) + 1
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&Mov); err != nil {
				log.Fatal(err)
			}
		} else {
			// parse dari form-data
			title := r.PostFormValue("title")
			getYear := r.PostFormValue("year")
			year, _ := strconv.Atoi(getYear)
			Mov.Title = title
			Mov.Year = year
		}

		AllMovies = append(AllMovies, Mov)
		dataMovie, _ := json.Marshal(Mov) // to byte
		w.Write(dataMovie)                // cetak di browser
		return
	}

	http.Error(w, "NOT FOUND", http.StatusNotFound)
	return
}

// Fungi Log yang berguna sebagai middleware
func basicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "admin" && pwd == "admin" {
			next.ServeHTTP(w, r)
			return
		}
		w.Write([]byte("Username atau Password tidak sesuai"))
		return
	})
}

func main() {
	http.HandleFunc("/movies", getMovies)

	http.Handle("/create-movie", basicAuth(http.HandlerFunc(PostMovie)))
	fmt.Println("server running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
