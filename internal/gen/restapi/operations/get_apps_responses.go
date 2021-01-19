// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/aklinker1/miasma/internal/gen/models"
)

// GetAppsOKCode is the HTTP code returned for type GetAppsOK
const GetAppsOKCode int = 200

/*GetAppsOK OK

swagger:response getAppsOK
*/
type GetAppsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.App `json:"body,omitempty"`
}

// NewGetAppsOK creates GetAppsOK with default headers values
func NewGetAppsOK() *GetAppsOK {

	return &GetAppsOK{}
}

// WithPayload adds the payload to the get apps o k response
func (o *GetAppsOK) WithPayload(payload []*models.App) *GetAppsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get apps o k response
func (o *GetAppsOK) SetPayload(payload []*models.App) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.App, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetAppsDefault Unknown Error

swagger:response getAppsDefault
*/
type GetAppsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetAppsDefault creates GetAppsDefault with default headers values
func NewGetAppsDefault(code int) *GetAppsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAppsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get apps default response
func (o *GetAppsDefault) WithStatusCode(code int) *GetAppsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get apps default response
func (o *GetAppsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get apps default response
func (o *GetAppsDefault) WithPayload(payload string) *GetAppsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get apps default response
func (o *GetAppsDefault) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
