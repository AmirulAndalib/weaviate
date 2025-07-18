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

package authz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/weaviate/weaviate/entities/models"
)

// AssignRoleToUserReader is a Reader for the AssignRoleToUser structure.
type AssignRoleToUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignRoleToUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssignRoleToUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAssignRoleToUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAssignRoleToUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAssignRoleToUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAssignRoleToUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAssignRoleToUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAssignRoleToUserOK creates a AssignRoleToUserOK with default headers values
func NewAssignRoleToUserOK() *AssignRoleToUserOK {
	return &AssignRoleToUserOK{}
}

/*
AssignRoleToUserOK describes a response with status code 200, with default header values.

Role assigned successfully
*/
type AssignRoleToUserOK struct {
}

// IsSuccess returns true when this assign role to user o k response has a 2xx status code
func (o *AssignRoleToUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this assign role to user o k response has a 3xx status code
func (o *AssignRoleToUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign role to user o k response has a 4xx status code
func (o *AssignRoleToUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this assign role to user o k response has a 5xx status code
func (o *AssignRoleToUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this assign role to user o k response a status code equal to that given
func (o *AssignRoleToUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the assign role to user o k response
func (o *AssignRoleToUserOK) Code() int {
	return 200
}

func (o *AssignRoleToUserOK) Error() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleToUserOK ", 200)
}

func (o *AssignRoleToUserOK) String() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleToUserOK ", 200)
}

func (o *AssignRoleToUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignRoleToUserBadRequest creates a AssignRoleToUserBadRequest with default headers values
func NewAssignRoleToUserBadRequest() *AssignRoleToUserBadRequest {
	return &AssignRoleToUserBadRequest{}
}

/*
AssignRoleToUserBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type AssignRoleToUserBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this assign role to user bad request response has a 2xx status code
func (o *AssignRoleToUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign role to user bad request response has a 3xx status code
func (o *AssignRoleToUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign role to user bad request response has a 4xx status code
func (o *AssignRoleToUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this assign role to user bad request response has a 5xx status code
func (o *AssignRoleToUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this assign role to user bad request response a status code equal to that given
func (o *AssignRoleToUserBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the assign role to user bad request response
func (o *AssignRoleToUserBadRequest) Code() int {
	return 400
}

func (o *AssignRoleToUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleToUserBadRequest  %+v", 400, o.Payload)
}

func (o *AssignRoleToUserBadRequest) String() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleToUserBadRequest  %+v", 400, o.Payload)
}

func (o *AssignRoleToUserBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AssignRoleToUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignRoleToUserUnauthorized creates a AssignRoleToUserUnauthorized with default headers values
func NewAssignRoleToUserUnauthorized() *AssignRoleToUserUnauthorized {
	return &AssignRoleToUserUnauthorized{}
}

/*
AssignRoleToUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type AssignRoleToUserUnauthorized struct {
}

// IsSuccess returns true when this assign role to user unauthorized response has a 2xx status code
func (o *AssignRoleToUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign role to user unauthorized response has a 3xx status code
func (o *AssignRoleToUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign role to user unauthorized response has a 4xx status code
func (o *AssignRoleToUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this assign role to user unauthorized response has a 5xx status code
func (o *AssignRoleToUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this assign role to user unauthorized response a status code equal to that given
func (o *AssignRoleToUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the assign role to user unauthorized response
func (o *AssignRoleToUserUnauthorized) Code() int {
	return 401
}

func (o *AssignRoleToUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleToUserUnauthorized ", 401)
}

func (o *AssignRoleToUserUnauthorized) String() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleToUserUnauthorized ", 401)
}

func (o *AssignRoleToUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignRoleToUserForbidden creates a AssignRoleToUserForbidden with default headers values
func NewAssignRoleToUserForbidden() *AssignRoleToUserForbidden {
	return &AssignRoleToUserForbidden{}
}

/*
AssignRoleToUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AssignRoleToUserForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this assign role to user forbidden response has a 2xx status code
func (o *AssignRoleToUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign role to user forbidden response has a 3xx status code
func (o *AssignRoleToUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign role to user forbidden response has a 4xx status code
func (o *AssignRoleToUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this assign role to user forbidden response has a 5xx status code
func (o *AssignRoleToUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this assign role to user forbidden response a status code equal to that given
func (o *AssignRoleToUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the assign role to user forbidden response
func (o *AssignRoleToUserForbidden) Code() int {
	return 403
}

func (o *AssignRoleToUserForbidden) Error() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleToUserForbidden  %+v", 403, o.Payload)
}

func (o *AssignRoleToUserForbidden) String() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleToUserForbidden  %+v", 403, o.Payload)
}

func (o *AssignRoleToUserForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AssignRoleToUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignRoleToUserNotFound creates a AssignRoleToUserNotFound with default headers values
func NewAssignRoleToUserNotFound() *AssignRoleToUserNotFound {
	return &AssignRoleToUserNotFound{}
}

/*
AssignRoleToUserNotFound describes a response with status code 404, with default header values.

role or user is not found.
*/
type AssignRoleToUserNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this assign role to user not found response has a 2xx status code
func (o *AssignRoleToUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign role to user not found response has a 3xx status code
func (o *AssignRoleToUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign role to user not found response has a 4xx status code
func (o *AssignRoleToUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this assign role to user not found response has a 5xx status code
func (o *AssignRoleToUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this assign role to user not found response a status code equal to that given
func (o *AssignRoleToUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the assign role to user not found response
func (o *AssignRoleToUserNotFound) Code() int {
	return 404
}

func (o *AssignRoleToUserNotFound) Error() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleToUserNotFound  %+v", 404, o.Payload)
}

func (o *AssignRoleToUserNotFound) String() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleToUserNotFound  %+v", 404, o.Payload)
}

func (o *AssignRoleToUserNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AssignRoleToUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignRoleToUserInternalServerError creates a AssignRoleToUserInternalServerError with default headers values
func NewAssignRoleToUserInternalServerError() *AssignRoleToUserInternalServerError {
	return &AssignRoleToUserInternalServerError{}
}

/*
AssignRoleToUserInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type AssignRoleToUserInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this assign role to user internal server error response has a 2xx status code
func (o *AssignRoleToUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign role to user internal server error response has a 3xx status code
func (o *AssignRoleToUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign role to user internal server error response has a 4xx status code
func (o *AssignRoleToUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this assign role to user internal server error response has a 5xx status code
func (o *AssignRoleToUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this assign role to user internal server error response a status code equal to that given
func (o *AssignRoleToUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the assign role to user internal server error response
func (o *AssignRoleToUserInternalServerError) Code() int {
	return 500
}

func (o *AssignRoleToUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleToUserInternalServerError  %+v", 500, o.Payload)
}

func (o *AssignRoleToUserInternalServerError) String() string {
	return fmt.Sprintf("[POST /authz/users/{id}/assign][%d] assignRoleToUserInternalServerError  %+v", 500, o.Payload)
}

func (o *AssignRoleToUserInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AssignRoleToUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AssignRoleToUserBody assign role to user body
swagger:model AssignRoleToUserBody
*/
type AssignRoleToUserBody struct {

	// the roles that assigned to user
	Roles []string `json:"roles"`

	// user type
	UserType models.UserTypeInput `json:"userType,omitempty"`
}

// Validate validates this assign role to user body
func (o *AssignRoleToUserBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateUserType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AssignRoleToUserBody) validateUserType(formats strfmt.Registry) error {
	if swag.IsZero(o.UserType) { // not required
		return nil
	}

	if err := o.UserType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "userType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "userType")
		}
		return err
	}

	return nil
}

// ContextValidate validate this assign role to user body based on the context it is used
func (o *AssignRoleToUserBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateUserType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AssignRoleToUserBody) contextValidateUserType(ctx context.Context, formats strfmt.Registry) error {

	if err := o.UserType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "userType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "userType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AssignRoleToUserBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AssignRoleToUserBody) UnmarshalBinary(b []byte) error {
	var res AssignRoleToUserBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
