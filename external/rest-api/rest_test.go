package rest_api

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

// MockClient is the mock client
type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

var (
	// GetDoFunc fetches the mock client's `Do` func
	GetDoFunc func(req *http.Request) (*http.Response, error)
)

// Do is the mock client's `Do` func
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}

func TestConfig_ExecuteApiCall(t *testing.T) {
	Client = &MockClient{}
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		resp := &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader("test")),
		}
		return resp, nil

	}
	// Testing POST
	cfg := Config{
		Url:                  "https://localhost",
		BearerToken:          "some bearer token",
		RequestType:          "POST",
		Payload:              "{\"name\":\"test\",\"salary\":\"123\",\"age\":\"23\"}",
		ExpectedResponseCode: "200",
	}
	resp, err := cfg.ExecuteApiCall()
	if err != nil {
		t.Fatal(err)
	}
	// Testing to ensure that the body is correctly parsed
	bodyBytes, err := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	if bodyString != "test" {
		t.Errorf("Response body is %s", bodyString)
	}

	// Testing PUT
	cfg = Config{
		Url:                  "https://localhost",
		Username:             "admin",
		Password:             "admin",
		RequestType:          "PUT",
		Payload:              "{\"name\":\"test\",\"salary\":\"123\",\"age\":\"23\"}",
		ExpectedResponseCode: "200",
	}
	resp, err = cfg.ExecuteApiCall()
	if err != nil {
		t.Fatal(err)
	}

	// Testing GET
	cfg = Config{
		Url:                  "https://localhost",
		RequestType:          "GET",
		ExpectedResponseCode: "200",
	}
	resp, err = cfg.ExecuteApiCall()
	if err != nil {
		t.Fatal(err)
	}

	//Testing DELETE
	cfg = Config{
		Url:                  "https://localhost",
		RequestType:          "DELETE",
		ExpectedResponseCode: "200",
	}
	resp, err = cfg.ExecuteApiCall()
	if err != nil {
		t.Fatal(err)
	}

	// Switching mock client response code for expected response code
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		resp := &http.Response{
			StatusCode: 400,
		}
		return resp, nil

	}

	// Ensuring having an expected response code outside 200-299 works
	cfg = Config{
		Url:                  "https://localhost",
		RequestType:          "DELETE",
		ExpectedResponseCode: "400",
	}
	resp, err = cfg.ExecuteApiCall()
	if err != nil {
		t.Fatalf("Failed testing expected response outside of 200-299: %v", err)
	}

	// Switching mock client response code for other expected errors
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		resp := &http.Response{
			StatusCode: 400,
		}
		return resp, errors.New(
			"error testing",
		)

	}

	// Test to ensure it fails for a bad request type
	cfg = Config{
		Url:         "https://localhost",
		RequestType: "badType",
	}
	resp, err = cfg.ExecuteApiCall()
	if err == nil {
		t.Fatalf("Failed for incorrect type")
	}

	// Test to ensure it fails if expectedResponseCode doesn't match the response when not specified
	cfg = Config{
		Url:                  "https://localhost",
		RequestType:          "GET",
		ExpectedResponseCode: "405",
	}
	resp, err = cfg.ExecuteApiCall()
	if err == nil {
		t.Fatal("Failed not matching ExpectedResponseCode")
	}

	// Test to ensure it fails if response code is outside 200-299 when not providing an expected response code
	cfg = Config{
		Url:         "https://localhost",
		RequestType: "GET",
	}
	resp, err = cfg.ExecuteApiCall()
	if err == nil {
		t.Fatalf("Failed for being outside of 200-299: %v", err)
	}

}
