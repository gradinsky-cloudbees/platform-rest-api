package rest_api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

func (c *Config) ExecuteApiCall(ctx context.Context) (string, error) {
	var req *http.Request
	var err error
	if c.RequestType == "GET" {
		req, err = GetBuilder(c.Url)
	} else if c.RequestType == "POST" {
		req, err = PostBuilder(c.Url, c.Payload)
	} else if c.RequestType == "PUT" {
		req, err = PutBuilder(c.Url, c.Payload)
	} else if c.RequestType == "DELETE" {
		req, err = DeleteBuilder(c.Url)
	}
	if err != nil {
		return "", err
	}
	if req == nil {
		return "", status.Error(codes.Internal, "request creation failed")
	}

	var bearer = "Bearer " + c.BearerToken
	if c.BearerToken == "" {
		req.SetBasicAuth(c.Username, c.Password)
	} else {
		req.Header.Set("Authorization", bearer)
	}
	resp, err := Client.Do(req)
	if err != nil {
		err = status.Error(codes.NotFound, err.Error())
		fmt.Println("error2:", err)
		return "", err
	}

	if resp != nil && (resp.StatusCode < 200 || resp.StatusCode > 299) {
		fmt.Println("Error3:", resp.Status)
		return "", err
	}
	if resp == nil {
		fmt.Println("Error4:")
		return "", err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	return bodyString, err
}

func GetBuilder(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err = status.Error(codes.Unknown, err.Error())
		fmt.Println("error:", err)
		return nil, err
	}
	return req, nil
}

func PostBuilder(url string, body string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		err = status.Error(codes.Unknown, err.Error())
		fmt.Println("error:", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func PutBuilder(url string, body string) (*http.Request, error) {
	req, err := http.NewRequest("PUT", url, strings.NewReader(body))
	if err != nil {
		err = status.Error(codes.Unknown, err.Error())
		fmt.Println("error:", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func DeleteBuilder(url string) (*http.Request, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		err = status.Error(codes.Unknown, err.Error())
		fmt.Println("error:", err)
		return nil, err
	}
	return req, nil
}
