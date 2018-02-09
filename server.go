package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const LocalServicePort string = ":3000"
const ConfigFilename string = "config.txt"

type Endpoint struct {
	IP   string `json:"ip"`
	Port uint32 `json:"port"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api/v1/config").Subrouter()
	sub.HandleFunc("/remote_address", handleAddress).Methods("GET", "POST")
	log.Fatal(http.ListenAndServe(LocalServicePort, router))
}

// handle GET and POST on http://localhost:3000/api/v1/config/remote_address
func handleAddress(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		dat, err := ioutil.ReadFile(ConfigFilename)
		check(err, res)
		fmt.Printf("read from %s: %s\n", ConfigFilename, string(dat))

		res.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		res.Write(dat)
	case "POST":
		body, err := ioutil.ReadAll(req.Body)
		check(err, res)

		// save text to config file
		f, err := os.Create(ConfigFilename)
		check(err, res)
		defer f.Close()
		b, err := f.Write(body)
    	check(err, res)
		fmt.Printf("wrote %d bytes to %s: %s\n", b, ConfigFilename, string(body))

		//w.Header().Set("Location", r.URL.Path+"/" + com.Name)
		res.WriteHeader(http.StatusCreated)
	}
}

func check(err error, res http.ResponseWriter) {
    if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
    }
}
