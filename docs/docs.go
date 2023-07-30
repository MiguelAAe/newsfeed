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
        "/newsFeed": {
            "get": {
                "description": "Gets a list of all feeds",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Gets a list of all feeds",
                "parameters": [
                    {
                        "description": "a map of sources to list of categories",
                        "name": "query",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/feed.feedQuery"
                        }
                    },
                    {
                        "type": "string",
                        "description": "source\t\t\t\tof the feed",
                        "name": "source",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "category\tof\tthe\tfeed",
                        "name": "category",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "a list of feeds",
                        "schema": {
                            "$ref": "#/definitions/feed.feedQuery"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/newsFeed/categories": {
            "get": {
                "description": "Gets a list of all feed categories",
                "produces": [
                    "application/json"
                ],
                "summary": "Gets a list of all feed categories",
                "responses": {
                    "200": {
                        "description": "a list of categories",
                        "schema": {
                            "$ref": "#/definitions/feed.category"
                        }
                    }
                }
            }
        },
        "/newsFeed/names": {
            "get": {
                "description": "Gets a list of all feed names",
                "produces": [
                    "application/json"
                ],
                "summary": "Gets a list of all feed names",
                "responses": {
                    "200": {
                        "description": "a list of names",
                        "schema": {
                            "$ref": "#/definitions/feed.names"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "feed.category": {
            "type": "object",
            "properties": {
                "categories": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "feed.feedQuery": {
            "type": "object",
            "properties": {
                "query": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "array",
                        "items": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "feed.names": {
            "type": "object",
            "properties": {
                "names": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}