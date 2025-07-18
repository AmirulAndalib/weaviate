//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// TenantsGetOneReader is a Reader for the TenantsGetOne structure.
type TenantsGetOneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenantsGetOneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenantsGetOneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewTenantsGetOneUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTenantsGetOneForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTenantsGetOneNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewTenantsGetOneUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTenantsGetOneInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTenantsGetOneOK creates a TenantsGetOneOK with default headers values
func NewTenantsGetOneOK() *TenantsGetOneOK {
	return &TenantsGetOneOK{}
}

/*
TenantsGetOneOK describes a response with status code 200, with default header values.

load the tenant given the specified class
*/
type TenantsGetOneOK struct {
	Payload *models.Tenant
}

// IsSuccess returns true when this tenants get one o k response has a 2xx status code
func (o *TenantsGetOneOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenants get one o k response has a 3xx status code
func (o *TenantsGetOneOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenants get one o k response has a 4xx status code
func (o *TenantsGetOneOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenants get one o k response has a 5xx status code
func (o *TenantsGetOneOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenants get one o k response a status code equal to that given
func (o *TenantsGetOneOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tenants get one o k response
func (o *TenantsGetOneOK) Code() int {
	return 200
}

func (o *TenantsGetOneOK) Error() string {
	return fmt.Sprintf("[GET /schema/{className}/tenants/{tenantName}][%d] tenantsGetOneOK  %+v", 200, o.Payload)
}

func (o *TenantsGetOneOK) String() string {
	return fmt.Sprintf("[GET /schema/{className}/tenants/{tenantName}][%d] tenantsGetOneOK  %+v", 200, o.Payload)
}

func (o *TenantsGetOneOK) GetPayload() *models.Tenant {
	return o.Payload
}

func (o *TenantsGetOneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenantsGetOneUnauthorized creates a TenantsGetOneUnauthorized with default headers values
func NewTenantsGetOneUnauthorized() *TenantsGetOneUnauthorized {
	return &TenantsGetOneUnauthorized{}
}

/*
TenantsGetOneUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type TenantsGetOneUnauthorized struct {
}

// IsSuccess returns true when this tenants get one unauthorized response has a 2xx status code
func (o *TenantsGetOneUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tenants get one unauthorized response has a 3xx status code
func (o *TenantsGetOneUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenants get one unauthorized response has a 4xx status code
func (o *TenantsGetOneUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this tenants get one unauthorized response has a 5xx status code
func (o *TenantsGetOneUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this tenants get one unauthorized response a status code equal to that given
func (o *TenantsGetOneUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the tenants get one unauthorized response
func (o *TenantsGetOneUnauthorized) Code() int {
	return 401
}

func (o *TenantsGetOneUnauthorized) Error() string {
	return fmt.Sprintf("[GET /schema/{className}/tenants/{tenantName}][%d] tenantsGetOneUnauthorized ", 401)
}

func (o *TenantsGetOneUnauthorized) String() string {
	return fmt.Sprintf("[GET /schema/{className}/tenants/{tenantName}][%d] tenantsGetOneUnauthorized ", 401)
}

func (o *TenantsGetOneUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTenantsGetOneForbidden creates a TenantsGetOneForbidden with default headers values
func NewTenantsGetOneForbidden() *TenantsGetOneForbidden {
	return &TenantsGetOneForbidden{}
}

/*
TenantsGetOneForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TenantsGetOneForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this tenants get one forbidden response has a 2xx status code
func (o *TenantsGetOneForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tenants get one forbidden response has a 3xx status code
func (o *TenantsGetOneForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenants get one forbidden response has a 4xx status code
func (o *TenantsGetOneForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this tenants get one forbidden response has a 5xx status code
func (o *TenantsGetOneForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this tenants get one forbidden response a status code equal to that given
func (o *TenantsGetOneForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the tenants get one forbidden response
func (o *TenantsGetOneForbidden) Code() int {
	return 403
}

func (o *TenantsGetOneForbidden) Error() string {
	return fmt.Sprintf("[GET /schema/{className}/tenants/{tenantName}][%d] tenantsGetOneForbidden  %+v", 403, o.Payload)
}

func (o *TenantsGetOneForbidden) String() string {
	return fmt.Sprintf("[GET /schema/{className}/tenants/{tenantName}][%d] tenantsGetOneForbidden  %+v", 403, o.Payload)
}

func (o *TenantsGetOneForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TenantsGetOneForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenantsGetOneNotFound creates a TenantsGetOneNotFound with default headers values
func NewTenantsGetOneNotFound() *TenantsGetOneNotFound {
	return &TenantsGetOneNotFound{}
}

/*
TenantsGetOneNotFound describes a response with status code 404, with default header values.

Tenant not found
*/
type TenantsGetOneNotFound struct {
}

// IsSuccess returns true when this tenants get one not found response has a 2xx status code
func (o *TenantsGetOneNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tenants get one not found response has a 3xx status code
func (o *TenantsGetOneNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenants get one not found response has a 4xx status code
func (o *TenantsGetOneNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this tenants get one not found response has a 5xx status code
func (o *TenantsGetOneNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this tenants get one not found response a status code equal to that given
func (o *TenantsGetOneNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the tenants get one not found response
func (o *TenantsGetOneNotFound) Code() int {
	return 404
}

func (o *TenantsGetOneNotFound) Error() string {
	return fmt.Sprintf("[GET /schema/{className}/tenants/{tenantName}][%d] tenantsGetOneNotFound ", 404)
}

func (o *TenantsGetOneNotFound) String() string {
	return fmt.Sprintf("[GET /schema/{className}/tenants/{tenantName}][%d] tenantsGetOneNotFound ", 404)
}

func (o *TenantsGetOneNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTenantsGetOneUnprocessableEntity creates a TenantsGetOneUnprocessableEntity with default headers values
func NewTenantsGetOneUnprocessableEntity() *TenantsGetOneUnprocessableEntity {
	return &TenantsGetOneUnprocessableEntity{}
}

/*
TenantsGetOneUnprocessableEntity describes a response with status code 422, with default header values.

Invalid tenant or class
*/
type TenantsGetOneUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this tenants get one unprocessable entity response has a 2xx status code
func (o *TenantsGetOneUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tenants get one unprocessable entity response has a 3xx status code
func (o *TenantsGetOneUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenants get one unprocessable entity response has a 4xx status code
func (o *TenantsGetOneUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this tenants get one unprocessable entity response has a 5xx status code
func (o *TenantsGetOneUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this tenants get one unprocessable entity response a status code equal to that given
func (o *TenantsGetOneUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the tenants get one unprocessable entity response
func (o *TenantsGetOneUnprocessableEntity) Code() int {
	return 422
}

func (o *TenantsGetOneUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /schema/{className}/tenants/{tenantName}][%d] tenantsGetOneUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *TenantsGetOneUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /schema/{className}/tenants/{tenantName}][%d] tenantsGetOneUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *TenantsGetOneUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TenantsGetOneUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenantsGetOneInternalServerError creates a TenantsGetOneInternalServerError with default headers values
func NewTenantsGetOneInternalServerError() *TenantsGetOneInternalServerError {
	return &TenantsGetOneInternalServerError{}
}

/*
TenantsGetOneInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type TenantsGetOneInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this tenants get one internal server error response has a 2xx status code
func (o *TenantsGetOneInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tenants get one internal server error response has a 3xx status code
func (o *TenantsGetOneInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenants get one internal server error response has a 4xx status code
func (o *TenantsGetOneInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenants get one internal server error response has a 5xx status code
func (o *TenantsGetOneInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this tenants get one internal server error response a status code equal to that given
func (o *TenantsGetOneInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the tenants get one internal server error response
func (o *TenantsGetOneInternalServerError) Code() int {
	return 500
}

func (o *TenantsGetOneInternalServerError) Error() string {
	return fmt.Sprintf("[GET /schema/{className}/tenants/{tenantName}][%d] tenantsGetOneInternalServerError  %+v", 500, o.Payload)
}

func (o *TenantsGetOneInternalServerError) String() string {
	return fmt.Sprintf("[GET /schema/{className}/tenants/{tenantName}][%d] tenantsGetOneInternalServerError  %+v", 500, o.Payload)
}

func (o *TenantsGetOneInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TenantsGetOneInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
