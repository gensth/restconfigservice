package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var configFilename string
var postSetCommand string

func main() {
	// parse command line arguments
	localServicePort := flag.Uint("port", 3000, "the port of the local HTTP REST service")
	flag.StringVar(&configFilename, "config", "config.txt", "the name of the config file to read/write")
	flag.StringVar(&postSetCommand, "postset", "./postSetConfig.sh", "the command to execute after writing the config or empty string if nothing to execute")
	flag.Parse()

	// configuring the HTTP REST service
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api/v1/config").Subrouter()
	sub.HandleFunc("/remote_address", handleAddress).Methods("GET", "POST")
	listenAddress := fmt.Sprintf(":%d", *localServicePort) // e.g. ":3000"
	// logging the service configuration
	fmt.Print("launching REST service\n")
	fmt.Printf("    listening to: %s\n", listenAddress)
	fmt.Printf("    reading/writing from/to: %s\n", configFilename)
	if postSetCommand != "" {
		fmt.Printf("    executing after SET: %s\n", postSetCommand)
	}
	// launching the HTTP REST service
	log.Fatal(http.ListenAndServe(listenAddress, router))
}

// Handle GET and POST on http://localhost:3000/api/v1/config/remote_address
func handleAddress(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		// read text from config file
		dat, err := ioutil.ReadFile(configFilename)
		check(err, res)
		fmt.Printf("read from %s: %s\n", configFilename, string(dat))

		// respond http request
		res.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		res.Write(dat)
	case "POST":
		body, err := ioutil.ReadAll(req.Body)
		check(err, res)

		// save text to config file
		f, err := os.Create(configFilename)
		check(err, res)
		defer f.Close()
		b, err := f.Write(body)
    	check(err, res)
		fmt.Printf("wrote %d bytes to %s: %s\n", b, configFilename, string(body))

		// call the post-set command
		if postSetCommand != "" {
            fmt.Printf("executing the post-set command: %s\n", postSetCommand)
			err = exec.Command(postSetCommand).Run()
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
