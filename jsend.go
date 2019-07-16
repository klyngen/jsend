package jsend

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

/**
* GOAL of these utils, is to ensure unity in the API
 */

// FormatResponse formats a response after the jsend-standard
func FormatResponse(w http.ResponseWriter, data interface{}, status ResponseStatus) error {

	var result *responseResult

	switch status {
	// We want the result to be 0 anyway
	case NoContent:
		// Closest we get to nil
		result = nil
	// Hdndle the 400-series
	case NotFound:
		fallthrough
	case BadRequest:
		fallthrough
	case UnAuthorized:
		fallthrough
	case MethodNotAllowed:
		fallthrough
	case ServiceNotAvailable:
		fallthrough
	case InternalServerError:
		fallthrough
	case Forbidden:
		result = &responseResult{
			Status: status.Status,
		}

		if str, ok := data.(string); ok {
			result.Message = str
		} else {
			result.Message = "missing"
		}
	default:
		result = &responseResult{
			Data:   data,
			Status: status.Status,
		}
	}

	// 204
	if result == nil {
		w.WriteHeader(204)
		return nil
	}

	// This is obviously json...
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status.StatusCode)

	err := json.NewEncoder(w).Encode(result)

	// Lets not panic. Just return something nice
	if err != nil {
		// TODO add some logging here
		// Basic fallback printing a simple status-message
		fmt.Fprint(w, "{"+
			"'status': 'success',"+
			"'message': 'internalServerError'"+
			"}")
		return errors.New("Unable to marshal request")
	}

	return nil
}

type responseResult struct {
	Data    interface{} `json:"data,omitempty"`
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
}

// ResponseStatus refers to a status. It has a simple status-code and a simple status-message
type ResponseStatus struct {
	Status     string
	StatusCode int
}

// ################ 200's ########################################
var (
	// Success type response
	Success = ResponseStatus{Status: "success", StatusCode: 200}
	// NoContent successfull but nothing to report
	NoContent = ResponseStatus{Status: "success", StatusCode: 204}
)

// ################ 400's ########################################
var (
	// NotFound what ever the query was... We could not find it...
	NotFound = ResponseStatus{Status: "client-error", StatusCode: 404}
	// BadRequest symbolises a developer / user doing crazy shit
	BadRequest = ResponseStatus{Status: "client-error", StatusCode: 400}
	// UnAuthorized symbolises a developer / user doing illegal shit
	UnAuthorized = ResponseStatus{Status: "client-error", StatusCode: 401}
	// Forbidden symbolises a developer / user doing illegal shit
	Forbidden = ResponseStatus{Status: "client-error", StatusCode: 403}
	// MethodNotAllowed symbolises a developer / user doing crazy shit
	MethodNotAllowed = ResponseStatus{Status: "client-error", StatusCode: 405}
)

// ################ 500's ########################################
var (
	// ServiceNotAvailable means that we might not have that service for some reason...
	ServiceNotAvailable = ResponseStatus{Status: "server-error", StatusCode: 503}
	// InternalServerError means that I am having a bad day...
	InternalServerError = ResponseStatus{Status: "server-error", StatusCode: 500}
)
