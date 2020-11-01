package gjrc

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/btnguyen2k/consu/reddo"
)

func TestNewGjrc(t *testing.T) {
	name := "TestNewGjrc"
	client := NewGjrc(nil, 10*time.Second)
	if client == nil {
		t.Fatalf("%s failed: nil", name)
	}
	if client.httpClient == nil {
		t.Fatalf("%s failed: httpClient is nil", name)
	}
}

func TestNewGjrc_PrebuiltHttpClient(t *testing.T) {
	name := "TestNewGjrc_PrebuiltHttpClient"
	httpClient := &http.Client{}
	client := NewGjrc(httpClient, 0)
	if client == nil {
		t.Fatalf("%s failed: nil", name)
	}
	if client.httpClient != httpClient {
		t.Fatalf("%s failed: pre-built http.Client is not used", name)
	}
}

func TestGjrcResponse_Error(t *testing.T) {
	name := "TestGjrcResponse_Error"
	client := NewGjrc(nil, 10*time.Second)
	url := "https://mydomain.notreal"
	resp := client.Get(url)
	if resp == nil {
		t.Fatalf("%s failed: nil response", name)
	}
	if resp.Error() == nil {
		t.Fatalf("%s failed: nil error", name)
	}
}

func TestGjrcResponse_StatusCode(t *testing.T) {
	name := "TestGjrcResponse_StatusCode"
	client := NewGjrc(nil, 10*time.Second)
	url := "https://httpbin.org/notfound"
	resp := client.Get(url)
	if resp == nil {
		t.Fatalf("%s failed: nil response", name)
	}
	if resp.Error() != nil {
		t.Fatalf("%s failed: %e", name, resp.Error())
	}
	if resp.StatusCode() != 404 {
		t.Fatalf("%s failed: expected status code %#v but received %#v", name, 404, resp.StatusCode())
	}
}

func TestGjrcResponse_Body(t *testing.T) {
	name := "TestGjrcResponse_Body"
	client := NewGjrc(nil, 10*time.Second)
	url := "https://httpbin.org/get"
	resp := client.Get(url)
	if resp == nil {
		t.Fatalf("%s failed: nil response", name)
	}
	if resp.Error() != nil {
		t.Fatalf("%s failed: %e", name, resp.Error())
	}
	if resp.StatusCode() != 200 {
		t.Fatalf("%s failed: expected status code %#v but received %#v", name, 404, resp.StatusCode())
	}
	body, err := resp.Body()
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	if !strings.HasPrefix(string(body), "{") {
		t.Fatalf("%s failed: unexpected response body %s", name, body)
	}
}

func TestGjrcResponse_HttpResponse(t *testing.T) {
	name := "TestGjrcResponse_HttpResponse"
	client := NewGjrc(nil, 10*time.Second)
	url := "https://httpbin.org/get"
	resp := client.Get(url)
	if resp == nil {
		t.Fatalf("%s failed: nil response", name)
	}
	if resp.Error() != nil {
		t.Fatalf("%s failed: %e", name, resp.Error())
	}
	if resp.StatusCode() != 200 {
		t.Fatalf("%s failed: expected status code %#v but received %#v", name, 404, resp.StatusCode())
	}
	if resp.HttpResponse() == nil {
		t.Fatalf("%s failed: nil http.Response", name)
	}
}

func TestGjrc_DeleteJson(t *testing.T) {
	name := "TestGjrc_DeleteJson"
	client := NewGjrc(nil, 10*time.Second)
	url := "https://httpbin.org/delete"
	resp := client.DeleteJson(url, map[string]interface{}{"key1": "value", "key2": 1, "key3": true})
	if resp == nil {
		t.Fatalf("%s failed: nil response", name)
	}

	valUrl, err := resp.GetValueAsType("url", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	if valUrl != url {
		t.Fatalf("%s failed: expected %s but received %s", name, url, valUrl)
	}

	valHost, err := resp.GetValueAsType("headers.Host", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	host := "httpbin.org"
	if valHost != host {
		t.Fatalf("%s failed: expected %s but received %s", name, host, valHost)
	}

	val1, err := resp.GetValueAsType("json.key1", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	v1 := "value"
	if val1 != v1 {
		t.Fatalf("%s failed: expected %#v but received %#v", name, v1, val1)
	}

	val2, err := resp.GetValueAsType("json.key2", reddo.TypeInt)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	v2 := int64(1)
	if val2 != v2 {
		t.Fatalf("%s failed: expected %#v but received %#v", name, v2, val2)
	}

	val3, err := resp.GetValueAsType("json.key3", reddo.TypeBool)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	v3 := true
	if val3 != v3 {
		t.Fatalf("%s failed: expected %#v but received %#v", name, v3, val3)
	}
}

func TestGjrc_Get(t *testing.T) {
	name := "TestGjrc_Get"
	client := NewGjrc(nil, 10*time.Second)
	url := "https://httpbin.org/get"
	resp := client.Get(url)
	if resp == nil {
		t.Fatalf("%s failed: nil response", name)
	}

	valUrl, err := resp.GetValueAsType("url", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	if valUrl != url {
		t.Fatalf("%s failed: expected %s but received %s", name, url, valUrl)
	}

	valHost, err := resp.GetValueAsType("headers.Host", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	host := "httpbin.org"
	if valHost != host {
		t.Fatalf("%s failed: expected %s but received %s", name, host, valHost)
	}
}

func TestGjrc_PatchJson(t *testing.T) {
	name := "TestGjrc_PatchJson"
	client := NewGjrc(nil, 10*time.Second)
	url := "https://httpbin.org/patch"
	resp := client.PatchJson(url, map[string]interface{}{"key1": "value", "key2": 1, "key3": true})
	if resp == nil {
		t.Fatalf("%s failed: nil response", name)
	}

	valUrl, err := resp.GetValueAsType("url", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	if valUrl != url {
		t.Fatalf("%s failed: expected %s but received %s", name, url, valUrl)
	}

	valHost, err := resp.GetValueAsType("headers.Host", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	host := "httpbin.org"
	if valHost != host {
		t.Fatalf("%s failed: expected %s but received %s", name, host, valHost)
	}

	val1, err := resp.GetValueAsType("json.key1", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	v1 := "value"
	if val1 != v1 {
		t.Fatalf("%s failed: expected %#v but received %#v", name, v1, val1)
	}

	val2, err := resp.GetValueAsType("json.key2", reddo.TypeInt)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	v2 := int64(1)
	if val2 != v2 {
		t.Fatalf("%s failed: expected %#v but received %#v", name, v2, val2)
	}

	val3, err := resp.GetValueAsType("json.key3", reddo.TypeBool)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	v3 := true
	if val3 != v3 {
		t.Fatalf("%s failed: expected %#v but received %#v", name, v3, val3)
	}
}

func TestGjrc_PostJson(t *testing.T) {
	name := "TestGjrc_PostJson"
	client := NewGjrc(nil, 10*time.Second)
	url := "https://httpbin.org/post"
	resp := client.PostJson(url, map[string]interface{}{"key1": "value", "key2": 1, "key3": true})
	if resp == nil {
		t.Fatalf("%s failed: nil response", name)
	}

	valUrl, err := resp.GetValueAsType("url", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	if valUrl != url {
		t.Fatalf("%s failed: expected %s but received %s", name, url, valUrl)
	}

	valHost, err := resp.GetValueAsType("headers.Host", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	host := "httpbin.org"
	if valHost != host {
		t.Fatalf("%s failed: expected %s but received %s", name, host, valHost)
	}

	val1, err := resp.GetValueAsType("json.key1", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	v1 := "value"
	if val1 != v1 {
		t.Fatalf("%s failed: expected %#v but received %#v", name, v1, val1)
	}

	val2, err := resp.GetValueAsType("json.key2", reddo.TypeInt)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	v2 := int64(1)
	if val2 != v2 {
		t.Fatalf("%s failed: expected %#v but received %#v", name, v2, val2)
	}

	val3, err := resp.GetValueAsType("json.key3", reddo.TypeBool)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	v3 := true
	if val3 != v3 {
		t.Fatalf("%s failed: expected %#v but received %#v", name, v3, val3)
	}
}

func TestGjrc_PostForm(t *testing.T) {
	name := "TestGjrc_PostForm"
	client := NewGjrc(nil, 10*time.Second)
	_url := "https://httpbin.org/post"
	resp := client.PostForm(_url, url.Values{"key1": {"value"}, "key2": {"1"}, "key3": {"true"}})
	if resp == nil {
		t.Fatalf("%s failed: nil response", name)
	}

	valUrl, err := resp.GetValueAsType("url", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	if valUrl != _url {
		t.Fatalf("%s failed: expected %s but received %s", name, _url, valUrl)
	}

	valHost, err := resp.GetValueAsType("headers.Host", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	host := "httpbin.org"
	if valHost != host {
		t.Fatalf("%s failed: expected %s but received %s", name, host, valHost)
	}

	val1, err := resp.GetValueAsType("form.key1", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	v1 := "value"
	if val1 != v1 {
		t.Fatalf("%s failed: expected %#v but received %#v", name, v1, val1)
	}

	val2, err := resp.GetValueAsType("form.key2", reddo.TypeInt)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	v2 := int64(1)
	if val2 != v2 {
		t.Fatalf("%s failed: expected %#v but received %#v", name, v2, val2)
	}

	val3, err := resp.GetValueAsType("form.key3", reddo.TypeBool)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	v3 := true
	if val3 != v3 {
		t.Fatalf("%s failed: expected %#v but received %#v", name, v3, val3)
	}
}

func TestGjrc_PutJson(t *testing.T) {
	name := "TestGjrc_PutJson"
	client := NewGjrc(nil, 10*time.Second)
	url := "https://httpbin.org/put"
	resp := client.PutJson(url, map[string]interface{}{"key1": "value", "key2": 1, "key3": true})
	if resp == nil {
		t.Fatalf("%s failed: nil response", name)
	}

	valUrl, err := resp.GetValueAsType("url", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	if valUrl != url {
		t.Fatalf("%s failed: expected %s but received %s", name, url, valUrl)
	}

	valHost, err := resp.GetValueAsType("headers.Host", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	host := "httpbin.org"
	if valHost != host {
		t.Fatalf("%s failed: expected %s but received %s", name, host, valHost)
	}

	val1, err := resp.GetValueAsType("json.key1", reddo.TypeString)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	v1 := "value"
	if val1 != v1 {
		t.Fatalf("%s failed: expected %#v but received %#v", name, v1, val1)
	}

	val2, err := resp.GetValueAsType("json.key2", reddo.TypeInt)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	v2 := int64(1)
	if val2 != v2 {
		t.Fatalf("%s failed: expected %#v but received %#v", name, v2, val2)
	}

	val3, err := resp.GetValueAsType("json.key3", reddo.TypeBool)
	if err != nil {
		t.Fatalf("%s failed: %e", name, err)
	}
	v3 := true
	if val3 != v3 {
		t.Fatalf("%s failed: expected %#v but received %#v", name, v3, val3)
	}
}
