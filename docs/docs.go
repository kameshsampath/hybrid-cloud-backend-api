// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Kamesh Sampath",
            "email": "kamesh.sampath@solo.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health/live": {
            "get": {
                "description": "Checks the API liveness, can be used with Kubernetes Probes",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Checks the API liveness",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/health/ready": {
            "get": {
                "description": "Checks the API readiness, can be used with Kubernetes Probes",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Checks the API readiness",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/process": {
            "post": {
                "description": "process the request message from the front end and apply the needed transformations",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "backend"
                ],
                "summary": "process the request message",
                "parameters": [
                    {
                        "description": "Message to process",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routes.Message"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/routes.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "routes.Message": {
            "type": "object",
            "properties": {
                "request": {
                    "$ref": "#/definitions/routes.Request"
                },
                "requestId": {
                    "description": "RequestId the unique request id",
                    "type": "string"
                }
            }
        },
        "routes.Request": {
            "type": "object",
            "properties": {
                "reverse": {
                    "description": "Reverse   reverse Text",
                    "type": "boolean"
                },
                "sleepMillis": {
                    "description": "SleepMillis add some sleep to processing",
                    "type": "integer"
                },
                "text": {
                    "description": "Text is any text to process",
                    "type": "string"
                },
                "upperCase": {
                    "description": "Uppercase change the Text to uppercase",
                    "type": "boolean"
                }
            }
        },
        "routes.Response": {
            "type": "object",
            "properties": {
                "cloudId": {
                    "description": "CloudId the cloud which processed the request",
                    "type": "string"
                },
                "requestId": {
                    "description": "RequestId the request id",
                    "type": "string"
                },
                "text": {
                    "description": "Text the processed text with all applied transformations",
                    "type": "string"
                },
                "workerId": {
                    "description": "WorkerId the worker id",
                    "type": "string"
                }
            }
        },
        "utils.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/v1/api",
	Schemes:     []string{"http", "https"},
	Title:       "Hybrid Cloud Demo Backend API",
	Description: "The backend API that processes the message from frontend",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
