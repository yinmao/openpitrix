// Code generated by go-swagger; DO NOT EDIT.

package market_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// NewCreateMarketParams creates a new CreateMarketParams object
// with the default values initialized.
func NewCreateMarketParams() *CreateMarketParams {
	var ()
	return &CreateMarketParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMarketParamsWithTimeout creates a new CreateMarketParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateMarketParamsWithTimeout(timeout time.Duration) *CreateMarketParams {
	var ()
	return &CreateMarketParams{

		timeout: timeout,
	}
}

// NewCreateMarketParamsWithContext creates a new CreateMarketParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateMarketParamsWithContext(ctx context.Context) *CreateMarketParams {
	var ()
	return &CreateMarketParams{

		Context: ctx,
	}
}

// NewCreateMarketParamsWithHTTPClient creates a new CreateMarketParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateMarketParamsWithHTTPClient(client *http.Client) *CreateMarketParams {
	var ()
	return &CreateMarketParams{
		HTTPClient: client,
	}
}

/*CreateMarketParams contains all the parameters to send to the API endpoint
for the create market operation typically these are written to a http.Request
*/
type CreateMarketParams struct {

	/*Body*/
	Body *models.OpenpitrixCreateMarketRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create market params
func (o *CreateMarketParams) WithTimeout(timeout time.Duration) *CreateMarketParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create market params
func (o *CreateMarketParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create market params
func (o *CreateMarketParams) WithContext(ctx context.Context) *CreateMarketParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create market params
func (o *CreateMarketParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create market params
func (o *CreateMarketParams) WithHTTPClient(client *http.Client) *CreateMarketParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create market params
func (o *CreateMarketParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create market params
func (o *CreateMarketParams) WithBody(body *models.OpenpitrixCreateMarketRequest) *CreateMarketParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create market params
func (o *CreateMarketParams) SetBody(body *models.OpenpitrixCreateMarketRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMarketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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