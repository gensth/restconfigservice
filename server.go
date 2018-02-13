package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const LocalServicePort uint16 = 3000
const ConfigFilename string = "config.txt"
const PostSetCommand string = "./postSetConfig.sh"

func main() {
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api/v1/config").Subrouter()
	sub.HandleFunc("/remote_address", handleAddress).Methods("GET", "POST")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", LocalServicePort), router))
}

// Handle GET and POST on http://localhost:3000/api/v1/config/remote_address
func handleAddress(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		// read text from config file
		dat, err := ioutil.ReadFile(ConfigFilename)
		check(err, res)
		fmt.Printf("read from %s: %s\n", ConfigFilename, string(dat))

		// respond http request
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

		// call the post-set command
		if PostSetCommand != "" {
			fmt.Printf("executing the post-set command: %s\n", PostSetCommand)
			err = exec.Command(PostSetCommand).Run()
			check(err, res)
		}

		// respond http request
		//w.Header().Set("Location", r.URL.Path+"/" + com.Name)
		res.WriteHeader(http.StatusCreated)
	}
}

// Check an error code and if an error occured respond an http error.
func check(err error, res http.ResponseWriter) {
    if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
    }
}
