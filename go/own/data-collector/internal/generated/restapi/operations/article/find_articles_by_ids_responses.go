// Code generated by go-swagger; DO NOT EDIT.

package article

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"data-collector/internal/generated/models"
)

// FindArticlesByIdsOKCode is the HTTP code returned for type FindArticlesByIdsOK
const FindArticlesByIdsOKCode int = 200

/*
FindArticlesByIdsOK successful operation

swagger:response findArticlesByIdsOK
*/
type FindArticlesByIdsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Article `json:"body,omitempty"`
}

// NewFindArticlesByIdsOK creates FindArticlesByIdsOK with default headers values
func NewFindArticlesByIdsOK() *FindArticlesByIdsOK {

	return &FindArticlesByIdsOK{}
}

// WithPayload adds the payload to the find articles by ids o k response
func (o *FindArticlesByIdsOK) WithPayload(payload []*models.Article) *FindArticlesByIdsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find articles by ids o k response
func (o *FindArticlesByIdsOK) SetPayload(payload []*models.Article) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindArticlesByIdsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Article, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// FindArticlesByIdsBadRequestCode is the HTTP code returned for type FindArticlesByIdsBadRequest
const FindArticlesByIdsBadRequestCode int = 400

/*
FindArticlesByIdsBadRequest Invalid status value

swagger:response findArticlesByIdsBadRequest
*/
type FindArticlesByIdsBadRequest struct {
}

// NewFindArticlesByIdsBadRequest creates FindArticlesByIdsBadRequest with default headers values
func NewFindArticlesByIdsBadRequest() *FindArticlesByIdsBadRequest {

	return &FindArticlesByIdsBadRequest{}
}

// WriteResponse to the client
func (o *FindArticlesByIdsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
