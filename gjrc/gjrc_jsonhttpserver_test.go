package gjrc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"time"
)

func (s *jsonHttpServer) serveRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
}

func (s *jsonHttpServer) serveNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func (s *jsonHttpServer) serveDELAY(w http.ResponseWriter, r *http.Request) {
	delayTime, err := strconv.Atoi(r.URL.Query().Get("time"))
	if err == nil && delayTime > 0 {
		time.Sleep(time.Duration(delayTime) * time.Second)
	}
	respData := s.commonResponse(r, r.Method)
	js, _ := json.Marshal(respData)
	w.Write(js)
}

func (s *jsonHttpServer) commonResponse(r *http.Request, method string) map[string]interface{} {
	headers := map[string]string{"Host": "localhost"}
	for k := range r.Header {
		headers[k] = r.Header.Get(k)
	}
	respData := map[string]interface{}{
		"method":  method,
		"headers": headers,
		"time":    time.Now().Format(time.RFC3339),
		"url":     fmt.Sprintf("http://localhost:%d/%s", s.listenPort, method),
	}
	if strings.ToUpper(method) != "GET" && strings.ToLower(headers["Content-Type"]) == "application/json" {
		body, _ := ioutil.ReadAll(r.Body)
		respData["data"] = string(body)

		var jsonData interface{}
		json.Unmarshal(body, &jsonData)
		respData["json"] = jsonData
	}
	if strings.ToUpper(method) == "POST" || strings.ToUpper(method) == "PUT" || strings.ToUpper(method) == "PATCH" &&
		strings.ToLower(headers["Content-Type"]) == "application/x-www-form-urlencoded" {
		r.ParseForm()
		respData["data"] = r.PostForm.Encode()

		formData := map[string]string{}
		for k := range r.PostForm {
			formData[k] = r.PostFormValue(k)
		}
		respData["form"] = formData
	}
	return respData
}

func (s *jsonHttpServer) serveDELETE(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	respData := s.commonResponse(r, "delete")
	js, _ := json.Marshal(respData)
	w.Write(js)
}

func (s *jsonHttpServer) serveGET(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	respData := s.commonResponse(r, "get")
	js, _ := json.Marshal(respData)
	w.Write(js)
}

func (s *jsonHttpServer) servePATCH(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	respData := s.commonResponse(r, "patch")
	js, _ := json.Marshal(respData)
	w.Write(js)
}

func (s *jsonHttpServer) servePOST(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	respData := s.commonResponse(r, "post")
	js, _ := json.Marshal(respData)
	w.Write(js)
}

func (s *jsonHttpServer) servePUT(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	respData := s.commonResponse(r, "put")
	js, _ := json.Marshal(respData)
	w.Write(js)
}

type jsonHttpServer struct {
	listenPort int
	server     *httptest.Server
}

func (s *jsonHttpServer) ListenAndServe() error {
	srvmx := http.NewServeMux()
	srvmx.HandleFunc("/", s.serveRoot)
	srvmx.HandleFunc("/notfound", s.serveNotFound)
	srvmx.HandleFunc("/delete", s.serveDELETE)
	srvmx.HandleFunc("/get", s.serveGET)
	srvmx.HandleFunc("/patch", s.servePATCH)
	srvmx.HandleFunc("/post", s.servePOST)
	srvmx.HandleFunc("/put", s.servePUT)
	srvmx.HandleFunc("/delay", s.serveDELAY)

	s.server = httptest.NewUnstartedServer(srvmx)
	var err error
	listenHostAndPort := fmt.Sprintf("127.0.0.1:%d", s.listenPort)
	s.server.Listener, err = net.Listen("tcp", listenHostAndPort)
	if err != nil {
		return err
	}
	fmt.Printf("\t[DEBUG] running HTTP server on <%s>\n", listenHostAndPort)
	s.server.Start()
	return nil
}

func (s *jsonHttpServer) Shutdown() error {
	s.server.Close()
	return nil
}

func newJsonHttpServer(port int) *jsonHttpServer {
	rand.Seed(time.Now().UnixNano())
	if port <= 0 {
		port = 1024 + rand.Intn(10000-1024)
	}
	return &jsonHttpServer{listenPort: port}
}
