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

package batch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/weaviate/weaviate/entities/models"
)

// BatchObjectsCreateHandlerFunc turns a function with the right signature into a batch objects create handler
type BatchObjectsCreateHandlerFunc func(BatchObjectsCreateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn BatchObjectsCreateHandlerFunc) Handle(params BatchObjectsCreateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// BatchObjectsCreateHandler interface for that can handle valid batch objects create params
type BatchObjectsCreateHandler interface {
	Handle(BatchObjectsCreateParams, *models.Principal) middleware.Responder
}

// NewBatchObjectsCreate creates a new http.Handler for the batch objects create operation
func NewBatchObjectsCreate(ctx *middleware.Context, handler BatchObjectsCreateHandler) *BatchObjectsCreate {
	return &BatchObjectsCreate{Context: ctx, Handler: handler}
}

/*
	BatchObjectsCreate swagger:route POST /batch/objects batch objects batchObjectsCreate

Creates new Objects based on a Object template as a batch.

Register new Objects in bulk. Provided meta-data and schema values are validated.
*/
type BatchObjectsCreate struct {
	Context *middleware.Context
	Handler BatchObjectsCreateHandler
}

func (o *BatchObjectsCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewBatchObjectsCreateParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// BatchObjectsCreateBody batch objects create body
//
// swagger:model BatchObjectsCreateBody
type BatchObjectsCreateBody struct {

	// Define which fields need to be returned. Default value is ALL
	Fields []*string `json:"fields" yaml:"fields"`

	// objects
	Objects []*models.Object `json:"objects" yaml:"objects"`
}

// Validate validates this batch objects create body
func (o *BatchObjectsCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var batchObjectsCreateBodyFieldsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALL","class","schema","id","creationTimeUnix"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		batchObjectsCreateBodyFieldsItemsEnum = append(batchObjectsCreateBodyFieldsItemsEnum, v)
	}
}

func (o *BatchObjectsCreateBody) validateFieldsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, batchObjectsCreateBodyFieldsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *BatchObjectsCreateBody) validateFields(formats strfmt.Registry) error {
	if swag.IsZero(o.Fields) { // not required
		return nil
	}

	for i := 0; i < len(o.Fields); i++ {
		if swag.IsZero(o.Fields[i]) { // not required
			continue
		}

		// value enum
		if err := o.validateFieldsItemsEnum("body"+"."+"fields"+"."+strconv.Itoa(i), "body", *o.Fields[i]); err != nil {
			return err
		}

	}

	return nil
}

func (o *BatchObjectsCreateBody) validateObjects(formats strfmt.Registry) error {
	if swag.IsZero(o.Objects) { // not required
		return nil
	}

	for i := 0; i < len(o.Objects); i++ {
		if swag.IsZero(o.Objects[i]) { // not required
			continue
		}

		if o.Objects[i] != nil {
			if err := o.Objects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this batch objects create body based on the context it is used
func (o *BatchObjectsCreateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BatchObjectsCreateBody) contextValidateObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Objects); i++ {

		if o.Objects[i] != nil {
			if err := o.Objects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *BatchObjectsCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BatchObjectsCreateBody) UnmarshalBinary(b []byte) error {
	var res BatchObjectsCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
