// Code generated by go-swagger; DO NOT EDIT.

package article

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"data-collector/internal/generated/models"
)

// GetArticleByIDOKCode is the HTTP code returned for type GetArticleByIDOK
const GetArticleByIDOKCode int = 200

/*
GetArticleByIDOK successful operation

swagger:response getArticleByIdOK
*/
type GetArticleByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Article `json:"body,omitempty"`
}

// NewGetArticleByIDOK creates GetArticleByIDOK with default headers values
func NewGetArticleByIDOK() *GetArticleByIDOK {

	return &GetArticleByIDOK{}
}

// WithPayload adds the payload to the get article by Id o k response
func (o *GetArticleByIDOK) WithPayload(payload *models.Article) *GetArticleByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get article by Id o k response
func (o *GetArticleByIDOK) SetPayload(payload *models.Article) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetArticleByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetArticleByIDBadRequestCode is the HTTP code returned for type GetArticleByIDBadRequest
const GetArticleByIDBadRequestCode int = 400

/*
GetArticleByIDBadRequest Invalid ID supplied

swagger:response getArticleByIdBadRequest
*/
type GetArticleByIDBadRequest struct {
}

// NewGetArticleByIDBadRequest creates GetArticleByIDBadRequest with default headers values
func NewGetArticleByIDBadRequest() *GetArticleByIDBadRequest {

	return &GetArticleByIDBadRequest{}
}

// WriteResponse to the client
func (o *GetArticleByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetArticleByIDNotFoundCode is the HTTP code returned for type GetArticleByIDNotFound
const GetArticleByIDNotFoundCode int = 404

/*
GetArticleByIDNotFound Pet not found

swagger:response getArticleByIdNotFound
*/
type GetArticleByIDNotFound struct {
}

// NewGetArticleByIDNotFound creates GetArticleByIDNotFound with default headers values
func NewGetArticleByIDNotFound() *GetArticleByIDNotFound {

	return &GetArticleByIDNotFound{}
}

// WriteResponse to the client
func (o *GetArticleByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}