package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	chartisan "github.com/chartisan/golang"
	"github.com/julienschmidt/httprouter"
)

func setHeaders(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
}

// randomData returns a random number between 0 and 50 n tmes.
func randomData(n int) []float64 {
	list := []float64{}
	for i := 0; i < n; i++ {
		list = append(list, rand.Float64()*100)
	}
	return list
}

// randomString generates a random strings of 10 chars n times.
func randomStrings(n int) []string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	strings := []string{}
	for i := 0; i < n; i++ {
		b := make([]byte, 10)
		for i := range b {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		}
		strings = append(strings, string(b))
	}
	return strings
}

// example shows how to send a chart response.
func example(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	setHeaders(w)
	values := 3
	chart := chartisan.Build().
		Labels(randomStrings(values)).
		Dataset("Sample 1", randomData(values)).
		Dataset("Sample 2", randomData(values)).
		ToJSON()
	fmt.Fprintf(w, chart)
}

// invalid is an example of an invalid chart.
func invalid(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	setHeaders(w)
	fmt.Fprintf(w, "{\"a\": 10}")
}

func main() {
	router := httprouter.New()
	router.GET("/", example)
	router.GET("/invalid", invalid)
	log.Println("Starting server...")
	log.Fatalln(http.ListenAndServe(":9000", router))
}
