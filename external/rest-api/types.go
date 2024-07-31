package rest_api

import (
	"context"
)

type Config struct {
	context.Context

	//URL to trigger
	Url string `json:"url,omitempty"`
	//Bearer token
	BearerToken string `json:"bearer_token,omitempty"`
	//Username (used for basic auth if no bearer token)
	Username string `json:"username,omitempty"`
	//Password (used for basic auth if no bearer token)
	Password string `json:"password,omitempty"`
	//Request type
	RequestType string `json:"request_type,omitempty"`
	//Payload for request
	Payload string `json:"payload,omitempty"`
	//Expected response code
	ExpectedResponseCode string `json:"expected_response_code,omitempty"`
}
