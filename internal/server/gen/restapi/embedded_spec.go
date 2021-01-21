// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Miasma",
    "version": "0.0.0"
  },
  "paths": {
    "/api/apps": {
      "get": {
        "summary": "List all the running applications",
        "operationId": "getApps",
        "parameters": [
          {
            "type": "boolean",
            "description": "Whether or not to show hidden apps",
            "name": "hidden",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/App"
              }
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "summary": "Create and start a new application",
        "operationId": "createApp",
        "parameters": [
          {
            "name": "app",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/App"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/App"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/api/apps/{appName}": {
      "get": {
        "summary": "Get an application by name",
        "operationId": "getApp",
        "parameters": [
          {
            "$ref": "#/parameters/appName"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/App"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "delete": {
        "summary": "Stop and delete an application",
        "operationId": "deleteApp",
        "parameters": [
          {
            "$ref": "#/parameters/appName"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/App"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/api/apps/{appName}/start": {
      "post": {
        "summary": "start the application",
        "operationId": "startApp",
        "parameters": [
          {
            "$ref": "#/parameters/appName"
          }
        ],
        "responses": {
          "200": {
            "description": "Started"
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "$ref": "#/responses/unknown"
          }
        }
      }
    },
    "/api/apps/{appName}/stop": {
      "post": {
        "summary": "stop the application",
        "operationId": "stopApp",
        "parameters": [
          {
            "$ref": "#/parameters/appName"
          }
        ],
        "responses": {
          "200": {
            "description": "Stopped"
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "$ref": "#/responses/unknown"
          }
        }
      }
    },
    "/api/health": {
      "get": {
        "summary": "Standard health check endpoint that checks all the service's statuses",
        "operationId": "getHealthCheck",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "$ref": "#/definitions/Health"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "App": {
      "type": "object",
      "required": [
        "name",
        "image",
        "running"
      ],
      "properties": {
        "hidden": {
          "description": "Whether or not the app is hidden during regular requests",
          "type": "boolean"
        },
        "image": {
          "description": "The image the app is based off of",
          "type": "string"
        },
        "name": {
          "description": "The apps name, used in the CLI with the ` + "`" + `-a|--app` + "`" + ` flag",
          "type": "string"
        },
        "running": {
          "type": "boolean"
        }
      }
    },
    "Health": {
      "type": "object",
      "required": [
        "version",
        "dockerVersion",
        "swarm"
      ],
      "properties": {
        "dockerVersion": {
          "description": "The version of docker running on the host, or null if docker is not running",
          "type": "string"
        },
        "swarm": {
          "description": "The info about the docker swarm if the host running miasma is apart of one. If it is not apart of a swarm, it returns ` + "`" + `null` + "`" + `",
          "type": "object",
          "properties": {
            "createdAt": {
              "description": "UTC timestamps when the swarm was created",
              "type": "string"
            },
            "id": {
              "description": "The swarm's ID",
              "type": "string"
            },
            "joinCommand": {
              "description": "The command for a node to run to join the swarm",
              "type": "string"
            },
            "updatedAt": {
              "description": "UTC timestamps when the swarm was last updated",
              "type": "string"
            }
          }
        },
        "version": {
          "description": "Miasma's current version",
          "type": "string"
        }
      }
    }
  },
  "parameters": {
    "appName": {
      "type": "string",
      "description": "App name from the ` + "`" + `-a|--app` + "`" + ` flag",
      "name": "appName",
      "in": "path",
      "required": true
    }
  },
  "responses": {
    "unknown": {
      "description": "Unknown Error",
      "schema": {
        "type": "string"
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Miasma",
    "version": "0.0.0"
  },
  "paths": {
    "/api/apps": {
      "get": {
        "summary": "List all the running applications",
        "operationId": "getApps",
        "parameters": [
          {
            "type": "boolean",
            "description": "Whether or not to show hidden apps",
            "name": "hidden",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/App"
              }
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "summary": "Create and start a new application",
        "operationId": "createApp",
        "parameters": [
          {
            "name": "app",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/App"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/App"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/api/apps/{appName}": {
      "get": {
        "summary": "Get an application by name",
        "operationId": "getApp",
        "parameters": [
          {
            "type": "string",
            "description": "App name from the ` + "`" + `-a|--app` + "`" + ` flag",
            "name": "appName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/App"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "delete": {
        "summary": "Stop and delete an application",
        "operationId": "deleteApp",
        "parameters": [
          {
            "type": "string",
            "description": "App name from the ` + "`" + `-a|--app` + "`" + ` flag",
            "name": "appName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/App"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/api/apps/{appName}/start": {
      "post": {
        "summary": "start the application",
        "operationId": "startApp",
        "parameters": [
          {
            "type": "string",
            "description": "App name from the ` + "`" + `-a|--app` + "`" + ` flag",
            "name": "appName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Started"
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/api/apps/{appName}/stop": {
      "post": {
        "summary": "stop the application",
        "operationId": "stopApp",
        "parameters": [
          {
            "type": "string",
            "description": "App name from the ` + "`" + `-a|--app` + "`" + ` flag",
            "name": "appName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Stopped"
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/api/health": {
      "get": {
        "summary": "Standard health check endpoint that checks all the service's statuses",
        "operationId": "getHealthCheck",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "$ref": "#/definitions/Health"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "App": {
      "type": "object",
      "required": [
        "name",
        "image",
        "running"
      ],
      "properties": {
        "hidden": {
          "description": "Whether or not the app is hidden during regular requests",
          "type": "boolean"
        },
        "image": {
          "description": "The image the app is based off of",
          "type": "string"
        },
        "name": {
          "description": "The apps name, used in the CLI with the ` + "`" + `-a|--app` + "`" + ` flag",
          "type": "string"
        },
        "running": {
          "type": "boolean"
        }
      }
    },
    "Health": {
      "type": "object",
      "required": [
        "version",
        "dockerVersion",
        "swarm"
      ],
      "properties": {
        "dockerVersion": {
          "description": "The version of docker running on the host, or null if docker is not running",
          "type": "string"
        },
        "swarm": {
          "description": "The info about the docker swarm if the host running miasma is apart of one. If it is not apart of a swarm, it returns ` + "`" + `null` + "`" + `",
          "type": "object",
          "properties": {
            "createdAt": {
              "description": "UTC timestamps when the swarm was created",
              "type": "string"
            },
            "id": {
              "description": "The swarm's ID",
              "type": "string"
            },
            "joinCommand": {
              "description": "The command for a node to run to join the swarm",
              "type": "string"
            },
            "updatedAt": {
              "description": "UTC timestamps when the swarm was last updated",
              "type": "string"
            }
          }
        },
        "version": {
          "description": "Miasma's current version",
          "type": "string"
        }
      }
    },
    "HealthSwarm": {
      "description": "The info about the docker swarm if the host running miasma is apart of one. If it is not apart of a swarm, it returns ` + "`" + `null` + "`" + `",
      "type": "object",
      "properties": {
        "createdAt": {
          "description": "UTC timestamps when the swarm was created",
          "type": "string"
        },
        "id": {
          "description": "The swarm's ID",
          "type": "string"
        },
        "joinCommand": {
          "description": "The command for a node to run to join the swarm",
          "type": "string"
        },
        "updatedAt": {
          "description": "UTC timestamps when the swarm was last updated",
          "type": "string"
        }
      }
    }
  },
  "parameters": {
    "appName": {
      "type": "string",
      "description": "App name from the ` + "`" + `-a|--app` + "`" + ` flag",
      "name": "appName",
      "in": "path",
      "required": true
    }
  },
  "responses": {
    "unknown": {
      "description": "Unknown Error",
      "schema": {
        "type": "string"
      }
    }
  }
}`))
}
