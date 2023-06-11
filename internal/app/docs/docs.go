package docs

import "github.com/swaggo/swag"

const docSwagger = `{
	"swagger": "2.0",
	"info": {
	  "title": "{{.Title}}",
	  "version": "{{.Version}}",
	  "description": "{{escape .Description}}"
	},
	"host": "{{.Host}}",
	"basePath": "{{.BasePath}}",
	"schemes": [
	  "http"
	],
	"produces": [
	  "application/json"
	],
	"consumes": [
	  "application/json"
	],
	"paths": {
	  "/users/all": {
		"get": {
		  "tags": ["users"],
		  "summary": "Lists all users.",
		  "responses": {
			"200": {
			  "description": "Success",
			  "schema": {
				"type": "array",
				"items": {
				  "$ref": "#/definitions/User"
				}
			  }
			},
			"500": {
			  "description": "Internal Server Error",
			  "schema": {
				"$ref": "#/definitions/ErrorResponse"
			  }
			}
		  }
		}
	  },
	  "/users/create": {
		"post": {
		  "tags": ["users"],
		  "summary": "Creates a new user.",
		  "responses": {
			"201": {
			  "description": "Success",
			  "schema": {
				"$ref": "#/definitions/User"
			  }
			},
			"400": {
			  "description": "Bad Request",
			  "schema": {
				"$ref": "#/definitions/ErrorResponse"
			  }
			},
			"500": {
			  "description": "Internal Server Error",
			  "schema": {
				"$ref": "#/definitions/ErrorResponse"
			  }
			}
		  }
		}
	  },
	  "/users/find/{id}": {
		"get": {
		  "tags": ["users"],
		  "summary": "Gets a user by ID.",
		  "parameters": [
			{
			  "name": "id",
			  "in": "path",
			  "required": true,
			  "type": "integer"
			}
		  ],
		  "responses": {
			"200": {
			  "description": "Success",
			  "schema": {
				"$ref": "#/definitions/User"
			  }
			},
			"404": {
			  "description": "Not Found",
			  "schema": {
				"$ref": "#/definitions/ErrorResponse"
			  }
			},
			"500": {
			  "description": "Internal Server Error",
			  "schema": {
				"$ref": "#/definitions/ErrorResponse"
			  }
			}
		  }
		}
	  }
	},
	"definitions": {
	  "User": {
		"type": "object",
		"properties": {
		  "id": {
			"type": "integer",
			"example": 1
		  },
		  "username": {
			"type": "string",
			"example": "johndoe"
		  },
		  "email": {
			"type": "string",
			"example": "johndoe@example.com"
		  }
		},
		"required": ["id", "username", "email"]
	  },
	  "ErrorResponse": {
		"type": "object",
		"properties": {
		  "message": {
			"type": "string",
			"example": "Invalid request"
		  }
		},
		"required": ["message"]
	  }
	}
  }
  `

var SwaggerInfo_swagger = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:6476",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Swagger API",
	Description:      "This is a API Server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docSwagger,
}

func SwaggerInit() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}
