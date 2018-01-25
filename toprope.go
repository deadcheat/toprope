package toprope

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"time"
)

const (
	modeTCP       = "tcp"
	retryToListen = 3
)

// NewHttptestTCPServerFromURL generate *httptest.Server from url-string
func NewHttptestTCPServerFromURL(urlstring string, handler http.Handler) (ts *httptest.Server, err error) {
	u, err := url.Parse(urlstring)
	if err != nil {
		return
	}
	return listenAndCreateServer(u, handler)
}

// NewHttptestTCPServer generates *httptest.Server from hostname and portnum
func NewHttptestTCPServer(hostName string, port int, handler http.Handler) (ts *httptest.Server, err error) {
	host := hostName
	if port >= 0 {
		host = fmt.Sprintf("%s:%d", hostName, port)
	}
	u, err := url.Parse(host)
	if err != nil {
		return
	}
	return listenAndCreateServer(u, handler)
}

func listenAndCreateServer(u *url.URL, handler http.Handler) (ts *httptest.Server, err error) {
	var l net.Listener
	// 規定回数リトライする
	for index := 0; index < retryToListen; index++ {
		l, err = net.Listen(modeTCP, u.Host)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
	if err != nil {
		return
	}
	ts = &httptest.Server{
		URL:      u.String(),
		Listener: l,
		Config:   &http.Server{Handler: handler},
	}
	return
}
