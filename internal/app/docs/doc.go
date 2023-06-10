package docs

// Package main is a sample API for demonstrating Swagger with Go Echo.
//
// The API has various endpoints for managing users.
//
//     Schemes: http
//     BasePath: /api
//     Version: 1.0.0
//     License: MIT http://opensource.org/licenses/MIT
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// User represents a user in the system.
//
// swagger:model
type User struct {
	// The unique ID of the user.
	//
	// required: true
	// example: 1
	ID int `json:"id"`

	// The username of the user.
	//
	// required: true
	// example: johndoe
	Username string `json:"username"`

	// The email address of the user.
	//
	// required: true
	// example: johndoe@example.com
	Email string `json:"email"`
}

// ErrorResponse represents an error response from the API.
//
// swagger:model
type ErrorResponse struct {
	// The error message.
	//
	// required: true
	// example: Invalid request
	Message string `json:"message"`
}

// swagger:route GET /users users getUserAll
//
// Lists all users.
//
// This will show all available users.
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: []User
//	  500: ErrorResponse
func getUserAll(c echo.Context) error {
	// Your implementation to fetch and return all users.
	return c.JSON(200, []User{})
}

// swagger:route GET /users/{id} users getUserByID
//
// Gets a user by ID.
//
// This will return the details of a user based on the provided ID.
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: User
//	  404: ErrorResponse
//	  500: ErrorResponse
func getUserByID(c echo.Context) error {
	// Your implementation to fetch and return a user by ID.
	return c.JSON(200, User{})
}

// swagger:parameters createUser
type createUserRequest struct {
	// The user object to be created.
	//
	// in: body
	// required: true
	Body User
}

// swagger:route POST /users users createUser
//
// Creates a new user.
//
// This will create a new user with the provided details.
//
//	Consumes:
//	- application/json
//	Produces:
//	- application/json
//
//	Responses:
//	  201: User
//	  400: ErrorResponse
//	  500: ErrorResponse
func createUser(c echo.Context) error {
	// Your implementation to create a new user.
	return c.JSON(201, User{})
}

// ServeSwaggerUI serves the Swagger UI for viewing the generated documentation.
func ServeSwaggerUI(c echo.Context) error {
	url := "swagger.json" // Path to your generated swagger.json or swagger.yml file
	return c.Redirect(http.StatusTemporaryRedirect, url)
}
