# RestConfigService

A very simple REST service (written in Go) to read and write a config file.

Features

* Reading text from a config file on GET
* Writing text to a config file on POST

## Compile and run

Execute in the root directory of this project

	$ go build
	$ ./restconfigservice

## Interaction with the REST service

When the service is running, use e.g. [Postman](http://www.getpostman.com) to GET/POST from/to `http://localhost:3000/api/v1/config/remote_address`.
