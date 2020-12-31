package guacamole

import (
	"fmt"
	"net/http"

	"github.com/techBeck03/guacamole-api-client/types"
)

const (
	userGroupsBasePath           = "userGroups"
	userGroupPermissionsBasePath = "/connectionGroupPermissions/"
)

// CreateUserGroup creates a guacamole user group
func (c *Client) CreateUserGroup(userGroup *types.GuacUserGroup) (types.GuacUserGroup, error) {
	var ret types.GuacUserGroup
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s/%s", c.baseURL, userGroupsBasePath), userGroup)

	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

// ReadUserGroup gets a user group by name
func (c *Client) ReadUserGroup(name string) (types.GuacUserGroup, error) {
	var ret types.GuacUserGroup

	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", c.baseURL, userGroupsBasePath, name), nil)

	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

// UpdateUserGroup updates a user group by username
func (c *Client) UpdateUserGroup(group *types.GuacUserGroup) error {
	request, err := c.CreateJSONRequest(http.MethodPut, fmt.Sprintf("%s/%s/%s", c.baseURL, userGroupsBasePath, group.Identifier), group)

	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserGroup deletes a user group by name
func (c *Client) DeleteUserGroup(name string) error {
	request, err := c.CreateJSONRequest(http.MethodDelete, fmt.Sprintf("%s/%s/%s", c.baseURL, userGroupsBasePath, name), nil)

	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListUserGroups lists all user groups
func (c *Client) ListUserGroups() ([]types.GuacUserGroup, error) {
	var ret map[string]types.GuacUserGroup
	var userGroupList []types.GuacUserGroup

	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s/%s", c.baseURL, userGroupsBasePath), nil)

	if err != nil {
		return userGroupList, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return userGroupList, err
	}

	for i := range ret {
		userGroupList = append(userGroupList, ret[i])
	}

	return userGroupList, nil
}

// AddUserGroupConnectionPermissions adds connection permissions to a user group
func (c *Client) AddUserGroupConnectionPermissions(group string, identifiers []string) error {
	var permissionItems []types.GuacPermissionItem

	for identifer := range identifiers {
		permissionItems = append(permissionItems, types.GuacPermissionItem{
			Op:    "add",
			Path:  fmt.Sprintf("%s/%s", userGroupPermissionsBasePath, identifiers[identifer]),
			Value: "READ",
		})
	}
	request, err := c.CreateJSONRequest(http.MethodPatch, fmt.Sprintf("%s/%s/%s/permissions", c.baseURL, userGroupsBasePath, group), permissionItems)

	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}
	return nil
}

// RemoveUserGroupConnectionPermissions removes connection permissions from a user group
func (c *Client) RemoveUserGroupConnectionPermissions(group string, identifiers []string) error {
	var permissionItems []types.GuacPermissionItem

	for identifer := range identifiers {
		permissionItems = append(permissionItems, types.GuacPermissionItem{
			Op:    "remove",
			Path:  fmt.Sprintf("%s/%s", userGroupPermissionsBasePath, identifiers[identifer]),
			Value: "READ",
		})
	}
	request, err := c.CreateJSONRequest(http.MethodPatch, fmt.Sprintf("%s/%s/%s/permissions", c.baseURL, userGroupsBasePath, group), permissionItems)

	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}
	return nil
}
