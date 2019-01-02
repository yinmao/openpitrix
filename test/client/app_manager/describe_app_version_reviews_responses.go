// Code generated by go-swagger; DO NOT EDIT.

package app_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// DescribeAppVersionReviewsReader is a Reader for the DescribeAppVersionReviews structure.
type DescribeAppVersionReviewsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeAppVersionReviewsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDescribeAppVersionReviewsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeAppVersionReviewsOK creates a DescribeAppVersionReviewsOK with default headers values
func NewDescribeAppVersionReviewsOK() *DescribeAppVersionReviewsOK {
	return &DescribeAppVersionReviewsOK{}
}

/*DescribeAppVersionReviewsOK handles this case with default header values.

A successful response.
*/
type DescribeAppVersionReviewsOK struct {
	Payload *models.OpenpitrixDescribeAppVersionReviewsResponse
}

func (o *DescribeAppVersionReviewsOK) Error() string {
	return fmt.Sprintf("[GET /v1/app/{app_id}/version/{version_id}/reviews][%d] describeAppVersionReviewsOK  %+v", 200, o.Payload)
}

func (o *DescribeAppVersionReviewsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDescribeAppVersionReviewsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}