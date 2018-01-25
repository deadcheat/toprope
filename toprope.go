package toprope

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"time"
)

const (
	modeTCP       = "tcp"
	retryToListen = 3
)

// memo cache servers to manage close and listen
var memo = make(map[string]*Server)

// Server server instance used internal
type Server struct {
	i  *httptest.Server
	wg *sync.WaitGroup
}

// URL access internal property
func (s *Server) URL() string {
	return s.i.URL
}

// Listener access internal property
func (s *Server) Listener() net.Listener {
	return s.i.Listener
}

// Config access internal property
func (s *Server) Config() *http.Server {
	return s.i.Config
}

// Client Override and call internal
func (s *Server) Client() *http.Client {
	return s.i.Client()
}

// Close Override and call internal
func (s *Server) Close() {
	s.i.Close()
	s.wg.Done()
}

// CloseClientConnections Override and call internal
func (s *Server) CloseClientConnections() {
	s.i.CloseClientConnections()
}

// Start Override and call internal
func (s *Server) Start() {
	s.i.Start()
}

// NewHttptestTCPServerFromURL generate *httptest.Server from url-string
func NewHttptestTCPServerFromURL(urlstring string, handler http.Handler) (ts *Server, err error) {
	u, err := url.Parse(urlstring)
	if err != nil {
		return
	}
	return listenAndCreateServer(u, handler)
}

// NewHttptestTCPServer generates *httptest.Server from hostname and portnum
func NewHttptestTCPServer(hostName string, port int, handler http.Handler) (ts *Server, err error) {
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

func listenAndCreateServer(u *url.URL, handler http.Handler) (ts *Server, err error) {
	host := u.Host
	s, ok := memo[host]
	if ok && s.wg != nil {
		// if memo has same hostname, wait listener to close certainly
		s.wg.Wait()
	}
	var l net.Listener
	// 規定回数リトライする
	for index := 0; index < retryToListen; index++ {
		l, err = net.Listen(modeTCP, host)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
	if err != nil {
		return
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	server := &httptest.Server{
		Listener: l,
		Config:   &http.Server{Handler: handler},
	}
	ts = &Server{
		i:  server,
		wg: wg,
	}
	memo[host] = ts
	return
}
