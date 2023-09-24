// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	"github.com/e2b-dev/api/packages/env-instance-task-driver/internal/client/models"
)

// NewPutCPUConfigurationParams creates a new PutCPUConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutCPUConfigurationParams() *PutCPUConfigurationParams {
	return &PutCPUConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutCPUConfigurationParamsWithTimeout creates a new PutCPUConfigurationParams object
// with the ability to set a timeout on a request.
func NewPutCPUConfigurationParamsWithTimeout(timeout time.Duration) *PutCPUConfigurationParams {
	return &PutCPUConfigurationParams{
		timeout: timeout,
	}
}

// NewPutCPUConfigurationParamsWithContext creates a new PutCPUConfigurationParams object
// with the ability to set a context for a request.
func NewPutCPUConfigurationParamsWithContext(ctx context.Context) *PutCPUConfigurationParams {
	return &PutCPUConfigurationParams{
		Context: ctx,
	}
}

// NewPutCPUConfigurationParamsWithHTTPClient creates a new PutCPUConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutCPUConfigurationParamsWithHTTPClient(client *http.Client) *PutCPUConfigurationParams {
	return &PutCPUConfigurationParams{
		HTTPClient: client,
	}
}

/*
PutCPUConfigurationParams contains all the parameters to send to the API endpoint

	for the put Cpu configuration operation.

	Typically these are written to a http.Request.
*/
type PutCPUConfigurationParams struct {

	/* Body.

	   CPU configuration request
	*/
	Body *models.CPUConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put Cpu configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCPUConfigurationParams) WithDefaults() *PutCPUConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put Cpu configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCPUConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put Cpu configuration params
func (o *PutCPUConfigurationParams) WithTimeout(timeout time.Duration) *PutCPUConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put Cpu configuration params
func (o *PutCPUConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put Cpu configuration params
func (o *PutCPUConfigurationParams) WithContext(ctx context.Context) *PutCPUConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put Cpu configuration params
func (o *PutCPUConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put Cpu configuration params
func (o *PutCPUConfigurationParams) WithHTTPClient(client *http.Client) *PutCPUConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put Cpu configuration params
func (o *PutCPUConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put Cpu configuration params
func (o *PutCPUConfigurationParams) WithBody(body *models.CPUConfig) *PutCPUConfigurationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put Cpu configuration params
func (o *PutCPUConfigurationParams) SetBody(body *models.CPUConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutCPUConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
