// Code generated by go-swagger; DO NOT EDIT.

package repos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetReposParams creates a new GetReposParams object
// with the default values initialized.
func NewGetReposParams() *GetReposParams {
	var (
		pageNumberDefault = int64(1)
		pageSizeDefault   = int32(20)
	)
	return &GetReposParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReposParamsWithTimeout creates a new GetReposParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReposParamsWithTimeout(timeout time.Duration) *GetReposParams {
	var (
		pageNumberDefault = int64(1)
		pageSizeDefault   = int32(20)
	)
	return &GetReposParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetReposParamsWithContext creates a new GetReposParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReposParamsWithContext(ctx context.Context) *GetReposParams {
	var (
		pageNumberDefault = int64(1)
		pageSizeDefault   = int32(20)
	)
	return &GetReposParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetReposParamsWithHTTPClient creates a new GetReposParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReposParamsWithHTTPClient(client *http.Client) *GetReposParams {
	var (
		pageNumberDefault = int64(1)
		pageSizeDefault   = int32(20)
	)
	return &GetReposParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetReposParams contains all the parameters to send to the API endpoint
for the get repos operation typically these are written to a http.Request
*/
type GetReposParams struct {

	/*PageNumber
	  Page number

	*/
	PageNumber *int64
	/*PageSize
	  Number of clusters returned

	*/
	PageSize *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get repos params
func (o *GetReposParams) WithTimeout(timeout time.Duration) *GetReposParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repos params
func (o *GetReposParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repos params
func (o *GetReposParams) WithContext(ctx context.Context) *GetReposParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repos params
func (o *GetReposParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repos params
func (o *GetReposParams) WithHTTPClient(client *http.Client) *GetReposParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repos params
func (o *GetReposParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get repos params
func (o *GetReposParams) WithPageNumber(pageNumber *int64) *GetReposParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get repos params
func (o *GetReposParams) SetPageNumber(pageNumber *int64) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get repos params
func (o *GetReposParams) WithPageSize(pageSize *int32) *GetReposParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get repos params
func (o *GetReposParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetReposParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber int64
		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt64(qrPageNumber)
		if qPageNumber != "" {
			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
