/*
 * Swagger Petstore - OpenAPI 3.0
 *
 * 社員情報管理API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// EmployeeAPIController binds http requests to an api service and writes the service results to the http response
type EmployeeAPIController struct {
	service EmployeeAPIServicer
	errorHandler ErrorHandler
}

// EmployeeAPIOption for how the controller is set up.
type EmployeeAPIOption func(*EmployeeAPIController)

// WithEmployeeAPIErrorHandler inject ErrorHandler into controller
func WithEmployeeAPIErrorHandler(h ErrorHandler) EmployeeAPIOption {
	return func(c *EmployeeAPIController) {
		c.errorHandler = h
	}
}

// NewEmployeeAPIController creates a default api controller
func NewEmployeeAPIController(s EmployeeAPIServicer, opts ...EmployeeAPIOption) Router {
	controller := &EmployeeAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the EmployeeAPIController
func (c *EmployeeAPIController) Routes() Routes {
	return Routes{
		"CreateEmployee": Route{
			strings.ToUpper("Post"),
			"/api/v1/employees",
			c.CreateEmployee,
		},
		"DeleteEmployee": Route{
			strings.ToUpper("Delete"),
			"/api/v1/employees/{employeeId}",
			c.DeleteEmployee,
		},
		"GetEmployeeById": Route{
			strings.ToUpper("Get"),
			"/api/v1/employees/{employeeId}",
			c.GetEmployeeById,
		},
		"ListEmployees": Route{
			strings.ToUpper("Get"),
			"/api/v1/employees",
			c.ListEmployees,
		},
		"UpdateEmployee": Route{
			strings.ToUpper("Put"),
			"/api/v1/employees/{employeeId}",
			c.UpdateEmployee,
		},
	}
}

// CreateEmployee - Create a new employee
func (c *EmployeeAPIController) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	employeeParam := Employee{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&employeeParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertEmployeeRequired(employeeParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertEmployeeConstraints(employeeParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateEmployee(r.Context(), employeeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteEmployee - Delete an employee
func (c *EmployeeAPIController) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	employeeIdParam := params["employeeId"]
	if employeeIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"employeeId"}, nil)
		return
	}
	result, err := c.service.DeleteEmployee(r.Context(), employeeIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetEmployeeById - Find employee by ID
func (c *EmployeeAPIController) GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	employeeIdParam := params["employeeId"]
	if employeeIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"employeeId"}, nil)
		return
	}
	result, err := c.service.GetEmployeeById(r.Context(), employeeIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListEmployees - List all employees
func (c *EmployeeAPIController) ListEmployees(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.ListEmployees(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateEmployee - Update an existing employee
func (c *EmployeeAPIController) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	employeeIdParam := params["employeeId"]
	if employeeIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"employeeId"}, nil)
		return
	}
	employeeParam := Employee{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&employeeParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertEmployeeRequired(employeeParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertEmployeeConstraints(employeeParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateEmployee(r.Context(), employeeIdParam, employeeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
