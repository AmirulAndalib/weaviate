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

package graphql

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// GraphqlBatchReader is a Reader for the GraphqlBatch structure.
type GraphqlBatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GraphqlBatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGraphqlBatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGraphqlBatchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGraphqlBatchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGraphqlBatchUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGraphqlBatchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGraphqlBatchOK creates a GraphqlBatchOK with default headers values
func NewGraphqlBatchOK() *GraphqlBatchOK {
	return &GraphqlBatchOK{}
}

/*
GraphqlBatchOK describes a response with status code 200, with default header values.

Successful query (with select).
*/
type GraphqlBatchOK struct {
	Payload models.GraphQLResponses
}

// IsSuccess returns true when this graphql batch o k response has a 2xx status code
func (o *GraphqlBatchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this graphql batch o k response has a 3xx status code
func (o *GraphqlBatchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this graphql batch o k response has a 4xx status code
func (o *GraphqlBatchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this graphql batch o k response has a 5xx status code
func (o *GraphqlBatchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this graphql batch o k response a status code equal to that given
func (o *GraphqlBatchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the graphql batch o k response
func (o *GraphqlBatchOK) Code() int {
	return 200
}

func (o *GraphqlBatchOK) Error() string {
	return fmt.Sprintf("[POST /graphql/batch][%d] graphqlBatchOK  %+v", 200, o.Payload)
}

func (o *GraphqlBatchOK) String() string {
	return fmt.Sprintf("[POST /graphql/batch][%d] graphqlBatchOK  %+v", 200, o.Payload)
}

func (o *GraphqlBatchOK) GetPayload() models.GraphQLResponses {
	return o.Payload
}

func (o *GraphqlBatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGraphqlBatchUnauthorized creates a GraphqlBatchUnauthorized with default headers values
func NewGraphqlBatchUnauthorized() *GraphqlBatchUnauthorized {
	return &GraphqlBatchUnauthorized{}
}

/*
GraphqlBatchUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type GraphqlBatchUnauthorized struct {
}

// IsSuccess returns true when this graphql batch unauthorized response has a 2xx status code
func (o *GraphqlBatchUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this graphql batch unauthorized response has a 3xx status code
func (o *GraphqlBatchUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this graphql batch unauthorized response has a 4xx status code
func (o *GraphqlBatchUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this graphql batch unauthorized response has a 5xx status code
func (o *GraphqlBatchUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this graphql batch unauthorized response a status code equal to that given
func (o *GraphqlBatchUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the graphql batch unauthorized response
func (o *GraphqlBatchUnauthorized) Code() int {
	return 401
}

func (o *GraphqlBatchUnauthorized) Error() string {
	return fmt.Sprintf("[POST /graphql/batch][%d] graphqlBatchUnauthorized ", 401)
}

func (o *GraphqlBatchUnauthorized) String() string {
	return fmt.Sprintf("[POST /graphql/batch][%d] graphqlBatchUnauthorized ", 401)
}

func (o *GraphqlBatchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGraphqlBatchForbidden creates a GraphqlBatchForbidden with default headers values
func NewGraphqlBatchForbidden() *GraphqlBatchForbidden {
	return &GraphqlBatchForbidden{}
}

/*
GraphqlBatchForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GraphqlBatchForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this graphql batch forbidden response has a 2xx status code
func (o *GraphqlBatchForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this graphql batch forbidden response has a 3xx status code
func (o *GraphqlBatchForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this graphql batch forbidden response has a 4xx status code
func (o *GraphqlBatchForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this graphql batch forbidden response has a 5xx status code
func (o *GraphqlBatchForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this graphql batch forbidden response a status code equal to that given
func (o *GraphqlBatchForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the graphql batch forbidden response
func (o *GraphqlBatchForbidden) Code() int {
	return 403
}

func (o *GraphqlBatchForbidden) Error() string {
	return fmt.Sprintf("[POST /graphql/batch][%d] graphqlBatchForbidden  %+v", 403, o.Payload)
}

func (o *GraphqlBatchForbidden) String() string {
	return fmt.Sprintf("[POST /graphql/batch][%d] graphqlBatchForbidden  %+v", 403, o.Payload)
}

func (o *GraphqlBatchForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GraphqlBatchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGraphqlBatchUnprocessableEntity creates a GraphqlBatchUnprocessableEntity with default headers values
func NewGraphqlBatchUnprocessableEntity() *GraphqlBatchUnprocessableEntity {
	return &GraphqlBatchUnprocessableEntity{}
}

/*
GraphqlBatchUnprocessableEntity describes a response with status code 422, with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type GraphqlBatchUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this graphql batch unprocessable entity response has a 2xx status code
func (o *GraphqlBatchUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this graphql batch unprocessable entity response has a 3xx status code
func (o *GraphqlBatchUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this graphql batch unprocessable entity response has a 4xx status code
func (o *GraphqlBatchUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this graphql batch unprocessable entity response has a 5xx status code
func (o *GraphqlBatchUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this graphql batch unprocessable entity response a status code equal to that given
func (o *GraphqlBatchUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the graphql batch unprocessable entity response
func (o *GraphqlBatchUnprocessableEntity) Code() int {
	return 422
}

func (o *GraphqlBatchUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /graphql/batch][%d] graphqlBatchUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GraphqlBatchUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /graphql/batch][%d] graphqlBatchUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GraphqlBatchUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GraphqlBatchUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGraphqlBatchInternalServerError creates a GraphqlBatchInternalServerError with default headers values
func NewGraphqlBatchInternalServerError() *GraphqlBatchInternalServerError {
	return &GraphqlBatchInternalServerError{}
}

/*
GraphqlBatchInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type GraphqlBatchInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this graphql batch internal server error response has a 2xx status code
func (o *GraphqlBatchInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this graphql batch internal server error response has a 3xx status code
func (o *GraphqlBatchInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this graphql batch internal server error response has a 4xx status code
func (o *GraphqlBatchInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this graphql batch internal server error response has a 5xx status code
func (o *GraphqlBatchInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this graphql batch internal server error response a status code equal to that given
func (o *GraphqlBatchInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the graphql batch internal server error response
func (o *GraphqlBatchInternalServerError) Code() int {
	return 500
}

func (o *GraphqlBatchInternalServerError) Error() string {
	return fmt.Sprintf("[POST /graphql/batch][%d] graphqlBatchInternalServerError  %+v", 500, o.Payload)
}

func (o *GraphqlBatchInternalServerError) String() string {
	return fmt.Sprintf("[POST /graphql/batch][%d] graphqlBatchInternalServerError  %+v", 500, o.Payload)
}

func (o *GraphqlBatchInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GraphqlBatchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
