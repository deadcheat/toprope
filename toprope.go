package toprope

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
)

const (
	modeTCP = "tcp"
)

// NewHttptestTCPServerFromURL generate *httptest.Server from url-string
func NewHttptestTCPServerFromURL(urlstring string, handler http.Handler) (ts *httptest.Server, err error) {
	ur, err := url.Parse(urlstring)
	if err != nil {
		return
	}
	l, err := net.Listen(modeTCP, ur.Host)
	if err != nil {
		return
	}
	ts = &httptest.Server{
		Listener: l,
		Config:   &http.Server{Handler: handler},
	}
	return
}

// NewHttptestTCPServer generates *httptest.Server from hostname and portnum
func NewHttptestTCPServer(hostName string, port int, handler http.Handler) (ts *httptest.Server, err error) {
	host := hostName
	if port >= 0 {
		host = fmt.Sprintf("%s:%d", hostName, port)
	}
	l, err := net.Listen(modeTCP, host)
	if err != nil {
		return
	}
	ts = &httptest.Server{
		Listener: l,
		Config:   &http.Server{Handler: handler},
	}
	return
}
