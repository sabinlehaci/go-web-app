package main

//Example Test Server

// import necessary packages
// net/http package allows use of servemux multiplexer
import (
	"context"
	"fmt"
	"path/filepath"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"github.com/sabinlehaci/go-web-app/api"
)

func main() {

	// a servemux (aka router) stores mapping btwn URL path for app
	// and associated handlers

	mux := http.NewServeMux()
	//Convert handler func to a http.HandlerFunc type
	th := http.HandlerFunc(handler)

	//and add it to ServeMux
	mux.Handle("/", th)
	cwd, _ := os.Getwd()

	log.Print("listening..")
	log.Print( filepath.Join( cwd, "./assets/index.html" ) )

	http.ListenAndServe(":8080", mux)
}

var indexHTMLTemplate = template.Must(template.ParseGlob("assets/index.html"))  



//put this in its own package
const indexHTML = `
<!doctype html>
<html lang=en>
<head>
	<meta charset=utf-8>
	<title>Sabins Movie Night</title>
    <script src="https://cdn.tailwindcss.com"></script>
</head>
<body>
	<div class="bg-neutral-100">
		<h1 class="text-xl">Sabins Movie Night</h1>
		<p>You have been selected a random trending movie!</p>
		<p>Your movie title is: {{ .Title }}</p>
		<p class="text-sm">{{ .Overview }}</p>
	</div>
</body>
</html>
`

// create a predefined template that can be replicated for each of the selected movies
// upon refreshing the web app page


type MovieGetter interface {
	GetTrendingMovies(ctx context.Context) (*api.Response, error)
}

func handler(w http.ResponseWriter, r *http.Request) {
	//here we reference our env var that we set as TMBD
	var cli MovieGetter = &api.Client{

		// retrieve environment variable 'TMDB' that stores API key
		APIKey: os.Getenv("TMDB"),
	}

	response, err := cli.GetTrendingMovies(r.Context())

	// error handling
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get trending movies: %v", err), http.StatusInternalServerError)
		return
	}

	//logic for randomizing movies on every refresh
	randomIndex := rand.Intn(len(response.Movies))
	err = indexHTMLTemplate.Execute(w, response.Movies[randomIndex])
	if err != nil {
		// This is kinda hopeless
		http.Error(w, fmt.Sprintf("failed to write response: %v", err), http.StatusInternalServerError)
		return
	}
}


//Separate handlers into its own dir to clean up this file 
//Cannot show img object 
//func imgHandler(w http.ResponseWriter, r*http.Request) {
//	http.ServeFile(w,r,"index.html")
//	http.Handle()
//} 

