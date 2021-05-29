// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/aklinker1/miasma/package/models"
)

// ListPluginsOKCode is the HTTP code returned for type ListPluginsOK
const ListPluginsOKCode int = 200

/*ListPluginsOK OK

swagger:response listPluginsOK
*/
type ListPluginsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Plugin `json:"body,omitempty"`
}

// NewListPluginsOK creates ListPluginsOK with default headers values
func NewListPluginsOK() *ListPluginsOK {

	return &ListPluginsOK{}
}

// WithPayload adds the payload to the list plugins o k response
func (o *ListPluginsOK) WithPayload(payload []*models.Plugin) *ListPluginsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list plugins o k response
func (o *ListPluginsOK) SetPayload(payload []*models.Plugin) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPluginsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Plugin, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*ListPluginsDefault Unknown Error

swagger:response listPluginsDefault
*/
type ListPluginsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewListPluginsDefault creates ListPluginsDefault with default headers values
func NewListPluginsDefault(code int) *ListPluginsDefault {
	if code <= 0 {
		code = 500
	}

	return &ListPluginsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list plugins default response
func (o *ListPluginsDefault) WithStatusCode(code int) *ListPluginsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list plugins default response
func (o *ListPluginsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list plugins default response
func (o *ListPluginsDefault) WithPayload(payload string) *ListPluginsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list plugins default response
func (o *ListPluginsDefault) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPluginsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
