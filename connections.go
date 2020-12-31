package guacamole

import (
	"fmt"
	"net/http"

	"github.com/techBeck03/guacamole-api-client/types"
)

const (
	connectionsBasePath = "connections"
)

// CreateConnection creates a guacamole connection
func (c *Client) CreateConnection(connection *types.GuacConnection) (types.GuacConnection, error) {
	var ret types.GuacConnection
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s/%s", c.baseURL, connectionsBasePath), connection)

	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

// ReadConnection gets a connection by identifier
func (c *Client) ReadConnection(identifier string) (types.GuacConnection, error) {
	var ret types.GuacConnection
	var retParams types.GuacConnectionParameters

	// Get connection base details
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", c.baseURL, connectionsBasePath, identifier), nil)

	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}

	// Get connection parameters
	request, err = c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s/parameters", c.baseURL, connectionsBasePath, identifier), nil)

	if err != nil {
		return ret, err
	}

	err = c.Call(request, &retParams)
	if err != nil {
		return ret, err
	}

	ret.Properties = retParams

	return ret, nil
}

// UpdateConnection updates a connection by identifier
func (c *Client) UpdateConnection(connection *types.GuacConnection) error {
	request, err := c.CreateJSONRequest(http.MethodPut, fmt.Sprintf("%s/%s/%s", c.baseURL, connectionsBasePath, connection.Identifier), connection)

	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}
	return nil
}

// DeleteConnection deletes a connection by identifier
func (c *Client) DeleteConnection(identifier string) error {
	request, err := c.CreateJSONRequest(http.MethodDelete, fmt.Sprintf("%s/%s/%s", c.baseURL, connectionsBasePath, identifier), nil)

	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListConnections lists all connections
func (c *Client) ListConnections() ([]types.GuacConnection, error) {
	var ret []types.GuacConnection
	connectionTree, err := c.GetConnectionTree()

	if err != nil {
		return ret, err
	}

	flattenedConnections, _, err := flatten([]types.GuacConnectionGroup{connectionTree})

	ret = flattenedConnections
	return ret, nil
}
