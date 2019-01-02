// Code generated by go-swagger; DO NOT EDIT.

package app_manager

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

// NewDescribeActiveAppVersionsParams creates a new DescribeActiveAppVersionsParams object
// with the default values initialized.
func NewDescribeActiveAppVersionsParams() *DescribeActiveAppVersionsParams {
	var ()
	return &DescribeActiveAppVersionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeActiveAppVersionsParamsWithTimeout creates a new DescribeActiveAppVersionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeActiveAppVersionsParamsWithTimeout(timeout time.Duration) *DescribeActiveAppVersionsParams {
	var ()
	return &DescribeActiveAppVersionsParams{

		timeout: timeout,
	}
}

// NewDescribeActiveAppVersionsParamsWithContext creates a new DescribeActiveAppVersionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeActiveAppVersionsParamsWithContext(ctx context.Context) *DescribeActiveAppVersionsParams {
	var ()
	return &DescribeActiveAppVersionsParams{

		Context: ctx,
	}
}

// NewDescribeActiveAppVersionsParamsWithHTTPClient creates a new DescribeActiveAppVersionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeActiveAppVersionsParamsWithHTTPClient(client *http.Client) *DescribeActiveAppVersionsParams {
	var ()
	return &DescribeActiveAppVersionsParams{
		HTTPClient: client,
	}
}

/*DescribeActiveAppVersionsParams contains all the parameters to send to the API endpoint
for the describe active app versions operation typically these are written to a http.Request
*/
type DescribeActiveAppVersionsParams struct {

	/*AppID*/
	AppID []string
	/*Description*/
	Description []string
	/*DisplayColumns*/
	DisplayColumns []string
	/*Limit*/
	Limit *int64
	/*Name*/
	Name []string
	/*Offset*/
	Offset *int64
	/*Owner*/
	Owner []string
	/*PackageName*/
	PackageName []string
	/*Reverse*/
	Reverse *bool
	/*SearchWord*/
	SearchWord *string
	/*SortKey*/
	SortKey *string
	/*Status*/
	Status []string
	/*Type*/
	Type []string
	/*VersionID*/
	VersionID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) WithTimeout(timeout time.Duration) *DescribeActiveAppVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) WithContext(ctx context.Context) *DescribeActiveAppVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) WithHTTPClient(client *http.Client) *DescribeActiveAppVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) WithAppID(appID []string) *DescribeActiveAppVersionsParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) SetAppID(appID []string) {
	o.AppID = appID
}

// WithDescription adds the description to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) WithDescription(description []string) *DescribeActiveAppVersionsParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) SetDescription(description []string) {
	o.Description = description
}

// WithDisplayColumns adds the displayColumns to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) WithDisplayColumns(displayColumns []string) *DescribeActiveAppVersionsParams {
	o.SetDisplayColumns(displayColumns)
	return o
}

// SetDisplayColumns adds the displayColumns to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) SetDisplayColumns(displayColumns []string) {
	o.DisplayColumns = displayColumns
}

// WithLimit adds the limit to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) WithLimit(limit *int64) *DescribeActiveAppVersionsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) WithName(name []string) *DescribeActiveAppVersionsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) SetName(name []string) {
	o.Name = name
}

// WithOffset adds the offset to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) WithOffset(offset *int64) *DescribeActiveAppVersionsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOwner adds the owner to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) WithOwner(owner []string) *DescribeActiveAppVersionsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) SetOwner(owner []string) {
	o.Owner = owner
}

// WithPackageName adds the packageName to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) WithPackageName(packageName []string) *DescribeActiveAppVersionsParams {
	o.SetPackageName(packageName)
	return o
}

// SetPackageName adds the packageName to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) SetPackageName(packageName []string) {
	o.PackageName = packageName
}

// WithReverse adds the reverse to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) WithReverse(reverse *bool) *DescribeActiveAppVersionsParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithSearchWord adds the searchWord to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) WithSearchWord(searchWord *string) *DescribeActiveAppVersionsParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) WithSortKey(sortKey *string) *DescribeActiveAppVersionsParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WithStatus adds the status to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) WithStatus(status []string) *DescribeActiveAppVersionsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) SetStatus(status []string) {
	o.Status = status
}

// WithType adds the typeVar to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) WithType(typeVar []string) *DescribeActiveAppVersionsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) SetType(typeVar []string) {
	o.Type = typeVar
}

// WithVersionID adds the versionID to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) WithVersionID(versionID []string) *DescribeActiveAppVersionsParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the describe active app versions params
func (o *DescribeActiveAppVersionsParams) SetVersionID(versionID []string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeActiveAppVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesAppID := o.AppID

	joinedAppID := swag.JoinByFormat(valuesAppID, "multi")
	// query array param app_id
	if err := r.SetQueryParam("app_id", joinedAppID...); err != nil {
		return err
	}

	valuesDescription := o.Description

	joinedDescription := swag.JoinByFormat(valuesDescription, "multi")
	// query array param description
	if err := r.SetQueryParam("description", joinedDescription...); err != nil {
		return err
	}

	valuesDisplayColumns := o.DisplayColumns

	joinedDisplayColumns := swag.JoinByFormat(valuesDisplayColumns, "multi")
	// query array param display_columns
	if err := r.SetQueryParam("display_columns", joinedDisplayColumns...); err != nil {
		return err
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

	valuesName := o.Name

	joinedName := swag.JoinByFormat(valuesName, "multi")
	// query array param name
	if err := r.SetQueryParam("name", joinedName...); err != nil {
		return err
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

	valuesOwner := o.Owner

	joinedOwner := swag.JoinByFormat(valuesOwner, "multi")
	// query array param owner
	if err := r.SetQueryParam("owner", joinedOwner...); err != nil {
		return err
	}

	valuesPackageName := o.PackageName

	joinedPackageName := swag.JoinByFormat(valuesPackageName, "multi")
	// query array param package_name
	if err := r.SetQueryParam("package_name", joinedPackageName...); err != nil {
		return err
	}

	if o.Reverse != nil {

		// query param reverse
		var qrReverse bool
		if o.Reverse != nil {
			qrReverse = *o.Reverse
		}
		qReverse := swag.FormatBool(qrReverse)
		if qReverse != "" {
			if err := r.SetQueryParam("reverse", qReverse); err != nil {
				return err
			}
		}

	}

	if o.SearchWord != nil {

		// query param search_word
		var qrSearchWord string
		if o.SearchWord != nil {
			qrSearchWord = *o.SearchWord
		}
		qSearchWord := qrSearchWord
		if qSearchWord != "" {
			if err := r.SetQueryParam("search_word", qSearchWord); err != nil {
				return err
			}
		}

	}

	if o.SortKey != nil {

		// query param sort_key
		var qrSortKey string
		if o.SortKey != nil {
			qrSortKey = *o.SortKey
		}
		qSortKey := qrSortKey
		if qSortKey != "" {
			if err := r.SetQueryParam("sort_key", qSortKey); err != nil {
				return err
			}
		}

	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "multi")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	valuesType := o.Type

	joinedType := swag.JoinByFormat(valuesType, "multi")
	// query array param type
	if err := r.SetQueryParam("type", joinedType...); err != nil {
		return err
	}

	valuesVersionID := o.VersionID

	joinedVersionID := swag.JoinByFormat(valuesVersionID, "multi")
	// query array param version_id
	if err := r.SetQueryParam("version_id", joinedVersionID...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}