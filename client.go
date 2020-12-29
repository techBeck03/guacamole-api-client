package guacamole

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/techBeck03/guacamole-api-client/types"
)

const (
	tokenPath string = "api/tokens"
)

// Config - Configuration details for connecting to guacamole
type Config struct {
	URI                    string
	Password               string
	Username               string
	DisableTLSVerification bool
}

// Client - base client for guacamole interactions
type Client struct {
	client *http.Client
	config Config
	token  string
}

// Connect - function for establishing connection to guacamole
func (c *Client) Connect() error {
	if c.config.DisableTLSVerification {
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		c.client = &http.Client{Transport: transport}
	} else {
		c.client = http.DefaultClient
	}
	resp, err := c.client.PostForm(fmt.Sprintf("%s/api/tokens", c.config.URI),
		url.Values{
			"username": {c.config.Username},
			"password": {c.config.Password},
		})
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var tokenresp types.ConnectResponse

	err = json.Unmarshal(body, &tokenresp)
	if err != nil {
		return err
	}
	c.token = tokenresp.AuthToken
	return nil
}

// RefreshToken - function for refreshing login token
func (c *Client) RefreshToken() error {
	resp, err := c.client.PostForm(fmt.Sprintf("%s/%s", c.config.URI, tokenPath),
		url.Values{
			"token": {c.token},
		})
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var tokenresp types.ConnectResponse

	err = json.Unmarshal(body, &tokenresp)
	if err != nil {
		return err
	}
	c.token = tokenresp.AuthToken
	return nil
}

// CreateJSONRequest - helper function for creating json based http requests
func (c *Client) CreateJSONRequest(method string, path string, params interface{}) (*http.Request, error) {
	var request *http.Request
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(&params)
	if err != nil {
		return request, err
	}
	request, err = http.NewRequest(method, fmt.Sprintf("%s/%s", c.config.URI, path), &buf)
	if err != nil {
		return request, err
	}
	if params == nil {
		request.Body = http.NoBody
	}
	request.Header.Set("Content-Type", "application/json")
	return request, nil
}

// Call - function for handling http requests
func (c *Client) Call(request *http.Request, result interface{}) error {
	err := c.RefreshToken()
	if err != nil {
		return err
	}

	// URL query params
	q := request.URL.Query()
	q.Add("token", c.token)

	request.URL.RawQuery = q.Encode()

	response, err := c.client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	if !(response.StatusCode >= 200 && response.StatusCode <= 299) {
		var rawBodyBuffer bytes.Buffer
		// Decode raw response, usually contains
		// additional error details
		body := io.TeeReader(response.Body, &rawBodyBuffer)
		var responseBody interface{}
		json.NewDecoder(body).Decode(&responseBody)
		return fmt.Errorf("Request %+v\n failed with status code %d\n response %+v\n%+v", request,
			response.StatusCode, responseBody,
			response)
	}
	// If no result is expected, don't attempt to decode a potentially
	// empty response stream and avoid incurring EOF errors
	if result == nil {
		return nil
	}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return err
	}
	return nil
}
