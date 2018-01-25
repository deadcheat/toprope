package toprope_test

import (
	"fmt"
	"net"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/deadcheat/toprope"
)

// Test NewHttptestTCPServerFromURL will return successfully
func TestNewHttptestTCPServerFromURL_Success(t *testing.T) {
	testUrl := "http://127.0.0.1:9999"
	// error must be nil
	ts, err := toprope.NewHttptestTCPServerFromURL(testUrl, nil)
	if err != nil {
		t.Error("Occurred unexpected error : ", err)
		t.Fail()
	}
	defer func() {
		ts.CloseClientConnections()
		ts.Close()
	}()
	ts.Start()
	// ts must return correct URL
	if ts.URL() != testUrl || ts.Listener() == nil || ts.Config() == nil {
		t.Error("Server created with unexpected url:", ts.URL())
		t.Fail()
	}

}

// Test NewHttptestTCPServerFromURL will return error when failed to parse url
func TestNewHttptestTCPServerFromURL_ParseError(t *testing.T) {
	testUrl := "://localhost:9999"
	// error must be nil
	_, err := toprope.NewHttptestTCPServerFromURL(testUrl, nil)
	if err == nil {
		t.Error("toprope must return error but returned nil")
		t.Fail()
	}
	if _, ok := err.(*url.Error); !ok {
		t.Error("returned error must be typed *url.Error")
		t.Fail()
	}
}

// Test NewHttptestTCPServer will return successfully
func TestNewHttptestTCPServer_Success(t *testing.T) {
	testHost := "http://127.0.0.1"
	testPort := 9999
	// error must be nil
	ts, err := toprope.NewHttptestTCPServer(testHost, testPort, nil)
	if err != nil {
		t.Error("Occurred unexpected error : ", err)
		t.Fail()
	}
	defer func() {
		ts.CloseClientConnections()
		ts.Close()
	}()
	expectedURL := fmt.Sprintf("%s:%d", testHost, testPort)
	ts.Start()
	// ts must return correct URL
	if ts.URL() != expectedURL {
		t.Error("Server created with unexpected url:", ts.URL())
		t.Fail()
	}

}

// Test NewHttptestTCPServer will return error when failed to parse url
func TestNewHttptestTCPServer_ParseError(t *testing.T) {
	testHost := "://127.0.0.1"
	testPort := 9999
	// error must be nil
	_, err := toprope.NewHttptestTCPServer(testHost, testPort, nil)
	if err == nil {
		t.Error("toprope must return error but returned nil")
		t.Fail()
	}
	if _, ok := err.(*url.Error); !ok {
		t.Error("returned error must be typed *url.Error")
		t.Fail()
	}
}

// Test NewHttptestTCPServerFromURL will return error when failed to listen port
func TestNewHttptestTCPServerFromURL_ListenError(t *testing.T) {
	ts := httptest.NewServer(nil)
	defer func() {
		ts.CloseClientConnections()
		ts.Close()
	}()
	testURL := ts.URL
	// error must be nil
	_, err := toprope.NewHttptestTCPServerFromURL(testURL, nil)
	if err == nil {
		t.Error("toprope must return error but returned nil")
		t.Fail()
	}
	if _, ok := err.(*net.OpError); !ok {
		t.Error("returned error must be typed *net.OpError")
		t.Fail()
	}
}

// Test NewHttptestTCPServerFromURL will return error when duplicate start using toprope
func TestNewHttptestTCPServerFromURL_ShouldSuccessDuplicatedRunning(t *testing.T) {
	testUrl := "http://127.0.0.1:9999"
	// error must be nil
	ts, err := toprope.NewHttptestTCPServerFromURL(testUrl, nil)
	go func() {
		time.Sleep(4 * time.Second)
		ts.CloseClientConnections()
		ts.Close()
	}()
	delayedTs, err := toprope.NewHttptestTCPServerFromURL(testUrl, nil)
	if err != nil {
		t.Error("Occurred unexpected error : ", err)
		t.Fail()
	}
	defer func() {
		delayedTs.CloseClientConnections()
		delayedTs.Close()
	}()
}

func TestInternals(t *testing.T) {
	testUrl := "http://127.0.0.1:9999"
	// error must be nil
	ts, _ := toprope.NewHttptestTCPServerFromURL(testUrl, nil)
	if ts.URL() != "" {
		t.Error("URL() must return empty value before start")
		t.Fail()
	}
	ts.Start()
	if ts.URL() == "" {
		t.Error("URL() must return not empty value after start")
		t.Fail()
	}
	ts.Close()
}
