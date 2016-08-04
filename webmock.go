package httpmock

import (
	"net/http/httptest"
	"net/http"
	"fmt"
	"net/url"
    "strings"
)

// DummyResponse represents a response that the server can send as answer to a response.
type DummyResponse struct {
	Code int
	HeaderValues map[string][]string
	Body string
}

// ServeGeneric will return the same output regardless the URL requested or anything else (body text, header and/or 
// forms).
func ServeGeneric(input DummyResponse) (*httptest.Server, *http.Client)  {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for key, values := range input.HeaderValues {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}
		w.WriteHeader(input.Code)
		fmt.Fprintln(w, input.Body)}))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {

            // If the request host address is the server URL, the client must redirect the request to our server.
            // Otherwise, do not proxy.
            if strings.Contains(server.URL, req.URL.Host) {
				return url.Parse(server.URL)
			} else {
				return nil, nil
			}}}

	httpClient := &http.Client{Transport: transport}
	return server, httpClient
}

// ServeMulti will return a response according to the path defined as key to that response in input parameter. Body and
// header content is ignored.
func ServeMulti(input map[string]DummyResponse) (*httptest.Server, *http.Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for key, values := range input[r.URL.Path].HeaderValues {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}
		w.WriteHeader(input[r.URL.Path].Code)
		fmt.Fprintln(w, input[r.URL.Path].Body)}))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)}}

	httpClient := &http.Client{Transport: transport}
	return server, httpClient
}
