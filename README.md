# RestConfigService

A very simple REST service (written in Go) to read and write a config file.

Features

*   On GET: Responding the text read from a config file.
*   On POST: Writing the submitted text to a config file and call a given command (e.g. restart another service).

## Compile and run

Execute in the root directory of this project

	$ go build
	$ ./restconfigservice

## Interaction with the REST service

When the service is running, use e.g. [Postman](http://www.getpostman.com) to GET/POST from/to `http://localhost:3000/api/v1/config`.
