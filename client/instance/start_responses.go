// Copyright (c) 2022 Yaohui Wang (yaohuiwang@outlook.com)
// FlexLB is licensed under Mulan PubL v2.
// You can use this software according to the terms and conditions of the Mulan PubL v2.
// You may obtain a copy of Mulan PubL v2 at:
//         http://license.coscl.org.cn/MulanPubL-2.0
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PubL v2 for more details.

// Code generated by go-swagger; DO NOT EDIT.

package instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"gitee.com/flexlb/flexlb-client-go/models"
)

// StartReader is a Reader for the Start structure.
type StartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStartBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStartUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStartForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStartNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStartInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStartOK creates a StartOK with default headers values
func NewStartOK() *StartOK {
	return &StartOK{}
}

/* StartOK describes a response with status code 200, with default header values.

Start instance succeeded
*/
type StartOK struct {
	Payload *models.Instance
}

func (o *StartOK) Error() string {
	return fmt.Sprintf("[POST /instances/{name}/start][%d] startOK  %+v", 200, o.Payload)
}
func (o *StartOK) GetPayload() *models.Instance {
	return o.Payload
}

func (o *StartOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Instance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartBadRequest creates a StartBadRequest with default headers values
func NewStartBadRequest() *StartBadRequest {
	return &StartBadRequest{}
}

/* StartBadRequest describes a response with status code 400, with default header values.

StartBadRequest start bad request
*/
type StartBadRequest struct {
	Payload *StartBadRequestBody
}

func (o *StartBadRequest) Error() string {
	return fmt.Sprintf("[POST /instances/{name}/start][%d] startBadRequest  %+v", 400, o.Payload)
}
func (o *StartBadRequest) GetPayload() *StartBadRequestBody {
	return o.Payload
}

func (o *StartBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartUnauthorized creates a StartUnauthorized with default headers values
func NewStartUnauthorized() *StartUnauthorized {
	return &StartUnauthorized{}
}

/* StartUnauthorized describes a response with status code 401, with default header values.

StartUnauthorized start unauthorized
*/
type StartUnauthorized struct {
	Payload interface{}
}

func (o *StartUnauthorized) Error() string {
	return fmt.Sprintf("[POST /instances/{name}/start][%d] startUnauthorized  %+v", 401, o.Payload)
}
func (o *StartUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *StartUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartForbidden creates a StartForbidden with default headers values
func NewStartForbidden() *StartForbidden {
	return &StartForbidden{}
}

/* StartForbidden describes a response with status code 403, with default header values.

StartForbidden start forbidden
*/
type StartForbidden struct {
	Payload *StartForbiddenBody
}

func (o *StartForbidden) Error() string {
	return fmt.Sprintf("[POST /instances/{name}/start][%d] startForbidden  %+v", 403, o.Payload)
}
func (o *StartForbidden) GetPayload() *StartForbiddenBody {
	return o.Payload
}

func (o *StartForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartNotFound creates a StartNotFound with default headers values
func NewStartNotFound() *StartNotFound {
	return &StartNotFound{}
}

/* StartNotFound describes a response with status code 404, with default header values.

StartNotFound start not found
*/
type StartNotFound struct {
	Payload *StartNotFoundBody
}

func (o *StartNotFound) Error() string {
	return fmt.Sprintf("[POST /instances/{name}/start][%d] startNotFound  %+v", 404, o.Payload)
}
func (o *StartNotFound) GetPayload() *StartNotFoundBody {
	return o.Payload
}

func (o *StartNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartInternalServerError creates a StartInternalServerError with default headers values
func NewStartInternalServerError() *StartInternalServerError {
	return &StartInternalServerError{}
}

/* StartInternalServerError describes a response with status code 500, with default header values.

StartInternalServerError start internal server error
*/
type StartInternalServerError struct {
	Payload interface{}
}

func (o *StartInternalServerError) Error() string {
	return fmt.Sprintf("[POST /instances/{name}/start][%d] startInternalServerError  %+v", 500, o.Payload)
}
func (o *StartInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *StartInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StartBadRequestBody start bad request body
swagger:model StartBadRequestBody
*/
type StartBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this start bad request body
func (o *StartBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("startBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this start bad request body based on context it is used
func (o *StartBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartBadRequestBody) UnmarshalBinary(b []byte) error {
	var res StartBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartForbiddenBody start forbidden body
swagger:model StartForbiddenBody
*/
type StartForbiddenBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this start forbidden body
func (o *StartForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartForbiddenBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("startForbidden"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this start forbidden body based on context it is used
func (o *StartForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartForbiddenBody) UnmarshalBinary(b []byte) error {
	var res StartForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartNotFoundBody start not found body
swagger:model StartNotFoundBody
*/
type StartNotFoundBody struct {

	// error
	// Required: true
	Error *string `json:"error"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this start not found body
func (o *StartNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartNotFoundBody) validateError(formats strfmt.Registry) error {

	if err := validate.Required("startNotFound"+"."+"error", "body", o.Error); err != nil {
		return err
	}

	return nil
}

func (o *StartNotFoundBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("startNotFound"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this start not found body based on context it is used
func (o *StartNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartNotFoundBody) UnmarshalBinary(b []byte) error {
	var res StartNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
