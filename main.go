package main

//Example Test Server

// import necessary packages
// net/http package allows use of servemux multiplexer
import (
	"context"
	"fmt"
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
	log.Print("listening..")

	// create a server that listens to incoming requests
	http.ListenAndServe(":8080", mux)
}

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

var indexHTMLTemplate = template.Must(template.New("indexHTML").Parse(indexHTML))

type MovieGetter interface {
	GetTrendingMovies(ctx context.Context) (*api.Response, error)
}

func handler(w http.ResponseWriter, r *http.Request) {
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

	randomIndex := rand.Intn(len(response.Movies))
	err = indexHTMLTemplate.Execute(w, response.Movies[randomIndex])
	if err != nil {
		// This is kinda hopeless
		http.Error(w, fmt.Sprintf("failed to write response: %v", err), http.StatusInternalServerError)
		return
	}

	//use handler logic to servemux to a specific endpoint
	//ex: localhost:8080/api and displays pretty json
}

//A go interface is similar to a pure abstract class in C++
// https://stackoverflow.com/questions/8970157/is-it-possible-to-mimic-go-interface-in-c-c

//idea:
//display lists and movie reviews to users based on his genre of choice

/*
TODO: consume api and figure out a way to store api key ?
      [] Maybe use vault as a use-case ?
	  [] logic for consuming api use postman to validate that the server code is working as intended
	  //Go server to send a request to an api endpoint and get back data as a byte slice

	  [] request different params: trending movies, maybe ask a user for a movie id and then process it and display info of that movie -> original use case- prompt user to enter his favorite genre based on that input server sends a response



	  Structure:
	  	Make a struct to display a list of params we will use:

*/
