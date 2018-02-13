# RestConfigService

A very simple REST service (written in Go) to read and write the contents of a config file.

## Features

*   On HTTP GET
    *   Responding the text read from a config file.
*   On HTTP POST
    *   Write the submitted text to a config file
    *   and call a given command (e.g. to restart another service).

## Compile and run

Execute in the root directory of this project

	$ go build
	
	$ ./restconfigservice --help
	Usage of ./restconfigservice:
      -config string
        	the name of the config file to read/write (default "config.txt")
      -port uint
        	the port of the local HTTP REST service (default 3000)
      -postset string
        	the command to execute after writing the config or empty string if nothing to execute (default "./postSetConfig.sh")

	$ ./restconfigservice
	launching REST service
        listening to: http://localhost:3000/api/v1/config
        reading/writing from/to: config.txt
        executing after SET: ./postSetConfig.sh


## Interaction with the REST service

Interact with the running service by sending HTTP GET/POST requests with

*   curl

        # GET the current config
        curl "http://localhost:3000/api/v1/config"
        
        # or POST (write) the config
        curl --data "192.168.1.1:8080" "http://localhost:3000/api/v1/config"
*   [Postman](http://www.getpostman.com), an easy to use UI to interact with REST APIs for macOS, Windows and Linux
