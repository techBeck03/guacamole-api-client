package guacamole

import (
	"fmt"
	"github.com/techBeck03/guacamole-api-client/types"
	"net/http"
	"net/url"
)

const (
	sharingProfilesBasePath = "sharingProfiles"
)

// CreateSharingProfiles creates a guacamole SharingProfile
func (c *Client) CreateSharingProfiles(sharingProfiles *types.GuacSharingProfile) error {
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s/%s", c.baseURL, sharingProfilesBasePath), sharingProfiles)

	if err != nil {
		return err
	}

	err = c.Call(request, &sharingProfiles)
	if err != nil {
		return err
	}
	return nil
}

// ReadSharingProfile gets a sharingProfile by identifier
func (c *Client) ReadSharingProfile(identifier string) (types.GuacSharingProfile, error) {
	var ret types.GuacSharingProfile

	// Get connection base details
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", c.baseURL, sharingProfilesBasePath, url.QueryEscape(identifier)), nil)

	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// DeleteSharingProfile deletes a sharingProfile by identifier
func (c *Client) DeleteSharingProfile(identifier string) error {
	request, err := c.CreateJSONRequest(http.MethodDelete, fmt.Sprintf("%s/%s/%s", c.baseURL, sharingProfilesBasePath, url.QueryEscape(identifier)), nil)

	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListSharingProfile lists all sharingProfile
func (c *Client) ListSharingProfiles() ([]types.GuacSharingProfile, error) {
	var ret []types.GuacSharingProfile
	var connectionList map[string]types.GuacSharingProfile

	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s", c.baseURL, sharingProfilesBasePath), nil)

	if err != nil {
		return ret, err
	}

	err = c.Call(request, &connectionList)
	if err != nil {
		return ret, err
	}

	for _, connection := range connectionList {
		ret = append(ret, connection)
	}
	return ret, nil
}
