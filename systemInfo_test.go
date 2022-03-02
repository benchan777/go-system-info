package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomePageHandlerStatusCode(t *testing.T) {
	var tests = []struct {
		url                string
		expectedStatusCode int
	}{
		{"/", 200},
	}

	for i := 0; i < len(tests); i++ {
		req, _ := http.NewRequest("GET", tests[i].url, nil) // Generate request
		res := httptest.NewRecorder()                       // Create response recorder

		handler := http.HandlerFunc(homePageHandler)
		handler.ServeHTTP(res, req)

		if res.Code != tests[i].expectedStatusCode {
			t.Error("Test Failed: expected: " + fmt.Sprint(tests[i].expectedStatusCode) + " got: " + fmt.Sprint(res.Code))
		}
	}
}

func TestGetDataHandlerStatusCode(t *testing.T) {
	var tests = []struct {
		url                string
		expectedStatusCode int
	}{
		{"/api", 200},
	}

	for i := 0; i < len(tests); i++ {
		req, _ := http.NewRequest("GET", tests[i].url, nil)
		res := httptest.NewRecorder()

		handler := http.HandlerFunc(getDataHandler)
		handler.ServeHTTP(res, req)

		if res.Code != tests[i].expectedStatusCode {
			t.Error("Test Failed: expected: " + fmt.Sprint(tests[i].expectedStatusCode) + " got: " + fmt.Sprint(res.Code))
		}
	}
}

func TestGetDataHandler(t *testing.T) {
	var tests = []struct {
		apiUrl       string
		expectedData string
	}{
		{"/api", "HostName"},
		{"/api", "CPUName"},
		{"/api", "DiskCapacity"},
		{"/api", "DiskUsage"},
		{"/api", "DiskFree"},
		{"/api", "RamCapacity"},
		{"/api", "RamAvailable"},
	}

	for i := 0; i < len(tests); i++ {
		req, _ := http.NewRequest("GET", tests[i].apiUrl, nil)
		res := httptest.NewRecorder()

		handler := http.HandlerFunc(getDataHandler)
		handler.ServeHTTP(res, req)

		data := res.Body.String()
		if !strings.Contains(data, tests[i].expectedData) {
			t.Error("Test Failed: expected: " + tests[i].expectedData + " to be in string but could not be found.")
		}
	}
}
