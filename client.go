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
	URL                    string
	Password               string
	Username               string
	DisableTLSVerification bool
	DisableCookies         bool
}

// Client - base client for guacamole interactions
type Client struct {
	client  *http.Client
	config  Config
	baseURL string
	token   string
	cookies []*http.Cookie
}

// New - creates a new guacamole client
func New(config Config) Client {
	var client *http.Client
	if config.DisableTLSVerification {
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{Transport: transport}
	} else {
		client = http.DefaultClient
	}
	return Client{
		client: client,
		config: config,
	}
}

// Connect - function for establishing connection to guacamole
func (c *Client) Connect() error {
	resp, err := c.client.PostForm(fmt.Sprintf("%s/%s", c.config.URL, tokenPath),
		url.Values{
			"username": {c.config.Username},
			"password": {c.config.Password},
		})
	if err != nil {
		return err
	}
	if resp.StatusCode == 403 {
		return fmt.Errorf("Invalid Credentials")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var tokenresp types.AuthenticationResponse

	err = json.Unmarshal(body, &tokenresp)
	if err != nil {
		return err
	}
	c.token = tokenresp.AuthToken
	c.baseURL = fmt.Sprintf("%s/api/session/data/%s", c.config.URL, tokenresp.DataSource)
	if !(c.config.DisableCookies) {
		c.cookies = resp.Cookies()
	}
	return nil
}

// Disconnect deletes a user session token
func (c *Client) Disconnect() error {

	request, err := c.CreateJSONRequest(http.MethodDelete, fmt.Sprintf("%s/%s/%s", c.config.URL, tokenPath, c.token), nil)
	if err != nil {
		return err
	}
	err = c.Call(request, nil)
	return err
}

// CreateJSONRequest - helper function for creating json based http requests
func (c *Client) CreateJSONRequest(method string, path string, params interface{}) (*http.Request, error) {
	var request *http.Request
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(&params)
	if err != nil {
		return request, err
	}
	request, err = http.NewRequest(method, path, &buf)
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
	// Add authentication token to query params
	q := request.URL.Query()
	q.Add("token", c.token)

	request.URL.RawQuery = q.Encode()

	// Add cookies if configured
	if !(c.config.DisableCookies) {
		for i := range c.cookies {
			request.AddCookie(c.cookies[i])
		}
	}

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

func (c *Client) Token() (token string) {
	return c.token
}
