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
	"github.com/go-openapi/swag"
)

// NewListAppsParams creates a new ListAppsParams object
// with the default values initialized.
func NewListAppsParams() *ListAppsParams {
	var ()
	return &ListAppsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAppsParamsWithTimeout creates a new ListAppsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAppsParamsWithTimeout(timeout time.Duration) *ListAppsParams {
	var ()
	return &ListAppsParams{

		timeout: timeout,
	}
}

// NewListAppsParamsWithContext creates a new ListAppsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAppsParamsWithContext(ctx context.Context) *ListAppsParams {
	var ()
	return &ListAppsParams{

		Context: ctx,
	}
}

// NewListAppsParamsWithHTTPClient creates a new ListAppsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAppsParamsWithHTTPClient(client *http.Client) *ListAppsParams {
	var ()
	return &ListAppsParams{
		HTTPClient: client,
	}
}

/*ListAppsParams contains all the parameters to send to the API endpoint
for the list apps operation typically these are written to a http.Request
*/
type ListAppsParams struct {

	/*Hidden
	  Whether or not to show hidden apps

	*/
	Hidden *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list apps params
func (o *ListAppsParams) WithTimeout(timeout time.Duration) *ListAppsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list apps params
func (o *ListAppsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list apps params
func (o *ListAppsParams) WithContext(ctx context.Context) *ListAppsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list apps params
func (o *ListAppsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list apps params
func (o *ListAppsParams) WithHTTPClient(client *http.Client) *ListAppsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list apps params
func (o *ListAppsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHidden adds the hidden to the list apps params
func (o *ListAppsParams) WithHidden(hidden *bool) *ListAppsParams {
	o.SetHidden(hidden)
	return o
}

// SetHidden adds the hidden to the list apps params
func (o *ListAppsParams) SetHidden(hidden *bool) {
	o.Hidden = hidden
}

// WriteToRequest writes these params to a swagger request
func (o *ListAppsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Hidden != nil {

		// query param hidden
		var qrHidden bool
		if o.Hidden != nil {
			qrHidden = *o.Hidden
		}
		qHidden := swag.FormatBool(qrHidden)
		if qHidden != "" {
			if err := r.SetQueryParam("hidden", qHidden); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
