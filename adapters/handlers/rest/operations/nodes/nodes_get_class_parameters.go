//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewNodesGetClassParams creates a new NodesGetClassParams object
// with the default values initialized.
func NewNodesGetClassParams() NodesGetClassParams {

	var (
		// initialize parameters with default values

		outputDefault = string("minimal")
	)

	return NodesGetClassParams{
		Output: &outputDefault,
	}
}

// NodesGetClassParams contains all the bound params for the nodes get class operation
// typically these are obtained from a http.Request
//
// swagger:parameters nodes.get.class
type NodesGetClassParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	ClassName string
	/*Controls the verbosity of the output, possible values are: "minimal", "verbose". Defaults to "minimal".
	  In: query
	  Default: "minimal"
	*/
	Output *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewNodesGetClassParams() beforehand.
func (o *NodesGetClassParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rClassName, rhkClassName, _ := route.Params.GetOK("className")
	if err := o.bindClassName(rClassName, rhkClassName, route.Formats); err != nil {
		res = append(res, err)
	}

	qOutput, qhkOutput, _ := qs.GetOK("output")
	if err := o.bindOutput(qOutput, qhkOutput, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindClassName binds and validates parameter ClassName from path.
func (o *NodesGetClassParams) bindClassName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ClassName = raw

	return nil
}

// bindOutput binds and validates parameter Output from query.
func (o *NodesGetClassParams) bindOutput(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewNodesGetClassParams()
		return nil
	}
	o.Output = &raw

	return nil
}