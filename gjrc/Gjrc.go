/*
Package gjrc offers generic utilities to work with JSON-based RESTful API.

Sample usage:

	package main

	import (
		"fmt"

		"github.com/btnguyen2k/consu/gjrc"
		"github.com/btnguyen2k/consu/reddo"
	)

	func main() {
		// // pre-build a http.Client
		// httpClient := &http.Client{}
		// client := NewGjrc(httpClient, 0)

		// or, a new http.Client is created with 10 seconds timeout
		client := NewGjrc(nil, 10*time.Second)

		url := "https://httpbin.org/post"
		resp := client.PostJson(url, map[string]interface{}{"key1": "value", "key2": 1, "key3": true})

		val1, err := resp.GetValueAsType("json.key1", reddo.TypeString)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n", val1) // output: "value"

		val2, err := resp.GetValueAsType("json.key2", reddo.TypeInt)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n", val2) // output: 2

		val3, err := resp.GetValueAsType("json.key3", reddo.TypeBool)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n", val3) // output: true
	}
*/
package gjrc

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"sync"
	"time"

	"github.com/btnguyen2k/consu/semita"
)

const (
	// Version defines version number of this package
	Version = "0.1.1"
)

// NewGjrc creates a new Gjrc object.
// It reuses the http.Client if supplied. Otherwise, a new client is created with specific timeout.
func NewGjrc(httpClient *http.Client, timeout time.Duration) *Gjrc {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: timeout}
	}
	return &Gjrc{httpClient: httpClient}
}

// Gjrc sends HTTP requests and wraps the HTTP response in a GjrcResponse.
type Gjrc struct {
	httpClient *http.Client
}

func (c *Gjrc) buildResponse(resp *http.Response, err error) *GjrcResponse {
	result := &GjrcResponse{err: err, resp: resp}
	go result.ensureResponseData()
	return result
}

// Do sends a HTTP request and returns a GjrcResponse capturing the HTTP response.
func (c *Gjrc) Do(req *http.Request) *GjrcResponse {
	return c.buildResponse(c.httpClient.Do(req))
}

// Post sends a POST request and returns a GjrcResponse capturing the HTTP response.
func (c *Gjrc) Post(url, contentType string, body io.Reader) *GjrcResponse {
	return c.buildResponse(c.httpClient.Post(url, contentType, body))
}

// PostForm sends a POST request with content type "application/x-www-form-urlencoded" and
// returns a GjrcResponse capturing the HTTP response.
func (c *Gjrc) PostForm(url string, data url.Values) *GjrcResponse {
	return c.buildResponse(c.httpClient.PostForm(url, data))
}

// Get sends a GET request and returns a GjrcResponse capturing the HTTP response.
func (c *Gjrc) Get(url string) *GjrcResponse {
	return c.buildResponse(c.httpClient.Get(url))
}

func (c *Gjrc) buildJsonRequest(method, url string, bodyObj interface{}) (*http.Request, error) {
	buf := make([]byte, 0)
	if bodyObj != nil {
		buf, _ = json.Marshal(bodyObj)
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(buf))
	if req != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, err
}

// DeleteJson sends a DELETE request with content type "application/json" and
// returns a GjrcResponse capturing the HTTP response.
//
// Note: bodyObj is marshalled to JSON string and sent along the request.
func (c *Gjrc) DeleteJson(url string, bodyObj interface{}) *GjrcResponse {
	req, err := c.buildJsonRequest("DELETE", url, bodyObj)
	if err != nil {
		return c.buildResponse(nil, err)
	}
	return c.Do(req)
}

// PatchJson sends a PATCH request with content type "application/json" and
// returns a GjrcResponse capturing the HTTP response.
//
// Note: bodyObj is marshalled to JSON string and sent along the request.
func (c *Gjrc) PatchJson(url string, bodyObj interface{}) *GjrcResponse {
	req, err := c.buildJsonRequest("PATCH", url, bodyObj)
	if err != nil {
		return c.buildResponse(nil, err)
	}
	return c.Do(req)
}

// PostJson sends a POST request with content type "application/json" and
// returns a GjrcResponse capturing the HTTP response.
//
// Note: bodyObj is marshalled to JSON string and sent along the request.
func (c *Gjrc) PostJson(url string, bodyObj interface{}) *GjrcResponse {
	req, err := c.buildJsonRequest("POST", url, bodyObj)
	if err != nil {
		return c.buildResponse(nil, err)
	}
	return c.Do(req)
}

// PutJson sends a PUT request with content type "application/json" and
// returns a GjrcResponse capturing the HTTP response.
//
// Note: bodyObj is marshalled to JSON string and sent along the request.
func (c *Gjrc) PutJson(url string, bodyObj interface{}) *GjrcResponse {
	req, err := c.buildJsonRequest("PUT", url, bodyObj)
	if err != nil {
		return c.buildResponse(nil, err)
	}
	return c.Do(req)
}

// GjrcResponse wraps around the HTTP response.
// Assuming the response body is JSON, GjrcResponse provides utility functions to access response data in a tree-like manner.
type GjrcResponse struct {
	err  error          // raw error making HTTP request
	resp *http.Response // raw HTTP response

	mutex   sync.Mutex
	rawBody []byte      // raw HTTP response body
	objBody interface{} // HTTP response body converted to object
	s       *semita.Semita
}

// Error returns the error (if any) caused by performing the request.
func (r *GjrcResponse) Error() error {
	return r.err
}

// HttpResponse returns the raw http.Response instance.
func (r *GjrcResponse) HttpResponse() *http.Response {
	return r.resp
}

// StatusCode returns the response status code.
func (r *GjrcResponse) StatusCode() int {
	return r.resp.StatusCode
}

// Body returns the raw response body.
func (r *GjrcResponse) Body() ([]byte, error) {
	if r.rawBody == nil && r.err == nil {
		r.mutex.Lock()
		defer r.mutex.Unlock()
		if r.rawBody == nil && r.err == nil {
			defer r.resp.Body.Close()
			buff, err := ioutil.ReadAll(r.resp.Body)
			var obj interface{}
			json.Unmarshal(buff, &obj)
			r.rawBody = buff
			r.objBody = obj
			r.s = semita.NewSemita(r.objBody)
			if err != nil && r.err == nil {
				r.err = err
			}
		}
	}
	if r.s == nil {
		r.s = semita.NewSemita(r.objBody)
	}
	return r.rawBody, r.err
}

func (r *GjrcResponse) ensureResponseData() {
	if r.s == nil {
		r.Body()
	}
}

// GetValueAsType retrieves the value located at path and returns it casted to typ.
//
// Note: see semita.Semita documentation for path syntax.
func (r *GjrcResponse) GetValueAsType(path string, typ reflect.Type) (interface{}, error) {
	r.ensureResponseData()
	return r.s.GetValueOfType(path, typ)
}
