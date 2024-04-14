// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
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
        "/class": {
            "post": {
                "description": "create a new class",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Class"
                ],
                "summary": "CreateClass",
                "parameters": [
                    {
                        "description": "Class object",
                        "name": "class",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Class"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Class created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "some error message here (from err.Error())",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/class/:id": {
            "get": {
                "description": "get class search by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Class"
                ],
                "summary": "GetClassByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "class id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Class"
                        }
                    },
                    "404": {
                        "description": "some error message here (from err.Error())",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/class/course": {
            "get": {
                "description": "get class search by course_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Class"
                ],
                "summary": "GetClassByCourseID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "course id",
                        "name": "course_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Class"
                            }
                        }
                    },
                    "404": {
                        "description": "some error message here (from err.Error())",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/class/delete/:id": {
            "delete": {
                "description": "delete class by class_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Class"
                ],
                "summary": "DeleteClassByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "class_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Class deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "were not able to delete the class",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/class/semester": {
            "get": {
                "description": "get class search by semester",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Class"
                ],
                "summary": "GetClassBySemester",
                "parameters": [
                    {
                        "type": "string",
                        "description": "1, 2, ...",
                        "name": "semester",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "array of class_id",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    },
                    "404": {
                        "description": "some error message here (from err.Error())",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/course": {
            "get": {
                "description": "get all courses",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Course"
                ],
                "summary": "GetAllCourses",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Course"
                            }
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
            "post": {
                "description": "create a new course",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Course"
                ],
                "summary": "CreateCourse",
                "parameters": [
                    {
                        "description": "Course object",
                        "name": "program",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.Course"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Course created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "some error message here (from err.Error())",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/course/:id": {
            "get": {
                "description": "get a course by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Course"
                ],
                "summary": "GetCourseByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "program id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Course"
                        }
                    },
                    "404": {
                        "description": "some error message here (from err.Error())",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/course/delete/:id": {
            "delete": {
                "description": "delete a course by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Course"
                ],
                "summary": "DeleteCourseByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "course id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Course deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "some error message here (from err.Error())",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/course/section": {
            "get": {
                "description": "get section search by class_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Course"
                ],
                "summary": "GetSectionByClassID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "class id",
                        "name": "class_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "CourseCode, Section",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/service.GetSectionByClassIDField"
                            }
                        }
                    },
                    "404": {
                        "description": "some error message here (from err.Error())",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/course/semester": {
            "get": {
                "description": "get semester from all course [no duplicate]",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Course"
                ],
                "summary": "GetAllDistinctSemester",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    },
                    "404": {
                        "description": "some error message here (from err.Error())",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/course/update/:id": {
            "put": {
                "description": "update a course by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Course"
                ],
                "summary": "UpdateCourseByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "course id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Course updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "some error message here (from err.Error())",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/program": {
            "get": {
                "description": "get all programs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Program"
                ],
                "summary": "GetAllPrograms",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Program"
                            }
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
            "post": {
                "description": "create a new program",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Program"
                ],
                "summary": "CreateProgram",
                "parameters": [
                    {
                        "description": "Program object",
                        "name": "program",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Program"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/program/:id": {
            "get": {
                "description": "get a program by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Program"
                ],
                "summary": "GetProgramByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "program id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Program"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/program/delete/:id": {
            "delete": {
                "description": "delete a program by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Program"
                ],
                "summary": "DeleteProgramByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "program id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/program/update/:id": {
            "put": {
                "description": "update a program by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Program"
                ],
                "summary": "UpdateProgramByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "program id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/timetable": {
            "post": {
                "description": "create a new timetable",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Timetable"
                ],
                "summary": "CreateTimeTable",
                "parameters": [
                    {
                        "description": "Timetable object",
                        "name": "timetable",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Timetable"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Timetable created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "some error message here (from err.Error())",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/timetable/class": {
            "get": {
                "description": "get timetable search by class_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Timetable"
                ],
                "summary": "GetTimetableByClassID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "class id",
                        "name": "class_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/service.GetTimetableByClassIDField"
                            }
                        }
                    },
                    "404": {
                        "description": "No timetables found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/timetable/count": {
            "get": {
                "description": "get timetable search by class_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Timetable"
                ],
                "summary": "CountTimeTable",
                "parameters": [
                    {
                        "type": "string",
                        "description": "you add your own search field in timetable model here (check equal only) but in php version it was searched by class_id",
                        "name": "{field_name}",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Count of timetables",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "404": {
                        "description": "some error message here (from err.Error())",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/timetable/delete/:id": {
            "delete": {
                "description": "delete a timetable by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Timetable"
                ],
                "summary": "DeleteTimeTableByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "timetable id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Timetable deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "some error message here (from err.Error())",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "model.Class": {
            "type": "object",
            "properties": {
                "course": {
                    "$ref": "#/definitions/model.Course"
                },
                "courseID": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "section": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.Course": {
            "type": "object",
            "properties": {
                "courseCode": {
                    "type": "string"
                },
                "courseName": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "credit": {
                    "type": "number"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "detail": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "semester": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "model.Faculty": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "department": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "major": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.Program": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "detail": {
                    "type": "string"
                },
                "faculty": {
                    "$ref": "#/definitions/model.Faculty"
                },
                "facultyID": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "prefix": {
                    "type": "string"
                },
                "pricePerTerm": {
                    "type": "number"
                },
                "programName": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.Timetable": {
            "type": "object",
            "properties": {
                "class": {
                    "$ref": "#/definitions/model.Class"
                },
                "classID": {
                    "type": "integer"
                },
                "classType": {
                    "type": "string"
                },
                "classroom": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "day": {
                    "$ref": "#/definitions/time.Weekday"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "endTime": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "startTime": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "service.GetSectionByClassIDField": {
            "type": "object",
            "properties": {
                "courseCode": {
                    "type": "string"
                },
                "section": {
                    "type": "string"
                }
            }
        },
        "service.GetTimetableByClassIDField": {
            "type": "object",
            "properties": {
                "classType": {
                    "type": "string"
                },
                "classroom": {
                    "type": "string"
                },
                "courseCode": {
                    "type": "string"
                },
                "day": {
                    "$ref": "#/definitions/time.Weekday"
                },
                "endTime": {
                    "type": "string"
                },
                "section": {
                    "type": "string"
                },
                "startTime": {
                    "type": "string"
                }
            }
        },
        "time.Weekday": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3,
                4,
                5,
                6
            ],
            "x-enum-varnames": [
                "Sunday",
                "Monday",
                "Tuesday",
                "Wednesday",
                "Thursday",
                "Friday",
                "Saturday"
            ]
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
