<p align="middle"><img src="logo.png" width="400px"></p>

<h1 align="middle" style="text-align:center">URL Shortener</h1><br>


<h4>HTTP server for shortening your links.</h4>

## Technologies used

* GoLang
* chi-router
* Redis

## Prerequisites

* GoLang

## Installation guide

#### To run this project, you will need to set up the following environment variables:

* PORT - informs the application on which port it should listen
* REDIS_ADDR - address of your Redis instance
* REDIS_PASS - if your Redis is secured with password, pass it using this variable

#### The instruction below will help you to get the working copy of the application on your machine:

````
# Clone the repository
go get github.com/dawidcyron/url-shortener

# Navigate to your Go project directory
cd $GOPATH/src/github.com/dawidcyron/url-shortener

# Get the required dependencies
go get

# Install the application
go install

# Navigate to the binary
cd $GOBIN

# Run the application
./shortener
````

