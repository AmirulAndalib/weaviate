// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/entities/models"
)

// ObjectsClassHeadNoContentCode is the HTTP code returned for type ObjectsClassHeadNoContent
const ObjectsClassHeadNoContentCode int = 204

/*
ObjectsClassHeadNoContent Object exists.

swagger:response objectsClassHeadNoContent
*/
type ObjectsClassHeadNoContent struct {
}

// NewObjectsClassHeadNoContent creates ObjectsClassHeadNoContent with default headers values
func NewObjectsClassHeadNoContent() *ObjectsClassHeadNoContent {

	return &ObjectsClassHeadNoContent{}
}

// WriteResponse to the client
func (o *ObjectsClassHeadNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// ObjectsClassHeadUnauthorizedCode is the HTTP code returned for type ObjectsClassHeadUnauthorized
const ObjectsClassHeadUnauthorizedCode int = 401

/*
ObjectsClassHeadUnauthorized Unauthorized or invalid credentials.

swagger:response objectsClassHeadUnauthorized
*/
type ObjectsClassHeadUnauthorized struct {
}

// NewObjectsClassHeadUnauthorized creates ObjectsClassHeadUnauthorized with default headers values
func NewObjectsClassHeadUnauthorized() *ObjectsClassHeadUnauthorized {

	return &ObjectsClassHeadUnauthorized{}
}

// WriteResponse to the client
func (o *ObjectsClassHeadUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ObjectsClassHeadForbiddenCode is the HTTP code returned for type ObjectsClassHeadForbidden
const ObjectsClassHeadForbiddenCode int = 403

/*
ObjectsClassHeadForbidden Forbidden

swagger:response objectsClassHeadForbidden
*/
type ObjectsClassHeadForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassHeadForbidden creates ObjectsClassHeadForbidden with default headers values
func NewObjectsClassHeadForbidden() *ObjectsClassHeadForbidden {

	return &ObjectsClassHeadForbidden{}
}

// WithPayload adds the payload to the objects class head forbidden response
func (o *ObjectsClassHeadForbidden) WithPayload(payload *models.ErrorResponse) *ObjectsClassHeadForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class head forbidden response
func (o *ObjectsClassHeadForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassHeadForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsClassHeadNotFoundCode is the HTTP code returned for type ObjectsClassHeadNotFound
const ObjectsClassHeadNotFoundCode int = 404

/*
ObjectsClassHeadNotFound Object doesn't exist.

swagger:response objectsClassHeadNotFound
*/
type ObjectsClassHeadNotFound struct {
}

// NewObjectsClassHeadNotFound creates ObjectsClassHeadNotFound with default headers values
func NewObjectsClassHeadNotFound() *ObjectsClassHeadNotFound {

	return &ObjectsClassHeadNotFound{}
}

// WriteResponse to the client
func (o *ObjectsClassHeadNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ObjectsClassHeadUnprocessableEntityCode is the HTTP code returned for type ObjectsClassHeadUnprocessableEntity
const ObjectsClassHeadUnprocessableEntityCode int = 422

/*
ObjectsClassHeadUnprocessableEntity Request is well-formed (i.e., syntactically correct), but erroneous.

swagger:response objectsClassHeadUnprocessableEntity
*/
type ObjectsClassHeadUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassHeadUnprocessableEntity creates ObjectsClassHeadUnprocessableEntity with default headers values
func NewObjectsClassHeadUnprocessableEntity() *ObjectsClassHeadUnprocessableEntity {

	return &ObjectsClassHeadUnprocessableEntity{}
}

// WithPayload adds the payload to the objects class head unprocessable entity response
func (o *ObjectsClassHeadUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *ObjectsClassHeadUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class head unprocessable entity response
func (o *ObjectsClassHeadUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassHeadUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ObjectsClassHeadInternalServerErrorCode is the HTTP code returned for type ObjectsClassHeadInternalServerError
const ObjectsClassHeadInternalServerErrorCode int = 500

/*
ObjectsClassHeadInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response objectsClassHeadInternalServerError
*/
type ObjectsClassHeadInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewObjectsClassHeadInternalServerError creates ObjectsClassHeadInternalServerError with default headers values
func NewObjectsClassHeadInternalServerError() *ObjectsClassHeadInternalServerError {

	return &ObjectsClassHeadInternalServerError{}
}

// WithPayload adds the payload to the objects class head internal server error response
func (o *ObjectsClassHeadInternalServerError) WithPayload(payload *models.ErrorResponse) *ObjectsClassHeadInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the objects class head internal server error response
func (o *ObjectsClassHeadInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ObjectsClassHeadInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
