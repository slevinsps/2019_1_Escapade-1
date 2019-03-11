// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-03-11 12:11:52.191419698 +0300 MSK m=+0.130443159

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "Documentation",
        "title": "Escapade API",
        "contact": {},
        "license": {},
        "version": "1.0"
    },
    "host": "https://escapade-backend.herokuapp.com",
    "basePath": "/api/v1",
    "paths": {
        "/users/{name}/profile": {
            "get": {
                "description": "return public information, such as name or best_score",
                "summary": "Get some of user fields",
                "operationId": "GetProfile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Profile email",
                        "name": "email",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "Profile nickname",
                        "name": "nickname",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Profile found successfully",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.UserPublicInfo"
                        }
                    },
                    "400": {
                        "description": "Invalid username",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Result"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Result"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Result": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "place": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "models.UserPublicInfo": {
            "type": "object",
            "properties": {
                "bestScore": {
                    "type": "string"
                },
                "bestTime": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
