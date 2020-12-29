package guacamole

import (
	"fmt"
	"net/http"

	"github.com/techBeck03/guacamole-api-client/types"
)

const (
	usersBasePath = "users"
)

// CreateUser creates a guacamole user
func (c *Client) CreateUser(user *types.GuacUser) (types.GuacUser, error) {
	var ret types.GuacUser
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s/%s", c.baseURL, usersBasePath), user)

	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

// ReadUser gets a user by username
func (c *Client) ReadUser(username string) (types.GuacUser, error) {
	var ret types.GuacUser

	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", c.baseURL, usersBasePath, username), nil)

	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

// UpdateUser updates a user by username
func (c *Client) UpdateUser(user *types.GuacUser) error {
	request, err := c.CreateJSONRequest(http.MethodPut, fmt.Sprintf("%s/%s/%s", c.baseURL, usersBasePath, user.Username), user)

	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user by username
func (c *Client) DeleteUser(username string) error {
	request, err := c.CreateJSONRequest(http.MethodDelete, fmt.Sprintf("%s/%s/%s", c.baseURL, usersBasePath, username), nil)

	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListUsers lists all users
func (c *Client) ListUsers() ([]types.GuacUser, error) {
	var ret map[string]types.GuacUser
	var userList []types.GuacUser

	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s", c.baseURL, usersBasePath), nil)

	if err != nil {
		return userList, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return userList, err
	}

	for i := range ret {
		userList = append(userList, ret[i])
	}

	return userList, nil
}
