// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewObjectsListParams creates a new ObjectsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewObjectsListParams() *ObjectsListParams {
	return &ObjectsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewObjectsListParamsWithTimeout creates a new ObjectsListParams object
// with the ability to set a timeout on a request.
func NewObjectsListParamsWithTimeout(timeout time.Duration) *ObjectsListParams {
	return &ObjectsListParams{
		timeout: timeout,
	}
}

// NewObjectsListParamsWithContext creates a new ObjectsListParams object
// with the ability to set a context for a request.
func NewObjectsListParamsWithContext(ctx context.Context) *ObjectsListParams {
	return &ObjectsListParams{
		Context: ctx,
	}
}

// NewObjectsListParamsWithHTTPClient creates a new ObjectsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewObjectsListParamsWithHTTPClient(client *http.Client) *ObjectsListParams {
	return &ObjectsListParams{
		HTTPClient: client,
	}
}

/*
ObjectsListParams contains all the parameters to send to the API endpoint

	for the objects list operation.

	Typically these are written to a http.Request.
*/
type ObjectsListParams struct {

	/* After.

	   The starting ID of the result window.
	*/
	After *string

	/* Class.

	   Class parameter specifies the class from which to query objects
	*/
	Class *string

	/* Include.

	   Include additional information, such as classification infos. Allowed values include: classification, vector, interpretation
	*/
	Include *string

	/* Limit.

	   The maximum number of items to be returned per page. Default value is set in Weaviate config.

	   Format: int64
	*/
	Limit *int64

	/* Offset.

	   The starting index of the result window. Default value is 0.

	   Format: int64
	*/
	Offset *int64

	/* Order.

	   Order parameter to tell how to order (asc or desc) data within given field
	*/
	Order *string

	/* Sort.

	   Sort parameter to pass an information about the names of the sort fields
	*/
	Sort *string

	/* Tenant.

	   Specifies the tenant in a request targeting a multi-tenant class
	*/
	Tenant *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the objects list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ObjectsListParams) WithDefaults() *ObjectsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the objects list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ObjectsListParams) SetDefaults() {
	var (
		offsetDefault = int64(0)
	)

	val := ObjectsListParams{
		Offset: &offsetDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the objects list params
func (o *ObjectsListParams) WithTimeout(timeout time.Duration) *ObjectsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the objects list params
func (o *ObjectsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the objects list params
func (o *ObjectsListParams) WithContext(ctx context.Context) *ObjectsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the objects list params
func (o *ObjectsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the objects list params
func (o *ObjectsListParams) WithHTTPClient(client *http.Client) *ObjectsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the objects list params
func (o *ObjectsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAfter adds the after to the objects list params
func (o *ObjectsListParams) WithAfter(after *string) *ObjectsListParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the objects list params
func (o *ObjectsListParams) SetAfter(after *string) {
	o.After = after
}

// WithClass adds the class to the objects list params
func (o *ObjectsListParams) WithClass(class *string) *ObjectsListParams {
	o.SetClass(class)
	return o
}

// SetClass adds the class to the objects list params
func (o *ObjectsListParams) SetClass(class *string) {
	o.Class = class
}

// WithInclude adds the include to the objects list params
func (o *ObjectsListParams) WithInclude(include *string) *ObjectsListParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the objects list params
func (o *ObjectsListParams) SetInclude(include *string) {
	o.Include = include
}

// WithLimit adds the limit to the objects list params
func (o *ObjectsListParams) WithLimit(limit *int64) *ObjectsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the objects list params
func (o *ObjectsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the objects list params
func (o *ObjectsListParams) WithOffset(offset *int64) *ObjectsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the objects list params
func (o *ObjectsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrder adds the order to the objects list params
func (o *ObjectsListParams) WithOrder(order *string) *ObjectsListParams {
	o.SetOrder(order)
	return o
}

// SetOrder adds the order to the objects list params
func (o *ObjectsListParams) SetOrder(order *string) {
	o.Order = order
}

// WithSort adds the sort to the objects list params
func (o *ObjectsListParams) WithSort(sort *string) *ObjectsListParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the objects list params
func (o *ObjectsListParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithTenant adds the tenant to the objects list params
func (o *ObjectsListParams) WithTenant(tenant *string) *ObjectsListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the objects list params
func (o *ObjectsListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WriteToRequest writes these params to a swagger request
func (o *ObjectsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.After != nil {

		// query param after
		var qrAfter string

		if o.After != nil {
			qrAfter = *o.After
		}
		qAfter := qrAfter
		if qAfter != "" {

			if err := r.SetQueryParam("after", qAfter); err != nil {
				return err
			}
		}
	}

	if o.Class != nil {

		// query param class
		var qrClass string

		if o.Class != nil {
			qrClass = *o.Class
		}
		qClass := qrClass
		if qClass != "" {

			if err := r.SetQueryParam("class", qClass); err != nil {
				return err
			}
		}
	}

	if o.Include != nil {

		// query param include
		var qrInclude string

		if o.Include != nil {
			qrInclude = *o.Include
		}
		qInclude := qrInclude
		if qInclude != "" {

			if err := r.SetQueryParam("include", qInclude); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Order != nil {

		// query param order
		var qrOrder string

		if o.Order != nil {
			qrOrder = *o.Order
		}
		qOrder := qrOrder
		if qOrder != "" {

			if err := r.SetQueryParam("order", qOrder); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if o.Tenant != nil {

		// query param tenant
		var qrTenant string

		if o.Tenant != nil {
			qrTenant = *o.Tenant
		}
		qTenant := qrTenant
		if qTenant != "" {

			if err := r.SetQueryParam("tenant", qTenant); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
