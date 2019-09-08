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

	strfmt "github.com/go-openapi/strfmt"
)

// NewPredictParams creates a new PredictParams object
// with the default values initialized.
func NewPredictParams() *PredictParams {
	var ()
	return &PredictParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPredictParamsWithTimeout creates a new PredictParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPredictParamsWithTimeout(timeout time.Duration) *PredictParams {
	var ()
	return &PredictParams{

		timeout: timeout,
	}
}

// NewPredictParamsWithContext creates a new PredictParams object
// with the default values initialized, and the ability to set a context for a request
func NewPredictParamsWithContext(ctx context.Context) *PredictParams {
	var ()
	return &PredictParams{

		Context: ctx,
	}
}

// NewPredictParamsWithHTTPClient creates a new PredictParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPredictParamsWithHTTPClient(client *http.Client) *PredictParams {
	var ()
	return &PredictParams{
		HTTPClient: client,
	}
}

/*PredictParams contains all the parameters to send to the API endpoint
for the predict operation typically these are written to a http.Request
*/
type PredictParams struct {

	/*PositionList*/
	PositionList PredictBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the predict params
func (o *PredictParams) WithTimeout(timeout time.Duration) *PredictParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the predict params
func (o *PredictParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the predict params
func (o *PredictParams) WithContext(ctx context.Context) *PredictParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the predict params
func (o *PredictParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the predict params
func (o *PredictParams) WithHTTPClient(client *http.Client) *PredictParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the predict params
func (o *PredictParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPositionList adds the positionList to the predict params
func (o *PredictParams) WithPositionList(positionList PredictBody) *PredictParams {
	o.SetPositionList(positionList)
	return o
}

// SetPositionList adds the positionList to the predict params
func (o *PredictParams) SetPositionList(positionList PredictBody) {
	o.PositionList = positionList
}

// WriteToRequest writes these params to a swagger request
func (o *PredictParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.PositionList); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
