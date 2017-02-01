package operations

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

// NewForumGetUsersParams creates a new ForumGetUsersParams object
// with the default values initialized.
func NewForumGetUsersParams() *ForumGetUsersParams {
	var (
		limitDefault = int32(100)
	)
	return &ForumGetUsersParams{
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewForumGetUsersParamsWithTimeout creates a new ForumGetUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewForumGetUsersParamsWithTimeout(timeout time.Duration) *ForumGetUsersParams {
	var (
		limitDefault = int32(100)
	)
	return &ForumGetUsersParams{
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewForumGetUsersParamsWithContext creates a new ForumGetUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewForumGetUsersParamsWithContext(ctx context.Context) *ForumGetUsersParams {
	var (
		limitDefault = int32(100)
	)
	return &ForumGetUsersParams{
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewForumGetUsersParamsWithHTTPClient creates a new ForumGetUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewForumGetUsersParamsWithHTTPClient(client *http.Client) *ForumGetUsersParams {
	var (
		limitDefault = int32(100)
	)
	return &ForumGetUsersParams{
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*ForumGetUsersParams contains all the parameters to send to the API endpoint
for the forum get users operation typically these are written to a http.Request
*/
type ForumGetUsersParams struct {

	/*Desc
	  Флаг сортировки по убыванию.


	*/
	Desc *bool
	/*Limit
	  Максимальное кол-во возвращаемых записей.

	*/
	Limit *int32
	/*Since
	  Идентификатор пользователя, с которого будут выводиться пользоватли
	(пользователь с данным идентификатором в результат не попадает).


	*/
	Since *string
	/*Slug
	  Идентификатор форума.

	*/
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the forum get users params
func (o *ForumGetUsersParams) WithTimeout(timeout time.Duration) *ForumGetUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the forum get users params
func (o *ForumGetUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the forum get users params
func (o *ForumGetUsersParams) WithContext(ctx context.Context) *ForumGetUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the forum get users params
func (o *ForumGetUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the forum get users params
func (o *ForumGetUsersParams) WithHTTPClient(client *http.Client) *ForumGetUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the forum get users params
func (o *ForumGetUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDesc adds the desc to the forum get users params
func (o *ForumGetUsersParams) WithDesc(desc *bool) *ForumGetUsersParams {
	o.SetDesc(desc)
	return o
}

// SetDesc adds the desc to the forum get users params
func (o *ForumGetUsersParams) SetDesc(desc *bool) {
	o.Desc = desc
}

// WithLimit adds the limit to the forum get users params
func (o *ForumGetUsersParams) WithLimit(limit *int32) *ForumGetUsersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the forum get users params
func (o *ForumGetUsersParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithSince adds the since to the forum get users params
func (o *ForumGetUsersParams) WithSince(since *string) *ForumGetUsersParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the forum get users params
func (o *ForumGetUsersParams) SetSince(since *string) {
	o.Since = since
}

// WithSlug adds the slug to the forum get users params
func (o *ForumGetUsersParams) WithSlug(slug string) *ForumGetUsersParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the forum get users params
func (o *ForumGetUsersParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *ForumGetUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Desc != nil {

		// query param desc
		var qrDesc bool
		if o.Desc != nil {
			qrDesc = *o.Desc
		}
		qDesc := swag.FormatBool(qrDesc)
		if qDesc != "" {
			if err := r.SetQueryParam("desc", qDesc); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Since != nil {

		// query param since
		var qrSince string
		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := qrSince
		if qSince != "" {
			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}

	}

	// path param slug
	if err := r.SetPathParam("slug", o.Slug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}