package helper

import (
	"net/http"
	"testing"
)

type CookieTransport struct {
	t         *testing.T
	Cookie    string
	Transport http.RoundTripper
}

func NewCookieTransport(
	t *testing.T,
	cookie string,
) *CookieTransport {
	return &CookieTransport{
		t:         t,
		Cookie:    cookie,
		Transport: http.DefaultTransport,
	}
}

func (ct *CookieTransport) transport() http.RoundTripper {
	ct.t.Helper()

	return ct.Transport
}

func (ct *CookieTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	ct.t.Helper()

	req.Header.Add("Cookie", ct.Cookie)

	resp, err := ct.transport().RoundTrip(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

type CookiesTransport struct {
	T         *testing.T
	Cookies   []*http.Cookie
	Transport http.RoundTripper
}

func NewCookiesTransport(
	t *testing.T,
	cookies []*http.Cookie,
) *CookiesTransport {
	t.Helper()

	return &CookiesTransport{
		T:         t,
		Cookies:   cookies,
		Transport: http.DefaultTransport,
	}
}

func (ct *CookiesTransport) transport() http.RoundTripper {
	ct.T.Helper()

	return ct.Transport
}

func (ct *CookiesTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	ct.T.Helper()

	for _, cookie := range ct.Cookies {
		req.AddCookie(cookie)
	}

	resp, err := ct.transport().RoundTrip(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
