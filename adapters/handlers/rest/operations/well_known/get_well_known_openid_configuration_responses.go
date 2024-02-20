// Code generated by go-swagger; DO NOT EDIT.

package well_known

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// GetWellKnownOpenidConfigurationOKCode is the HTTP code returned for type GetWellKnownOpenidConfigurationOK
const GetWellKnownOpenidConfigurationOKCode int = 200

/*
GetWellKnownOpenidConfigurationOK Successful response, inspect body

swagger:response getWellKnownOpenidConfigurationOK
*/
type GetWellKnownOpenidConfigurationOK struct {

	/*
	  In: Body
	*/
	Payload *GetWellKnownOpenidConfigurationOKBody `json:"body,omitempty"`
}

// NewGetWellKnownOpenidConfigurationOK creates GetWellKnownOpenidConfigurationOK with default headers values
func NewGetWellKnownOpenidConfigurationOK() *GetWellKnownOpenidConfigurationOK {

	return &GetWellKnownOpenidConfigurationOK{}
}

// WithPayload adds the payload to the get well known openid configuration o k response
func (o *GetWellKnownOpenidConfigurationOK) WithPayload(payload *GetWellKnownOpenidConfigurationOKBody) *GetWellKnownOpenidConfigurationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get well known openid configuration o k response
func (o *GetWellKnownOpenidConfigurationOK) SetPayload(payload *GetWellKnownOpenidConfigurationOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWellKnownOpenidConfigurationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetWellKnownOpenidConfigurationNotFoundCode is the HTTP code returned for type GetWellKnownOpenidConfigurationNotFound
const GetWellKnownOpenidConfigurationNotFoundCode int = 404

/*
GetWellKnownOpenidConfigurationNotFound Not found, no oidc provider present

swagger:response getWellKnownOpenidConfigurationNotFound
*/
type GetWellKnownOpenidConfigurationNotFound struct {
}

// NewGetWellKnownOpenidConfigurationNotFound creates GetWellKnownOpenidConfigurationNotFound with default headers values
func NewGetWellKnownOpenidConfigurationNotFound() *GetWellKnownOpenidConfigurationNotFound {

	return &GetWellKnownOpenidConfigurationNotFound{}
}

// WriteResponse to the client
func (o *GetWellKnownOpenidConfigurationNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetWellKnownOpenidConfigurationInternalServerErrorCode is the HTTP code returned for type GetWellKnownOpenidConfigurationInternalServerError
const GetWellKnownOpenidConfigurationInternalServerErrorCode int = 500

/*
GetWellKnownOpenidConfigurationInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response getWellKnownOpenidConfigurationInternalServerError
*/
type GetWellKnownOpenidConfigurationInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetWellKnownOpenidConfigurationInternalServerError creates GetWellKnownOpenidConfigurationInternalServerError with default headers values
func NewGetWellKnownOpenidConfigurationInternalServerError() *GetWellKnownOpenidConfigurationInternalServerError {

	return &GetWellKnownOpenidConfigurationInternalServerError{}
}

// WithPayload adds the payload to the get well known openid configuration internal server error response
func (o *GetWellKnownOpenidConfigurationInternalServerError) WithPayload(payload *models.ErrorResponse) *GetWellKnownOpenidConfigurationInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get well known openid configuration internal server error response
func (o *GetWellKnownOpenidConfigurationInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWellKnownOpenidConfigurationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
