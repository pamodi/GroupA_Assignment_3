# Group A - Assignment 3

This repository contains RESTful APIs and tests that manages a collection of items which is developed using golang by Group A.

## Purpose of the Project 
* To develop RESTful APIs and tests using Golang.
* To practice adding the code to github reporitory.
* Team Collaboration.
* Merging the code from different team members and resolving the conflicts before pushing to the repository.

## API Specification

* POST /items
	* Adds a new item to the collection.
	* Struct has been defined to store the item properties.
	* Used in-memory storage for the items.
	* Auto generated id is assigned as the item ID using Go’s standard library.
	* Request body contains the ID and name of the item.
	* API will returns the created item with its ID.

* GET /items 
	* Retrieves a list of all items along with all item details.

* Test for POST items API
	* Included tests to check for correct HTTP status codes.
	* Validated response bodies.

* Test for GET items API
	* Included tests to check for correct HTTP status codes.


## Getting Started 

Follow the following steps inorder to get the project run on your environment.

* Install Go

* Clone the project

```
git clone https://github.com/pamodi/GroupA_Assignment_3.git
```

* Open the project using Visual Studio Code

* Open the terminal and run the project

```
go run main.go
```

* Curl commands to run the APIs

```
curl -X GET http://localhost:8012/items
curl -X POST http://localhost:8012/items -H "Content-Type: application/json" -d '{"name":"Test Item"}'
```

Both VSCode terminal and postman application can be used to run the APIs.

* Run the tests

```
go test -v
```
