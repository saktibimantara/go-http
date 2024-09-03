package go_http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// Constant Header
const (
	Authorization = "Authorization"
	ContentType   = "Content-Type"
	Accept        = "Accept"
	Basic         = "Basic"
	Bearer        = "Bearer"
)

// Header custom type
type Header map[string]string

func NewBearerToken(token string) Header {
	return Header{
		Authorization: Bearer + " " + token,
	}
}

// Body custom type
type Body map[string]any

type Config struct {
	// BaseURL is the base URL for the API
	BaseURL string

	// It is used to authenticate the request using parameters in the URL
	APIKey string

	// It is used to authenticate the request using the Authorization header
	Header Header
}

type CallAPI interface {
	Get(url string) (*Response, error)
	Post(url string, body Body) (*Response, error)
	Put(url string, body Body) (*Response, error)
	Delete(url string) (*Response, error)
}

// Response callx model
type Response struct {
	Code   int
	Data   []byte
	Header http.Header
}

type GoHTTP struct {
	httpClient *http.Client
	config     *Config
}

func New(c *Config) *GoHTTP {
	httpClient := &http.Client{}
	return &GoHTTP{httpClient: httpClient, config: c}
}

func (g *GoHTTP) Get(url string) (*Response, error) {
	dataURL := url
	if g.config.BaseURL != "" {
		dataURL = g.config.BaseURL + url
	}

	if g.config.APIKey != "" {
		if strings.Contains(dataURL, "?") {
			dataURL += "&apikey=" + g.config.APIKey
		} else {
			dataURL += "?apikey=" + g.config.APIKey
		}
	}

	return g.Request(dataURL, http.MethodGet, nil)
}

func (g *GoHTTP) Post(url string, body Body) (*Response, error) {
	return g.Request(url, http.MethodPost, &body)
}

func (g *GoHTTP) Put(url string, body Body) (*Response, error) {
	return g.Request(url, http.MethodPut, &body)
}

func (g *GoHTTP) Delete(url string) (*Response, error) {
	return g.Request(url, http.MethodDelete, nil)
}

func (g *GoHTTP) Request(url string, method string, body *Body) (*Response, error) {

	req, err := CreateRequest(url, method, body)
	if err != nil {
		return nil, err
	}

	for key, value := range g.config.Header {
		req.Header.Set(key, value)
	}

	resp, err := g.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		Code:   resp.StatusCode,
		Data:   data,
		Header: resp.Header,
	}, nil
}

func CreateRequest(url string, method string, body *Body) (*http.Request, error) {
	var payload *bytes.Buffer

	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		payload = bytes.NewBuffer(jsonBody)
	}

	if payload == nil {
		payload = bytes.NewBuffer([]byte{})
	}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	return req, nil
}
