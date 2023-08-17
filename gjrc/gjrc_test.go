package gjrc

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/btnguyen2k/consu/reddo"
)

func TestNewGjrc(t *testing.T) {
	testName := "TestNewGjrc"
	client := NewGjrc(nil, 10*time.Second)
	if client == nil {
		t.Fatalf("%s failed: nil", testName)
	}
	if client.httpClient == nil {
		t.Fatalf("%s failed: httpClient is nil", testName)
	}
}

func TestNewGjrc_PrebuiltHttpClient(t *testing.T) {
	testName := "TestNewGjrc_PrebuiltHttpClient"
	httpClient := &http.Client{}
	client := NewGjrc(httpClient, 0)
	if client == nil {
		t.Fatalf("%s failed: nil", testName)
	}
	if client.httpClient != httpClient {
		t.Fatalf("%s failed: pre-built http.Client is not used", testName)
	}
}

func TestGjrcResponse_Error(t *testing.T) {
	testName := "TestGjrcResponse_Error"
	client := NewGjrc(nil, 10*time.Second)
	url := "https://mydomain.notreal"
	resp := client.Get(url)
	if resp == nil {
		t.Fatalf("%s failed: nil response", testName)
	}
	if resp.Error() == nil {
		t.Fatalf("%s failed: nil error", testName)
	}
}

func TestGjrcResponse_StatusCode(t *testing.T) {
	s := newJsonHttpServer(0)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	defer func() {
		if err := s.Shutdown(); err != nil {
			panic(err)
		}
	}()

	testName := "TestGjrcResponse_StatusCode"
	client := NewGjrc(nil, 10*time.Second)
	url := fmt.Sprintf("http://localhost:%d/notfound", s.listenPort)
	resp := client.Get(url)
	if resp == nil {
		t.Fatalf("%s failed: nil response", testName)
	}
	if resp.Error() != nil {
		t.Fatalf("%s failed: %e", testName, resp.Error())
	}
	if resp.StatusCode() != 404 {
		t.Fatalf("%s failed: expected status code %#v but received %#v", testName, 404, resp.StatusCode())
	}
}

func TestGjrcResponse_Body(t *testing.T) {
	s := newJsonHttpServer(0)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	defer func() {
		if err := s.Shutdown(); err != nil {
			panic(err)
		}
	}()

	testName := "TestGjrcResponse_Body"
	client := NewGjrc(nil, 10*time.Second)
	url := fmt.Sprintf("http://localhost:%d/get", s.listenPort)
	resp := client.Get(url)
	if resp == nil {
		t.Fatalf("%s failed: nil response", testName)
	}
	if resp.Error() != nil {
		t.Fatalf("%s failed: %e", testName, resp.Error())
	}
	if resp.StatusCode() != 200 {
		t.Fatalf("%s failed: expected status code %#v but received %#v", testName, 200, resp.StatusCode())
	}
	body, err := resp.Body()
	if err != nil {
		t.Fatalf("%s failed: %e", testName, err)
	}
	if !strings.HasPrefix(string(body), "{") {
		t.Fatalf("%s failed: unexpected response body %s", testName, body)
	}
}

func TestGjrcResponse_HttpResponse(t *testing.T) {
	s := newJsonHttpServer(0)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	defer func() {
		if err := s.Shutdown(); err != nil {
			panic(err)
		}
	}()

	testName := "TestGjrcResponse_HttpResponse"
	client := NewGjrc(nil, 10*time.Second)
	url := fmt.Sprintf("http://localhost:%d/get", s.listenPort)
	resp := client.Get(url)
	if resp == nil {
		t.Fatalf("%s failed: nil response", testName)
	}
	if resp.Error() != nil {
		t.Fatalf("%s failed: %e", testName, resp.Error())
	}
	if resp.StatusCode() != 200 {
		t.Fatalf("%s failed: expected status code %#v but received %#v", testName, 200, resp.StatusCode())
	}
	if resp.HttpResponse() == nil {
		t.Fatalf("%s failed: nil http.Response", testName)
	}
}

/*----------------------------------------------------------------------*/

var (
	requestDataJson = map[string]interface{}{"key1": "value1", "key2": 2, "key3": true}
	requestHeader1  = RequestMeta{Header: map[string][]string{"X-Header1": {"value1"}}}
	requestHeader2  = RequestMeta{Header: map[string][]string{"X-Header2": {"2"}}}
	requestHeader3  = RequestMeta{Header: map[string][]string{"X-Header3": {"true"}}}
	requestHeaders  = []RequestMeta{requestHeader1, requestHeader2, requestHeader3}

	verifyKeysAll    = []string{"json.key1", "json.key2", "json.key3", "headers.X-Header1", "headers.X-Header2", "headers.X-Header3"}
	verifyKeysHeader = []string{"headers.X-Header1", "headers.X-Header2", "headers.X-Header3"}
	verifyTypes      = map[string]reflect.Type{
		"json.key1":         reddo.TypeString,
		"json.key2":         reddo.TypeInt,
		"json.key3":         reddo.TypeBool,
		"headers.X-Header1": reddo.TypeString,
		"headers.X-Header2": reddo.TypeString,
		"headers.X-Header3": reddo.TypeString,
	}
	verifyValues = map[string]interface{}{
		"json.key1":         "value1",
		"json.key2":         int64(2),
		"json.key3":         true,
		"headers.X-Header1": "value1",
		"headers.X-Header2": "2",
		"headers.X-Header3": "true",
	}
)

func TestGjrc_DeleteJson(t *testing.T) {
	s := newJsonHttpServer(0)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	defer func() {
		if err := s.Shutdown(); err != nil {
			panic(err)
		}
	}()

	testName := "TestGjrc_DeleteJson"
	client := NewGjrc(nil, 10*time.Second)
	url := fmt.Sprintf("http://localhost:%d/delete", s.listenPort)
	resp := client.DeleteJson(url, requestDataJson, requestHeaders...)
	if resp == nil {
		t.Fatalf("%s failed: nil response", testName)
	}

	valUrl, err := resp.GetValueAsType("url", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", testName, err)
	}
	if valUrl != url {
		t.Fatalf("%s failed: expected %s but received %s", testName, url, valUrl)
	}

	valHost, err := resp.GetValueAsType("headers.Host", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", testName, err)
	}
	host := "localhost"
	if valHost != host {
		t.Fatalf("%s failed: expected %s but received %s", testName, host, valHost)
	}

	for _, key := range verifyKeysAll {
		val, err := resp.GetValueAsType(key, verifyTypes[key])
		if err != nil {
			t.Fatalf("%s failed: %e", testName, err)
		}
		if val != verifyValues[key] {
			t.Fatalf("%s failed: <%#v> expected %#v but received %#v", testName, key, verifyValues[key], val)
		}
	}
}

func TestGjrc_DeleteJsonError(t *testing.T) {
	testName := "TestGjrc_DeleteJsonError"
	client := NewGjrc(nil, 10*time.Second)
	url := "://localhost/delete"
	resp := client.DeleteJson(url, map[string]interface{}{"key1": "value1", "key2": 2, "key3": false})
	if resp == nil {
		t.Fatalf("%s failed: nil response", testName)
	}
	if resp.Error() == nil {
		t.Fatalf("%s failed: expected error", testName)
	}
	if resp.HttpResponse() != nil {
		t.Fatalf("%s failed: expected nil http response", testName)
	}
}

func TestGjrc_Get(t *testing.T) {
	s := newJsonHttpServer(0)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	defer func() {
		if err := s.Shutdown(); err != nil {
			panic(err)
		}
	}()

	testName := "TestGjrc_Get"
	client := NewGjrc(nil, 10*time.Second)
	url := fmt.Sprintf("http://localhost:%d/get", s.listenPort)
	resp := client.Get(url, requestHeaders...)
	if resp == nil {
		t.Fatalf("%s failed: nil response", testName)
	}

	valUrl, err := resp.GetValueAsType("url", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", testName, err)
	}
	if valUrl != url {
		t.Fatalf("%s failed: expected %s but received %s", testName, url, valUrl)
	}

	valHost, err := resp.GetValueAsType("headers.Host", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", testName, err)
	}
	host := "localhost"
	if valHost != host {
		t.Fatalf("%s failed: expected %s but received %s", testName, host, valHost)
	}

	for _, key := range verifyKeysHeader {
		val, err := resp.GetValueAsType(key, verifyTypes[key])
		if err != nil {
			t.Fatalf("%s failed: %e", testName, err)
		}
		if val != verifyValues[key] {
			t.Fatalf("%s failed: <%#v> expected %#v but received %#v", testName, key, verifyValues[key], val)
		}
	}
}

func TestGjrc_PatchJson(t *testing.T) {
	s := newJsonHttpServer(0)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	defer func() {
		if err := s.Shutdown(); err != nil {
			panic(err)
		}
	}()

	testName := "TestGjrc_PatchJson"
	client := NewGjrc(nil, 10*time.Second)
	url := fmt.Sprintf("http://localhost:%d/patch", s.listenPort)
	resp := client.PatchJson(url, requestDataJson, requestHeaders...)
	if resp == nil {
		t.Fatalf("%s failed: nil response", testName)
	}

	valUrl, err := resp.GetValueAsType("url", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", testName, err)
	}
	if valUrl != url {
		t.Fatalf("%s failed: expected %s but received %s", testName, url, valUrl)
	}

	valHost, err := resp.GetValueAsType("headers.Host", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", testName, err)
	}
	host := "localhost"
	if valHost != host {
		t.Fatalf("%s failed: expected %s but received %s", testName, host, valHost)
	}

	for _, key := range verifyKeysAll {
		val, err := resp.GetValueAsType(key, verifyTypes[key])
		if err != nil {
			t.Fatalf("%s failed: %e", testName, err)
		}
		if val != verifyValues[key] {
			t.Fatalf("%s failed: <%#v> expected %#v but received %#v", testName, key, verifyValues[key], val)
		}
	}
}

func TestGjrc_PatchJsonError(t *testing.T) {
	testName := "TestGjrc_PatchJsonError"
	client := NewGjrc(nil, 10*time.Second)
	url := "://localhost/patch"
	resp := client.PatchJson(url, map[string]interface{}{"key1": "value1", "key2": 2, "key3": true})
	if resp == nil {
		t.Fatalf("%s failed: nil response", testName)
	}
	if resp.Error() == nil {
		t.Fatalf("%s failed: expected error", testName)
	}
	if resp.HttpResponse() != nil {
		t.Fatalf("%s failed: expected nil http response", testName)
	}
}

func TestGjrc_PostJson(t *testing.T) {
	s := newJsonHttpServer(0)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	defer func() {
		if err := s.Shutdown(); err != nil {
			panic(err)
		}
	}()

	testName := "TestGjrc_PostJson"
	client := NewGjrc(nil, 10*time.Second)
	url := fmt.Sprintf("http://localhost:%d/post", s.listenPort)
	resp := client.PostJson(url, requestDataJson, requestHeaders...)
	if resp == nil {
		t.Fatalf("%s failed: nil response", testName)
	}

	valUrl, err := resp.GetValueAsType("url", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", testName, err)
	}
	if valUrl != url {
		t.Fatalf("%s failed: expected %s but received %s", testName, url, valUrl)
	}

	valHost, err := resp.GetValueAsType("headers.Host", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", testName, err)
	}
	host := "localhost"
	if valHost != host {
		t.Fatalf("%s failed: expected %s but received %s", testName, host, valHost)
	}

	for _, key := range verifyKeysAll {
		val, err := resp.GetValueAsType(key, verifyTypes[key])
		if err != nil {
			t.Fatalf("%s failed: %e", testName, err)
		}
		if val != verifyValues[key] {
			t.Fatalf("%s failed: <%#v> expected %#v but received %#v", testName, key, verifyValues[key], val)
		}
	}
}

func TestGjrc_PostJsonError(t *testing.T) {
	testName := "TestGjrc_PostJsonError"
	client := NewGjrc(nil, 10*time.Second)
	url := "://localhot/post"
	resp := client.PostJson(url, map[string]interface{}{"key1": "value1", "key2": 2, "key3": true})
	if resp == nil {
		t.Fatalf("%s failed: nil response", testName)
	}
	if resp.Error() == nil {
		t.Fatalf("%s failed: expected error", testName)
	}
	if resp.HttpResponse() != nil {
		t.Fatalf("%s failed: expected nil http response", testName)
	}
}

func TestGjrc_PostForm(t *testing.T) {
	s := newJsonHttpServer(0)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	defer func() {
		if err := s.Shutdown(); err != nil {
			panic(err)
		}
	}()

	testName := "TestGjrc_PostForm"
	client := NewGjrc(nil, 10*time.Second)
	_url := fmt.Sprintf("http://localhost:%d/post", s.listenPort)
	resp := client.PostForm(_url, url.Values{"key1": {"value1"}, "key2": {"2"}, "key3": {"true"}}, requestHeaders...)
	if resp == nil {
		t.Fatalf("%s failed: nil response", testName)
	}

	valUrl, err := resp.GetValueAsType("url", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", testName, err)
	}
	if valUrl != _url {
		t.Fatalf("%s failed: expected %s but received %s", testName, _url, valUrl)
	}

	valHost, err := resp.GetValueAsType("headers.Host", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", testName, err)
	}
	host := "localhost"
	if valHost != host {
		t.Fatalf("%s failed: expected %s but received %s", testName, host, valHost)
	}

	for _, key := range verifyKeysHeader {
		val, err := resp.GetValueAsType(key, verifyTypes[key])
		if err != nil {
			t.Fatalf("%s failed: %e", testName, err)
		}
		if val != verifyValues[key] {
			t.Fatalf("%s failed: <%#v> expected %#v but received %#v", testName, key, verifyValues[key], val)
		}
	}

	val1, err := resp.GetValueAsType("form.key1", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", testName, err)
	}
	v1 := "value1"
	if val1 != v1 {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, v1, val1)
	}

	val2, err := resp.GetValueAsType("form.key2", reddo.TypeInt)
	if err != nil {
		t.Fatalf("%s failed: %e", testName, err)
	}
	v2 := int64(2)
	if val2 != v2 {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, v2, val2)
	}

	val3, err := resp.GetValueAsType("form.key3", reddo.TypeBool)
	if err != nil {
		t.Fatalf("%s failed: %e", testName, err)
	}
	v3 := true
	if val3 != v3 {
		t.Fatalf("%s failed: expected %#v but received %#v", testName, v3, val3)
	}
}

func TestGjrc_PutJson(t *testing.T) {
	s := newJsonHttpServer(0)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	defer func() {
		if err := s.Shutdown(); err != nil {
			panic(err)
		}
	}()

	testName := "TestGjrc_PutJson"
	client := NewGjrc(nil, 10*time.Second)
	url := fmt.Sprintf("http://localhost:%d/put", s.listenPort)
	resp := client.PutJson(url, requestDataJson, requestHeaders...)
	if resp == nil {
		t.Fatalf("%s failed: nil response", testName)
	}

	valUrl, err := resp.GetValueAsType("url", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", testName, err)
	}
	if valUrl != url {
		t.Fatalf("%s failed: expected %s but received %s", testName, url, valUrl)
	}

	valHost, err := resp.GetValueAsType("headers.Host", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", testName, err)
	}
	host := "localhost"
	if valHost != host {
		t.Fatalf("%s failed: expected %s but received %s", testName, host, valHost)
	}

	for _, key := range verifyKeysAll {
		val, err := resp.GetValueAsType(key, verifyTypes[key])
		if err != nil {
			t.Fatalf("%s failed: %e", testName, err)
		}
		if val != verifyValues[key] {
			t.Fatalf("%s failed: <%#v> expected %#v but received %#v", testName, key, verifyValues[key], val)
		}
	}
}

func TestGjrc_PutJsonError(t *testing.T) {
	testName := "TestGjrc_PutJsonError"
	client := NewGjrc(nil, 10*time.Second)
	url := "://localhost/put"
	resp := client.PutJson(url, map[string]interface{}{"key1": "value1", "key2": 2, "key3": true})
	if resp == nil {
		t.Fatalf("%s failed: nil response", testName)
	}
	if resp.Error() == nil {
		t.Fatalf("%s failed: expected error", testName)
	}
	if resp.HttpResponse() != nil {
		t.Fatalf("%s failed: expected nil http response", testName)
	}
}

/*----------------------------------------------------------------------*/

func TestGjrc_Timeout_Default1(t *testing.T) {
	s := newJsonHttpServer(0)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	defer func() {
		if err := s.Shutdown(); err != nil {
			panic(err)
		}
	}()

	testName := "TestGjrc_TimeoutDefault1"
	client := NewGjrc(nil, 5*time.Second)
	url := fmt.Sprintf("http://localhost:%d/delay?time=3", s.listenPort)
	resp := client.Get(url)
	if resp == nil {
		t.Fatalf("%s failed: nil response", testName)
	}
	if resp.Error() != nil {
		t.Fatalf("%s failed: %e", testName, resp.Error())
	}
	if resp.StatusCode() != 200 {
		t.Fatalf("%s failed: expected status code %#v but received %#v", testName, 200, resp.StatusCode())
	}
}

func TestGjrc_Timeout_Default2(t *testing.T) {
	s := newJsonHttpServer(0)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	defer func() {
		if err := s.Shutdown(); err != nil {
			panic(err)
		}
	}()

	testName := "TestGjrc_Timeout_Default2"
	client := NewGjrc(nil, 5*time.Second)
	url := fmt.Sprintf("http://localhost:%d/delay?time=7", s.listenPort)
	resp := client.Get(url)
	if resp == nil {
		t.Fatalf("%s failed: nil response", testName)
	}
	if resp.Error() == nil {
		t.Fatalf("%s failed: expecting error", testName)
	}
	fmt.Printf("\t[DEBUG] %s\n", resp.Error())
}

func TestGjrc_Timeout_Custom1(t *testing.T) {
	s := newJsonHttpServer(0)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	defer func() {
		if err := s.Shutdown(); err != nil {
			panic(err)
		}
	}()

	testName := "TestGjrc_Timeout_Custom1"
	client := NewGjrc(nil, 10*time.Second)
	url := fmt.Sprintf("http://localhost:%d/delay?time=3", s.listenPort)
	resp := client.Get(url, RequestMeta{Timeout: 5 * time.Second})
	if resp == nil {
		t.Fatalf("%s failed: nil response", testName)
	}
	if resp.Error() != nil {
		t.Fatalf("%s failed: %e", testName, resp.Error())
	}
	if resp.StatusCode() != 200 {
		t.Fatalf("%s failed: expected status code %#v but received %#v", testName, 200, resp.StatusCode())
	}
}

func TestGjrc_Timeout_Custom2(t *testing.T) {
	s := newJsonHttpServer(0)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	defer func() {
		if err := s.Shutdown(); err != nil {
			panic(err)
		}
	}()

	testName := "TestGjrc_Timeout_Custom2"
	client := NewGjrc(nil, 10*time.Second)
	url := fmt.Sprintf("http://localhost:%d/delay?time=7", s.listenPort)
	resp := client.Get(url, RequestMeta{Timeout: 5 * time.Second})
	if resp == nil {
		t.Fatalf("%s failed: nil response", testName)
	}
	if resp.Error() == nil {
		t.Fatalf("%s failed: expecting error", testName)
	}
	fmt.Printf("\t[DEBUG] %s\n", resp.Error())
}

/*----------------------------------------------------------------------*/

func TestGjrc_Unmarshal(t *testing.T) {
	s := newJsonHttpServer(0)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	defer func() {
		if err := s.Shutdown(); err != nil {
			panic(err)
		}
	}()

	testName := "TestGjrc_Unmarshal"
	client := NewGjrc(nil, 10*time.Second)
	url := fmt.Sprintf("http://localhost:%d/get", s.listenPort)
	resp := client.Get(url, requestHeaders...)
	if resp == nil {
		t.Fatalf("%s failed: nil response", testName)
	}

	type mystruct struct {
		Headers map[string]string `json:"headers"`
		Method  string            `json:"method"`
		Time    string            `json:"time"`
		Url     string            `json:"url"`
	}

	var v1 mystruct
	err := resp.Unmarshal(&v1)
	if err != nil {
		t.Fatalf("%s failed: %e", testName, err)
	}
	if v1.Url != url {
		t.Fatalf("%s failed: expected %s but received %s", testName, url, v1.Url)
	}
	if v1.Method != "get" {
		t.Fatalf("%s failed: expected %s but received %s", testName, "get", v1.Method)
	}
	if v1.Headers["Host"] != "localhost" {
		t.Fatalf("%s failed: expected %s but received %s", testName, "localhost", v1.Headers["host"])
	}

	var v2 *mystruct
	err = resp.Unmarshal(&v2)
	if err != nil {
		t.Fatalf("%s failed: %e", testName, err)
	}
	if v2 == nil {
		t.Fatalf("%s failed: nil", testName)
	}
	if v2.Url != url {
		t.Fatalf("%s failed: expected %s but received %s", testName, url, v2.Url)
	}
	if v2.Method != "get" {
		t.Fatalf("%s failed: expected %s but received %s", testName, "get", v2.Method)
	}
	if v2.Headers["Host"] != "localhost" {
		t.Fatalf("%s failed: expected %s but received %s", testName, "localhost", v2.Headers["host"])
	}
}
