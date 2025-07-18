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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ObjectsGetResponse objects get response
//
// swagger:model ObjectsGetResponse
type ObjectsGetResponse struct {
	Object

	// deprecations
	Deprecations []*Deprecation `json:"deprecations"`

	// result
	Result *ObjectsGetResponseAO2Result `json:"result,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ObjectsGetResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Object
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Object = aO0

	// AO1
	var dataAO1 struct {
		Deprecations []*Deprecation `json:"deprecations"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Deprecations = dataAO1.Deprecations

	// AO2
	var dataAO2 struct {
		Result *ObjectsGetResponseAO2Result `json:"result,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO2); err != nil {
		return err
	}

	m.Result = dataAO2.Result

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ObjectsGetResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.Object)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Deprecations []*Deprecation `json:"deprecations"`
	}

	dataAO1.Deprecations = m.Deprecations

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	var dataAO2 struct {
		Result *ObjectsGetResponseAO2Result `json:"result,omitempty"`
	}

	dataAO2.Result = m.Result

	jsonDataAO2, errAO2 := swag.WriteJSON(dataAO2)
	if errAO2 != nil {
		return nil, errAO2
	}
	_parts = append(_parts, jsonDataAO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this objects get response
func (m *ObjectsGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Object
	if err := m.Object.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeprecations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectsGetResponse) validateDeprecations(formats strfmt.Registry) error {

	if swag.IsZero(m.Deprecations) { // not required
		return nil
	}

	for i := 0; i < len(m.Deprecations); i++ {
		if swag.IsZero(m.Deprecations[i]) { // not required
			continue
		}

		if m.Deprecations[i] != nil {
			if err := m.Deprecations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deprecations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deprecations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ObjectsGetResponse) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this objects get response based on the context it is used
func (m *ObjectsGetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Object
	if err := m.Object.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeprecations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectsGetResponse) contextValidateDeprecations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Deprecations); i++ {

		if m.Deprecations[i] != nil {
			if err := m.Deprecations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deprecations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deprecations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ObjectsGetResponse) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if m.Result != nil {
		if err := m.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectsGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectsGetResponse) UnmarshalBinary(b []byte) error {
	var res ObjectsGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ObjectsGetResponseAO2Result Results for this specific Object.
//
// swagger:model ObjectsGetResponseAO2Result
type ObjectsGetResponseAO2Result struct {

	// errors
	Errors *ErrorResponse `json:"errors,omitempty"`

	// status
	// Enum: [SUCCESS FAILED]
	Status *string `json:"status,omitempty"`
}

// Validate validates this objects get response a o2 result
func (m *ObjectsGetResponseAO2Result) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectsGetResponseAO2Result) validateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	if m.Errors != nil {
		if err := m.Errors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result" + "." + "errors")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("result" + "." + "errors")
			}
			return err
		}
	}

	return nil
}

var objectsGetResponseAO2ResultTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUCCESS","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		objectsGetResponseAO2ResultTypeStatusPropEnum = append(objectsGetResponseAO2ResultTypeStatusPropEnum, v)
	}
}

const (

	// ObjectsGetResponseAO2ResultStatusSUCCESS captures enum value "SUCCESS"
	ObjectsGetResponseAO2ResultStatusSUCCESS string = "SUCCESS"

	// ObjectsGetResponseAO2ResultStatusFAILED captures enum value "FAILED"
	ObjectsGetResponseAO2ResultStatusFAILED string = "FAILED"
)

// prop value enum
func (m *ObjectsGetResponseAO2Result) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, objectsGetResponseAO2ResultTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ObjectsGetResponseAO2Result) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("result"+"."+"status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this objects get response a o2 result based on the context it is used
func (m *ObjectsGetResponseAO2Result) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectsGetResponseAO2Result) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	if m.Errors != nil {
		if err := m.Errors.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result" + "." + "errors")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("result" + "." + "errors")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectsGetResponseAO2Result) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectsGetResponseAO2Result) UnmarshalBinary(b []byte) error {
	var res ObjectsGetResponseAO2Result
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
