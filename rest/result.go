package rest

import (
	"io"
	"net/http"
)

type Result struct {
	Status int
	Header http.Header
	Body   interface{}
}

func BuildResult() *Result {
	return &Result{Status: 200, Header: make(http.Header)}
}

func (r *Result) WithStatus(status int) *Result {
	r.Status = status
	return r
}

func (r *Result) WithBody(body interface{}) *Result {
	r.Body = body
	if r.Body == nil && r.Status == 200 {
		r.Status = 204
	}
	return r
}

func (r *Result) AddHeader(key, value string) *Result {
	r.Header.Add(key, value)
	return r
}

func (r *Result) Send(resp http.ResponseWriter, encoder ResponseEncoder) error {
	for key, values := range r.Header {
		for _, value := range values {
			resp.Header().Add(key, value)
		}
	}
	resp.WriteHeader(r.Status)
	switch r.Body.(type) {
	case nil:
		return nil
	case io.Reader:
		_, err := io.Copy(resp, r.Body.(io.Reader))
		return err
	default:
		return encoder(resp, r.Body)
	}
}

type restHandler func(request *http.Request) (interface{}, error)

func (h restHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	encoder := StdResponseEncoderChooser(req)
	var err error
	result, err := h(req)
	if err == nil {
		switch result.(type) {
		case *Result:
			result.(*Result).Send(resp, encoder)
		default:
			BuildResult().WithBody(result).Send(resp, encoder)
		}
		if err == nil {
			return
		}
	}
	WrapError(err).Send(resp, encoder)
}

type createHandler func(*http.Request) (Resource, error)

func (h createHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	encoder := StdResponseEncoderChooser(req)
	var err error
	resource, err := h(req)
	if err == nil {
		result, err := resource.Get(req)
		if err == nil {
			switch result.(type) {
			case *Result:
				result.(*Result).
					AddHeader("location", resource.Location()).
					WithStatus(201).
					Send(resp, encoder)
			default:
				BuildResult().
					AddHeader("location", resource.Location()).
					WithStatus(201).
					WithBody(result).
					Send(resp, encoder)
			}
		}
	}
	WrapError(err).Send(resp, encoder)
}
