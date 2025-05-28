package types

import (
	"fmt"
	"net/http"

	"github.com/aamgayle/gotmdb/consts"
)

type Config struct {
	APIKey      string
	BearerToken string
	SessionID   string
}

type Method int

const (
	GET Method = iota + 1
	POST
	DELETE
)

type TMDBReqProps struct {
	BaseURL string
	URL     string
	Method  Method
	Config  Config
}

func NewHTTPRequestFromProps(props *TMDBReqProps) (*http.Request, error) {
	method := ""

	switch props.Method {
	case GET:
		method = "GET"
	case POST:
		method = "POST"
	case DELETE:
		method = "DELETE"
	}

	req, err := http.NewRequest(method, props.BaseURL+props.URL, nil)
	if err != nil {
		fmt.Printf("NewHTTPRequestFromProps Error: %s", err.Error())
		return nil, err
	}

	return req, nil
}

func NewTMDBRequest(opts ...func(*TMDBReqProps)) (*http.Request, error) {
	props := &TMDBReqProps{
		BaseURL: consts.BaseURL,
		URL:     "",
		Method:  GET,
		Config:  Config{},
	}

	for _, opt := range opts {
		opt(props)
	}

	req, err := NewHTTPRequestFromProps(props)
	if err != nil {
		fmt.Printf("NewTMDBRequest Error: %s", err.Error())
		return nil, err
	}

	return req, nil
}
