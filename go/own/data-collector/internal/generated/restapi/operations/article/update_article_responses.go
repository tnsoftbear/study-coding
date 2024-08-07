// Code generated by go-swagger; DO NOT EDIT.

package article

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UpdateArticleBadRequestCode is the HTTP code returned for type UpdateArticleBadRequest
const UpdateArticleBadRequestCode int = 400

/*
UpdateArticleBadRequest Invalid ID supplied

swagger:response updateArticleBadRequest
*/
type UpdateArticleBadRequest struct {
}

// NewUpdateArticleBadRequest creates UpdateArticleBadRequest with default headers values
func NewUpdateArticleBadRequest() *UpdateArticleBadRequest {

	return &UpdateArticleBadRequest{}
}

// WriteResponse to the client
func (o *UpdateArticleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// UpdateArticleNotFoundCode is the HTTP code returned for type UpdateArticleNotFound
const UpdateArticleNotFoundCode int = 404

/*
UpdateArticleNotFound Article not found

swagger:response updateArticleNotFound
*/
type UpdateArticleNotFound struct {
}

// NewUpdateArticleNotFound creates UpdateArticleNotFound with default headers values
func NewUpdateArticleNotFound() *UpdateArticleNotFound {

	return &UpdateArticleNotFound{}
}

// WriteResponse to the client
func (o *UpdateArticleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// UpdateArticleMethodNotAllowedCode is the HTTP code returned for type UpdateArticleMethodNotAllowed
const UpdateArticleMethodNotAllowedCode int = 405

/*
UpdateArticleMethodNotAllowed Validation exception

swagger:response updateArticleMethodNotAllowed
*/
type UpdateArticleMethodNotAllowed struct {
}

// NewUpdateArticleMethodNotAllowed creates UpdateArticleMethodNotAllowed with default headers values
func NewUpdateArticleMethodNotAllowed() *UpdateArticleMethodNotAllowed {

	return &UpdateArticleMethodNotAllowed{}
}

// WriteResponse to the client
func (o *UpdateArticleMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}
