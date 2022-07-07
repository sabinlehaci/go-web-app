package handler

import (
	"context"
	"embed"
	"fmt"
	"math/rand"
	"net/http"
	"text/template"

	"github.com/sabinlehaci/go-web-app/db"
	"github.com/sabinlehaci/go-web-app/tmdbApi"
)

//go:embed assets
var static embed.FS

type MovieGetter interface {
	GetTrendingMovies(ctx context.Context) (*tmdbApi.Response, error)
}

type Handlers struct {
	MovieGetter MovieGetter
	DB          db.Querier
}

var indexHTMLTemplate = template.Must(template.ParseFS(static, "assets/index.html"))

type templateVariables struct {
	Favorites []db.FavoriteMovie
	Movie     tmdbApi.Movie
}

func (h *Handlers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response, err := h.MovieGetter.GetTrendingMovies(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get trending movies: %v", err), http.StatusInternalServerError)
		return
	}

	movies, err := h.DB.ListMovies(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get favorite movies: %v", err), http.StatusInternalServerError)
		return
	}

	//logic for randomizing movies on every refresh
	randomIndex := rand.Intn(len(response.Movies))
	vars := templateVariables{
		Movie:     response.Movies[randomIndex],
		Favorites: movies,
	}
	err = indexHTMLTemplate.Execute(w, vars)
	if err != nil {
		// This is kinda hopeless
		http.Error(w, fmt.Sprintf("failed to write response: %v", err), http.StatusInternalServerError)
		return
	}
}
