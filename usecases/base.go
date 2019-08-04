package usecases

import (

)

// UseCase is an interface which standardises the methods available on
// Use Cases.
type UseCase interface {
	Setup()
	Execute()
}

// ResponseError contains the Name (e.g. ARTICLE_NOT_FOUND) and Description (e.g. 
// "The article could not be found in the repository.")
type ResponseError struct {
	Name string
	Description string
}

// Response contains a string mapping Body which contains some output data which is
// only exists on successful execution of the use case.
type Response struct {
	Body map[string]interface{}
}

// ResponseCollector is the object which the use case calls to in order to set
// the response or error.
type ResponseCollector struct {
	Response *Response
	Error *ResponseError
}

// SetResponse sets the Response field on the ResponseCollector
func (rc *ResponseCollector) SetResponse(resp *Response) {
	rc.Response = resp
	rc.Error = nil
}

// SetError sets the Error field on the ResponseCollector.
func (rc *ResponseCollector) SetError(err *ResponseError) {
	rc.Response = nil
	rc.Error = err
}

func panicHandler(response *ResponseCollector) {
	if err := recover(); err != nil {
		respErr := ResponseError{
			Name: "SEVERE_FAILURE",
			Description: "Panic occured.",
		}
		response.SetError(&respErr)
	}
}