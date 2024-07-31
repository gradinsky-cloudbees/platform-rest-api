package rest

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func ExecuteApiCall(requestType string, url string, payload string, bearerToken string, username string, password string, expectedResponseCode string) (*http.Response, error) {
	var req *http.Request
	var err error
	//Build request
	if requestType == "GET" {
		req, err = GetBuilder(url)
	} else if requestType == "POST" {
		req, err = PostBuilder(url, payload)
	} else if requestType == "PUT" {
		req, err = PutBuilder(url, payload)
	} else if requestType == "DELETE" {
		req, err = DeleteBuilder(url)
	}
	// If errors building the request, error handle
	if err != nil {
		return nil, err
	}
	if req == nil {
		return nil, status.Error(codes.Internal, "request creation failed")
	}

	//Auth
	if bearerToken != "" || username != "" {
		if bearerToken == "" {
			req.SetBasicAuth(username, password)
		} else {
			bearer := "Bearer " + bearerToken
			req.Header.Set("Authorization", bearer)
		}
	}

	//Perform API call
	resp, err := Client.Do(req)

	//Error handling time
	if err != nil {
		err = status.Error(codes.NotFound, err.Error())
		log.Println("Error occurred:", err)
		return resp, status.Error(codes.Unknown, fmt.Sprintf("ERROR - %v", err.Error()))
	}

	if expectedResponseCode == "" && resp != nil && (resp.StatusCode < 200 || resp.StatusCode > 299) {
		log.Println("Response code was not 200-299:", resp.Status)
		return resp, MapHttpToGrpcErrorCode(resp)
	}
	if resp == nil {
		log.Println("Response was empty")
		return nil, status.Error(codes.NotFound, "Response was empty")
	}
	expectedCode, _ := strconv.Atoi(expectedResponseCode)
	if expectedResponseCode != "" && resp.StatusCode != expectedCode {
		return resp, status.Error(codes.OutOfRange, fmt.Sprintf("Expected status code (%v) does not match returned status code (%v)", resp.StatusCode, expectedCode))
	}

	return resp, err
}

func GetBuilder(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err = status.Error(codes.Unknown, err.Error())
		log.Println("error:", err)
		return nil, err
	}
	return req, nil
}

func PostBuilder(url string, body string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		err = status.Error(codes.Unknown, err.Error())
		log.Println("error:", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func PutBuilder(url string, body string) (*http.Request, error) {
	req, err := http.NewRequest("PUT", url, strings.NewReader(body))
	if err != nil {
		err = status.Error(codes.Unknown, err.Error())
		log.Println("error:", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func DeleteBuilder(url string) (*http.Request, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		err = status.Error(codes.Unknown, err.Error())
		log.Println("error:", err)
		return nil, err
	}
	return req, nil
}
