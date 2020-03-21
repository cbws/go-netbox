// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package dcim

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

// NewDcimDevicesNapalmParams creates a new DcimDevicesNapalmParams object
// with the default values initialized.
func NewDcimDevicesNapalmParams() *DcimDevicesNapalmParams {
	var ()
	return &DcimDevicesNapalmParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDevicesNapalmParamsWithTimeout creates a new DcimDevicesNapalmParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimDevicesNapalmParamsWithTimeout(timeout time.Duration) *DcimDevicesNapalmParams {
	var ()
	return &DcimDevicesNapalmParams{

		timeout: timeout,
	}
}

// NewDcimDevicesNapalmParamsWithContext creates a new DcimDevicesNapalmParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimDevicesNapalmParamsWithContext(ctx context.Context) *DcimDevicesNapalmParams {
	var ()
	return &DcimDevicesNapalmParams{

		Context: ctx,
	}
}

// NewDcimDevicesNapalmParamsWithHTTPClient creates a new DcimDevicesNapalmParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimDevicesNapalmParamsWithHTTPClient(client *http.Client) *DcimDevicesNapalmParams {
	var ()
	return &DcimDevicesNapalmParams{
		HTTPClient: client,
	}
}

/*DcimDevicesNapalmParams contains all the parameters to send to the API endpoint
for the dcim devices napalm operation typically these are written to a http.Request
*/
type DcimDevicesNapalmParams struct {

	/*ID
	  A unique integer value identifying this device.

	*/
	ID int64
	/*Method*/
	Method string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim devices napalm params
func (o *DcimDevicesNapalmParams) WithTimeout(timeout time.Duration) *DcimDevicesNapalmParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim devices napalm params
func (o *DcimDevicesNapalmParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim devices napalm params
func (o *DcimDevicesNapalmParams) WithContext(ctx context.Context) *DcimDevicesNapalmParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim devices napalm params
func (o *DcimDevicesNapalmParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim devices napalm params
func (o *DcimDevicesNapalmParams) WithHTTPClient(client *http.Client) *DcimDevicesNapalmParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim devices napalm params
func (o *DcimDevicesNapalmParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim devices napalm params
func (o *DcimDevicesNapalmParams) WithID(id int64) *DcimDevicesNapalmParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim devices napalm params
func (o *DcimDevicesNapalmParams) SetID(id int64) {
	o.ID = id
}

// WithMethod adds the method to the dcim devices napalm params
func (o *DcimDevicesNapalmParams) WithMethod(method string) *DcimDevicesNapalmParams {
	o.SetMethod(method)
	return o
}

// SetMethod adds the method to the dcim devices napalm params
func (o *DcimDevicesNapalmParams) SetMethod(method string) {
	o.Method = method
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDevicesNapalmParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// query param method
	qrMethod := o.Method
	qMethod := qrMethod
	if qMethod != "" {
		if err := r.SetQueryParam("method", qMethod); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}