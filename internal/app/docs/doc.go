package docs

import "github.com/labstack/echo/v4"

type User struct {
	// The unique ID of the user
	//
	// required: true
	// example: 1
	ID int `json:"id"`
	// The username of the user
	//
	// required: true
	// example: johndoe
	Username string `json:"username"`
	// The email address of the user
	//
	// required: true
	// example: johndoe@example.com
	Email string `json:"email"`
}

// ErrorResponse represents an error response from the API.
//
// swagger:model
type ErrorResponse struct {
	// The error message
	//
	// required: true
	// example: Invalid request
	Message string `json:"message"`
}

// swagger:route GET /users users listUsers
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
func listUsers(c echo.Context) error {
	// Your implementation to fetch and return users
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
	// Your implementation to fetch and return a user by ID
	return c.JSON(200, User{})
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
	// Your implementation to create a new user
	return c.JSON(201, User{})
}

// swagger:route PUT /users/{id} users updateUser
//
// Updates an existing user.
//
// This will update the details of an existing user based on the provided ID.
//
//	Consumes:
//	- application/json
//	Produces:
//	- application/json
//
//	Responses:
//	  200: User
//	  400: ErrorResponse
//	  404: ErrorResponse
//	  500: ErrorResponse
func updateUser(c echo.Context) error {
	// Your implementation to update an existing user
	return c.JSON(200, User{})
}

// swagger:route DELETE /users/{id} users deleteUser
//
// Deletes a user.
//
// This will delete a user based on the provided ID.
//
//	Produces:
//	- application/json
//
//	Responses:
//	  204:
//	  404: ErrorResponse
//	  500: ErrorResponse
func deleteUser(c echo.Context) error {
	return c.NoContent(204)
}
