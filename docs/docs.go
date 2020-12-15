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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/acceptCandidate": {
            "put": {
                "description": "Accepts a candidate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Accept candidate",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Candidate ID",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "qwerty"
                            }
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/arrangeMeeting": {
            "put": {
                "description": "Aranges a meeting with an available candidate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Arrange meeting",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Candidate ID",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Next Meeting Time (Ex. 2020-11-11T20:59:48.133+03:00)",
                        "name": "nextMeetingTime",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "qwerty"
                            }
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/completeMeeting": {
            "put": {
                "description": "Completes a meeting of specified candidate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Complete meeting",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Candidate ID",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "qwerty"
                            }
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/createCandidate": {
            "post": {
                "description": "Creates a new candidate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create candidate",
                "parameters": [
                    {
                        "type": "string",
                        "description": "First Name",
                        "name": "first_name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Last Name",
                        "name": "last_name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "E-Mail",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Department",
                        "name": "department",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "University",
                        "name": "university",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Experience",
                        "name": "experience",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Assignee",
                        "name": "assignee",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "qwerty"
                            }
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/deleteCandidate": {
            "delete": {
                "description": "Removes a candidate by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete candidate",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Candidate ID",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "qwerty"
                            }
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/denyCandidate": {
            "put": {
                "description": "Denies a candidate by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Deny candidate",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Candidate ID",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "qwerty"
                            }
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/findAssigneeIDByName": {
            "get": {
                "description": "Finds ID of an assignee by name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Find assignee ID by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Assignee Name",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "qwerty"
                            }
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/findAssigneesCandidates": {
            "get": {
                "description": "Finds candidates of an assignee by ID of assignee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Find asignees candidates",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Assignee ID",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "qwerty"
                            }
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/readCandidate": {
            "get": {
                "description": "Reads a specific candidate by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Read candidate",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Candidate ID",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        },
                        "headers": {
                            "Token": {
                                "type": "string",
                                "description": "qwerty"
                            }
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "object"
                        }
                    }
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
