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

// ThreadCreateReader is a Reader for the ThreadCreate structure.
type ThreadCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ThreadCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewThreadCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewThreadCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewThreadCreateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewThreadCreateCreated creates a ThreadCreateCreated with default headers values
func NewThreadCreateCreated() *ThreadCreateCreated {
	return &ThreadCreateCreated{}
}

/*ThreadCreateCreated handles this case with default header values.

Ветка обсуждения успешно создана.
Возвращает данные созданной ветки обсуждения.

*/
type ThreadCreateCreated struct {
	Payload *models.Thread
}

func (o *ThreadCreateCreated) Error() string {
	return fmt.Sprintf("[POST /forum/{slug}/create][%d] threadCreateCreated  %+v", 201, o.Payload)
}

func (o *ThreadCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Thread)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThreadCreateNotFound creates a ThreadCreateNotFound with default headers values
func NewThreadCreateNotFound() *ThreadCreateNotFound {
	return &ThreadCreateNotFound{}
}

/*ThreadCreateNotFound handles this case with default header values.

Автор ветки или форум не найдены.

*/
type ThreadCreateNotFound struct {
	Payload *models.Error
}

func (o *ThreadCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /forum/{slug}/create][%d] threadCreateNotFound  %+v", 404, o.Payload)
}

func (o *ThreadCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewThreadCreateConflict creates a ThreadCreateConflict with default headers values
func NewThreadCreateConflict() *ThreadCreateConflict {
	return &ThreadCreateConflict{}
}

/*ThreadCreateConflict handles this case with default header values.

Ветка обсуждения уже присутсвует в базе данных.
Возвращает данные ранее созданной ветки обсуждения.

*/
type ThreadCreateConflict struct {
	Payload *models.Thread
}

func (o *ThreadCreateConflict) Error() string {
	return fmt.Sprintf("[POST /forum/{slug}/create][%d] threadCreateConflict  %+v", 409, o.Payload)
}

func (o *ThreadCreateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Thread)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
