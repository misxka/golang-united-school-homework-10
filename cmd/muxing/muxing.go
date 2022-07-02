package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", ViewName).Methods(http.MethodGet)
	router.HandleFunc("/bad", ViewBad).Methods(http.MethodGet)
	router.HandleFunc("/data", ViewData).Methods(http.MethodPost)
	router.HandleFunc("/headers", ViewHeaders).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func ViewName(w http.ResponseWriter, r *http.Request) {
	if name, ok := mux.Vars(r)["PARAM"]; ok {
		w.Write([]byte("Hello, " + name + "!"))
	}
	return
}

func ViewBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	return
}

func ViewData(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		w.Write(append([]byte("I got message:\n"), body...))
	}
	return
}

func ViewHeaders(w http.ResponseWriter, r *http.Request) {
	headers := r.Header

	headerName := ""
	headerValue := ""
	total := 0

	for k, v := range headers {
		if headerName == "" {
			headerName += k
		} else {
			headerName += ("+" + k)
		}

		intValue, err := strconv.Atoi(v[0])
		if err == nil {
			total += intValue
		}
	}

	headerValue = strconv.Itoa(total)
	w.Header().Set(strings.ToLower(headerName), headerValue)
	return
}
