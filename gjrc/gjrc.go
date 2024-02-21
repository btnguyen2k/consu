package gjrc

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/btnguyen2k/consu/semita"
)

// NewGjrc creates a new Gjrc object.
// It reuses the http.Client if supplied (then timeout is ignored). Otherwise, a new client is created with the
// specified timeout.
func NewGjrc(httpClient *http.Client, timeout time.Duration) *Gjrc {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: timeout}
	}
	return &Gjrc{httpClient: httpClient}
}

// RequestMeta captures the metadata to be sent along with the request.
//
// Available since v0.2.0
type RequestMeta struct {
	Header  http.Header
	Timeout time.Duration
}

// Merge merges metadata from another instance into this one.
func (rm RequestMeta) Merge(other RequestMeta) RequestMeta {
	if other.Timeout > 0 {
		rm.Timeout = other.Timeout
	}
	if rm.Header == nil {
		rm.Header = http.Header{}
	}
	for k := range other.Header {
		rm.Header[k] = other.Header[k]
	}
	return rm
}

func mergeMetadata(starter RequestMeta, others ...RequestMeta) RequestMeta {
	for _, m := range others {
		starter = starter.Merge(m)
	}
	return starter
}

// Gjrc sends HTTP requests and wraps the HTTP response in a GjrcResponse.
type Gjrc struct {
	httpClient *http.Client
}

func (c *Gjrc) buildResponse(resp *http.Response, err error) *GjrcResponse {
	result := &GjrcResponse{err: err, resp: resp}
	go func() { _, _ = result.ensureResponseData() }()
	return result
}

// Do sends an HTTP request and returns a GjrcResponse capturing the HTTP response.
func (c *Gjrc) Do(req *http.Request) *GjrcResponse {
	return c.buildResponse(c.httpClient.Do(req))
}

/*----------------------------------------------------------------------*/

// buildContext creates is a convenience function to create a context and handle timeout.
//
// @Available since <<VERSION>>
func (c *Gjrc) buildContext(metadata RequestMeta) context.Context {
	ctx := context.Background()
	if metadata.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, metadata.Timeout)
		go func() {
			defer cancel()
			time.Sleep(metadata.Timeout)
		}()
	}
	return ctx
}

// Post sends a POST request and returns a GjrcResponse capturing the HTTP response.
//
// metadata is added since v0.2.0
func (c *Gjrc) Post(url, contentType string, body io.Reader, metadata ...RequestMeta) *GjrcResponse {
	h := http.Header{}
	h.Set("Content-Type", contentType)
	meta := mergeMetadata(RequestMeta{Header: h}, metadata...)
	ctx := c.buildContext(meta)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return c.buildResponse(nil, err)
	}

	for k := range meta.Header {
		req.Header.Set(k, meta.Header.Get(k))
	}

	return c.Do(req)
}

// PostForm sends a POST request with content type "application/x-www-form-urlencoded" and returns a GjrcResponse
// capturing the HTTP response.
//
// metadata is added since v0.2.0
func (c *Gjrc) PostForm(url string, data url.Values, metadata ...RequestMeta) *GjrcResponse {
	return c.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()), metadata...)
}

// Get sends a GET request and returns a GjrcResponse capturing the HTTP response.
//
// metadata is added since v0.2.0
func (c *Gjrc) Get(url string, metadata ...RequestMeta) *GjrcResponse {
	meta := mergeMetadata(RequestMeta{}, metadata...)
	ctx := c.buildContext(meta)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return c.buildResponse(nil, err)
	}

	for k := range meta.Header {
		req.Header.Set(k, meta.Header.Get(k))
	}

	return c.Do(req)
}

/*----------------------------------------------------------------------*/

func (c *Gjrc) buildJsonRequest(method, url string, bodyObj interface{}, metadata ...RequestMeta) (*http.Request, error) {
	h := http.Header{}
	h.Set("Content-Type", "application/json")
	meta := mergeMetadata(RequestMeta{Header: h}, metadata...)
	ctx := c.buildContext(meta)
	buf := make([]byte, 0)
	if bodyObj != nil {
		buf, _ = json.Marshal(bodyObj)
	}
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	for k := range meta.Header {
		req.Header.Set(k, meta.Header.Get(k))
	}
	return req, err
}

// DeleteJson sends a DELETE request with content type "application/json" and
// returns a GjrcResponse capturing the HTTP response.
//
// Note: bodyObj is marshalled to JSON string and sent along the request.
//
// metadata is added since v0.2.0
func (c *Gjrc) DeleteJson(url string, bodyObj interface{}, metadata ...RequestMeta) *GjrcResponse {
	req, err := c.buildJsonRequest(http.MethodDelete, url, bodyObj, metadata...)
	if err != nil {
		return c.buildResponse(nil, err)
	}
	return c.Do(req)
}

// PatchJson sends a PATCH request with content type "application/json" and
// returns a GjrcResponse capturing the HTTP response.
//
// Note: bodyObj is marshalled to JSON string and sent along the request.
//
// metadata is added since v0.2.0
func (c *Gjrc) PatchJson(url string, bodyObj interface{}, metadata ...RequestMeta) *GjrcResponse {
	req, err := c.buildJsonRequest(http.MethodPatch, url, bodyObj, metadata...)
	if err != nil {
		return c.buildResponse(nil, err)
	}
	return c.Do(req)
}

// PostJson sends a POST request with content type "application/json" and
// returns a GjrcResponse capturing the HTTP response.
//
// Note: bodyObj is marshalled to JSON string and sent along the request.
//
// metadata is added since v0.2.0
func (c *Gjrc) PostJson(url string, bodyObj interface{}, metadata ...RequestMeta) *GjrcResponse {
	req, err := c.buildJsonRequest(http.MethodPost, url, bodyObj, metadata...)
	if err != nil {
		return c.buildResponse(nil, err)
	}
	return c.Do(req)
}

// PutJson sends a PUT request with content type "application/json" and
// returns a GjrcResponse capturing the HTTP response.
//
// Note: bodyObj is marshalled to JSON string and sent along the request.
//
// metadata is added since v0.2.0
func (c *Gjrc) PutJson(url string, bodyObj interface{}, metadata ...RequestMeta) *GjrcResponse {
	req, err := c.buildJsonRequest(http.MethodPut, url, bodyObj, metadata...)
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
			defer func() { _ = r.resp.Body.Close() }()
			buff, err := ioutil.ReadAll(r.resp.Body) // leave it here as we are still supporting go v1.13
			if err != nil && r.err == nil {
				r.err = err
			}
			var obj interface{}
			err = json.Unmarshal(buff, &obj)
			if err != nil && r.err == nil {
				r.err = err
			}
			r.rawBody = buff
			r.objBody = obj
			r.s = semita.NewSemita(r.objBody)
		}
	}
	if r.s == nil {
		r.s = semita.NewSemita(r.objBody)
	}
	return r.rawBody, r.err
}

func (r *GjrcResponse) ensureResponseData() ([]byte, error) {
	if r.s == nil {
		return r.Body()
	}
	return r.rawBody, r.err
}

// GetValueAsType retrieves the value located at path and returns it casted to typ.
//
// Note: see semita.Semita documentation for path syntax.
func (r *GjrcResponse) GetValueAsType(path string, typ reflect.Type) (interface{}, error) {
	_, _ = r.ensureResponseData()
	return r.s.GetValueOfType(path, typ)
}

// Unmarshal parses the JSON-encoded response's body and puts the result to v.
//
// Note: v must be a pointer.
//
// Available since v0.2.1
func (r *GjrcResponse) Unmarshal(v interface{}) error {
	body, err := r.Body()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}
