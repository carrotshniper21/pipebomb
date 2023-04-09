// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/films/vip/search": {
            "get": {
                "description": "Searches for a film and returns a JSON response",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "films"
                ],
                "summary": "Search for a film",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search Query",
                        "name": "q",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/film.FilmResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "film.FilmResponse": {
            "description": "is the response struct",
            "type": "object",
            "properties": {
                "film": {
                    "$ref": "#/definitions/film.FilmStruct"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "film.FilmStruct": {
            "description": "stores the film data",
            "type": "object",
            "properties": {
                "href": {
                    "type": "string",
                    "example": "https://example.com/film/1"
                },
                "id": {
                    "type": "string",
                    "example": "film-1"
                },
                "idParts": {
                    "$ref": "#/definitions/film.IdSplit"
                },
                "poster": {
                    "type": "string",
                    "example": "https://example.com/poster/1.jpg"
                }
            }
        },
        "film.IdSplit": {
            "description": "stores the film ID parts",
            "type": "object",
            "properties": {
                "idNum": {
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "example": "Film 1"
                },
                "type": {
                    "type": "string",
                    "example": "film"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Pipebomb API",
	Description:      "Pipebomb API for searching and streaming movies",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}