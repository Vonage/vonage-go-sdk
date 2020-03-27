package nexmo

import "net/http"

type APITransport struct {
	APISecret string

	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

// RoundTrip implements the RoundTripper interface.
func (t *APITransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// To set extra querystring params, we must make a copy of the Request so
	// that we don't modify the Request we were given. This is required by the
	// specification of http.RoundTripper.
	req = cloneRequest(req)
	q := req.URL.Query()
	q.Set("api_secret", t.APISecret)
	req.URL.RawQuery = q.Encode()

	// Make the HTTP request.
	return t.transport().RoundTrip(req)
}

func (t *APITransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *APITransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

// cloneRequest returns a clone of the provided *http.Request. The clone is a
// shallow copy of the struct and its Header map.
func cloneRequest(r *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	return r2
}
