// Extended example of using select statements to collect data from multiple API requests
// Written by: Waqqas Sheikh (https://www.github.com/w-k-s)
// For: Dubai DevFest 2018 (https://www.meetup.com/en-AU/GDG-Dubai/events/253941428/)

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// You'll need to create an account at https://www.themoviedb.org/documentation/api
// to get an API key and run the sample
const TMDB_API_KEY = ""

// A struct used to contain the response of a netwrok request.
// If the response is successful, response will be stored in data as bytes.
// If the response fails, the error will be stored in err.
type NetworkResult struct {
	statusCode int
	data       []byte
	err        error
}

// A Movie type
type Movie struct {
	title      string
	trailerUrl string
}

// Represents movie as a string
func (m Movie) String() string {
	return fmt.Sprintf("%s : %s", m.title, m.trailerUrl)
}

// Sends an HTTP GET request asynchronously to the given 'url'.
// Sends a pointer to 'NetworkResult' to the given 'nc' channel
func getAsync(url string, nc chan NetworkResult) {

	timeout := time.Duration(30 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(url)
	if err != nil {
		nc <- NetworkResult{err: err}
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		nc <- NetworkResult{err: err}
		return
	}

	nc <- NetworkResult{statusCode: resp.StatusCode, data: bytes}
}

// A structure to contain the result of loading a Movie using 'loadMovie(movieId string, mc chan MovieResult)'
// If loading movie is successful, the movie field will be set and error will be nil
// If loading movie fails, err will be set and movie will be nil
type MovieResult struct {
	movie *Movie
	err   error
}

// loads movie with given 'movieId' asynchronously and sends the result to the 'mc' 'MovieResult' channel
func loadMovie(movieId string, mc chan MovieResult) {

	// channel that will receive the response of the /movie (movie details) request.
	detailsChannel := make(chan NetworkResult)
	// channel that will receive the response of the movie/{id}/videos (movie trailer) request.
	trailerUrlChannel := make(chan NetworkResult)

	// result of the /movie (movie details) request
	var detailsResult NetworkResult
	// result of the movie/{id}/videos (movie trailer) request
	var trailerUrlResult NetworkResult
	// total number of requests that will be inflight concurrently
	inflight := 2

	go getAsync("https://api.themoviedb.org/3/movie/"+movieId+"?api_key="+TMDB_API_KEY, detailsChannel)
	go getAsync("https://api.themoviedb.org/3/movie/"+movieId+"/videos"+"?api_key="+TMDB_API_KEY, trailerUrlChannel)

	for {
 		// This for loop will wait to receive a message from one of the two channels.
		// Every time a request completes, the inflight request count is decremented
		// and a check is done to see if there are any remaining inflight requests.
		// If there are, the loop continues from the beginning;
		// otherwise, the for loop is exited
		
		select {
		case detailsResult = <-detailsChannel:
			inflight--
		case trailerUrlResult = <-trailerUrlChannel:
			inflight--
		}

		//once there are no more inflight requests, exit the loop
		if inflight == 0 {
			break
		}
	}

	// if either of the result has an error, send a MovieResult with the error
	if err := checkAPIResults(detailsResult, trailerUrlResult); err != nil {
		mc <- MovieResult{err: err}
		return
	}

	// The syntax for the next part can be a bit confusing.
	//
	// We need to map the keys in the received jsons to fields in structs.
	//
	// Go allows us to declare a variable and define it's struct all in one line.
	// That's what we're doing below.
	// Notice in the case of 'trailers', we're able to define nested structs as well.
	// 
	// Each Field is tagged with its correspending JSON key.

	var detail struct {
		OriginalTitle string `json:"original_title"`
	}

	var trailers struct {
		Results []struct {
			Site string `json:"site"`
			Key  string `json:"key"`
		} `json:"results"`
	}

	// parse the bytes as json and map the keys into the structs declared above
	// We can ignore the errors here, sicne we expect the JSON to be valid
	json.Unmarshal(detailsResult.data, &detail)
	json.Unmarshal(trailerUrlResult.data, &trailers)

	// Create a new movie instance
	movie := &Movie{}
	movie.title = detail.OriginalTitle

	// iterate over the trailers to find the one hosted on YouTube.
	if len(trailers.Results) > 0 {
		for _, result := range trailers.Results {
			if result.Site == "YouTube" {
				movie.trailerUrl = "https://www.youtube.com/watch?v=" + result.Key
			}
		}
	}

	// Send the result back through the movie channel.
	mc <- MovieResult{movie: movie}
}

// Receives an array of NetworkResults
// Iterates over the results and returns the first non-nil error
// Returns nil if all NetworkResults are successful
func checkAPIResults(results ...NetworkResult) error {

	for _, result := range results {
		
		if result.err != nil {
			return result.err
		}

		if result.statusCode >= http.StatusBadRequest {
			
			var err struct {
				Message string `json:"status_message"`
			}
			json.Unmarshal(result.data, &err)

			return errors.New(err.Message)
		}
	}
	
	return nil
}

func main() {
	//movie channel
	movieChannel := make(chan MovieResult)

	movieId := "694"
	go loadMovie(movieId, movieChannel)

	result := <-movieChannel

	if result.err != nil {
		fmt.Printf("Error loading movie: %s\n", result.err)
		return
	}

	fmt.Println(result.movie)
}