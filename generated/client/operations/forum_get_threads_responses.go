// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bozaro/tech-db-forum/generated/models"
)

// ForumGetThreadsReader is a Reader for the ForumGetThreads structure.
type ForumGetThreadsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ForumGetThreadsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewForumGetThreadsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewForumGetThreadsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewForumGetThreadsOK creates a ForumGetThreadsOK with default headers values
func NewForumGetThreadsOK() *ForumGetThreadsOK {
	return &ForumGetThreadsOK{}
}

/*ForumGetThreadsOK handles this case with default header values.

Информация о ветках обсуждения на форуме.

*/
type ForumGetThreadsOK struct {
	Payload models.Threads
}

func (o *ForumGetThreadsOK) Error() string {
	return fmt.Sprintf("[GET /forum/{slug}/threads][%d] forumGetThreadsOK  %+v", 200, o.Payload)
}

func (o *ForumGetThreadsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewForumGetThreadsNotFound creates a ForumGetThreadsNotFound with default headers values
func NewForumGetThreadsNotFound() *ForumGetThreadsNotFound {
	return &ForumGetThreadsNotFound{}
}

/*ForumGetThreadsNotFound handles this case with default header values.

Форум отсутсвует в системе.

*/
type ForumGetThreadsNotFound struct {
	Payload *models.Error
}

func (o *ForumGetThreadsNotFound) Error() string {
	return fmt.Sprintf("[GET /forum/{slug}/threads][%d] forumGetThreadsNotFound  %+v", 404, o.Payload)
}

func (o *ForumGetThreadsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
