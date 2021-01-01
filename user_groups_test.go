// +build all unittests

package guacamole

import (
	"fmt"
	"os"
	"testing"

	"github.com/techBeck03/guacamole-api-client/types"
)

var (
	userGroupsConfig = Config{
		URI:                    os.Getenv("GUACAMOLE_URL"),
		Username:               os.Getenv("GUACAMOLE_USERNAME"),
		Password:               os.Getenv("GUACAMOLE_PASSWORD"),
		DisableTLSVerification: true,
	}
	testGroup                    = types.GuacUserGroup{Identifier: "(testing) Test Group"}
	testUserGroupConnectionGroup = types.GuacConnectionGroup{
		Name:             "Testing User Groups Group",
		ParentIdentifier: "ROOT",
		Type:             "ORGANIZATIONAL",
	}
	testUserGroupConnection = types.GuacConnection{
		Name:     "Testing User Connection",
		Protocol: "ssh",
	}
	testUserGroupPermissionItems = []types.GuacPermissionItem{}
)

func TestListUserGroups(t *testing.T) {
	client := New(userGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, userGroupsConfig)
	}

	_, err = client.ListUserGroups()
	if err != nil {
		t.Errorf("Error %s listing user groups with client %+v", err, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestCreateUserGroup(t *testing.T) {
	client := New(userGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, userGroupsConfig)
	}

	_, err = client.CreateUserGroup(&testGroup)
	if err != nil {
		t.Errorf("Error %s creating user group with identifier: %s with client %+v", err, testGroup.Identifier, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestReadUserGroup(t *testing.T) {
	client := New(userGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, userGroupsConfig)
	}

	group, err := client.ReadUserGroup(testGroup.Identifier)
	if err != nil {
		t.Errorf("Error %s reading user group with identifier: %s with client %+v", err, testGroup.Identifier, client)
	}

	if group.Identifier != testGroup.Identifier {
		t.Errorf("Expected group identifier = %s but got value = %s", testGroup.Identifier, group.Identifier)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestUpdateUserGroup(t *testing.T) {
	client := New(userGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, userGroupsConfig)
	}

	group, err := client.ReadUserGroup(testGroup.Identifier)
	if err != nil {
		t.Errorf("Error %s reading user group with identifier: %s with client %+v", err, testGroup.Identifier, client)
	}

	group.Attributes.Disabled = "true"

	err = client.UpdateUserGroup(&group)

	if err != nil {
		t.Errorf("Error %s updating user group with identifier: %s with client %+v", err, testGroup.Identifier, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestGetUserGroupPermissions(t *testing.T) {
	client := New(userGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, userGroupsConfig)
	}

	_, err = client.GetUserGroupPermissions(testGroup.Identifier)
	if err != nil {
		t.Errorf("Error %s reading user group permissions for: %s with client %+v", err, testGroup.Identifier, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestAddUserGroupConnectionPermissions(t *testing.T) {
	client := New(userGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, userGroupsConfig)
	}

	group, err := client.CreateConnectionGroup(&testUserGroupConnectionGroup)

	testUserGroupConnectionGroup = group

	if err != nil {
		t.Errorf("Error %s creating connection group: %s with client %+v", err, testUserGroupConnectionGroup.Name, client)
	}

	testUserGroupConnection.ParentIdentifier = group.Identifier

	connection, err := client.CreateConnection(&testUserGroupConnection)
	if err != nil {
		t.Errorf("Error %s creating user connection: %s with client %+v", err, testUserGroupConnection.Identifier, client)
	}

	testUserGroupConnection = connection

	testUserGroupPermissionItems = []types.GuacPermissionItem{
		{
			Op:    "add",
			Path:  fmt.Sprintf("%s/%s", connectionGroupPermissionsBasePath, testUserGroupConnectionGroup.Identifier),
			Value: "READ",
		},
		{
			Op:    "add",
			Path:  fmt.Sprintf("%s/%s", connectionPermissionsBasePath, testUserGroupConnection.Identifier),
			Value: "READ",
		},
	}

	err = client.SetUserGroupConnectionPermissions(testGroup.Identifier, &testUserGroupPermissionItems)

	if err != nil {
		t.Errorf("Error %s adding user connection permissions for group: %s with client %+v", err, testGroup.Identifier, client)
	}

	permissions, err := client.GetUserGroupPermissions(testGroup.Identifier)
	if err != nil {
		t.Errorf("Error %s reading group permissions for: %s with client %+v", err, testGroup.Identifier, client)
	}

	_, ok := permissions.ConnectionGroupPermissions[group.Identifier]

	if !ok {
		t.Errorf("Error adding connection group: %s permissions for: %s with client %+v", group.Identifier, testGroup.Identifier, client)
	}

	_, ok = permissions.ConnectionPermissions[connection.Identifier]

	if !ok {
		t.Errorf("Error adding connection: %s permissions for: %s with client %+v", connection.Identifier, testGroup.Identifier, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestRemoveUserGroupConnectionPermissions(t *testing.T) {
	client := New(userGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, userGroupsConfig)
	}

	testUserGroupPermissionItems = []types.GuacPermissionItem{
		{
			Op:    "remove",
			Path:  fmt.Sprintf("%s/%s", connectionGroupPermissionsBasePath, testUserGroupConnectionGroup.Identifier),
			Value: "READ",
		},
		{
			Op:    "remove",
			Path:  fmt.Sprintf("%s/%s", connectionPermissionsBasePath, testUserGroupConnection.Identifier),
			Value: "READ",
		},
	}

	err = client.SetUserGroupConnectionPermissions(testGroup.Identifier, &testUserGroupPermissionItems)

	if err != nil {
		t.Errorf("Error %s adding connection permissions for group: %s with client %+v", err, testGroup.Identifier, client)
	}

	permissions, err := client.GetUserGroupPermissions(testGroup.Identifier)
	if err != nil {
		t.Errorf("Error %s reading group permissions for: %s with client %+v", err, testGroup.Identifier, client)
	}

	_, ok := permissions.ConnectionGroupPermissions[testUserGroupConnectionGroup.Identifier]

	if ok {
		t.Errorf("Error %s deleting connection group permissions for: %s with client %+v", err, testGroup.Identifier, client)
	}

	_, ok = permissions.ConnectionPermissions[testUserGroupConnection.Identifier]

	if ok {
		t.Errorf("Error %s deleting connection permissions for: %s with client %+v", err, testGroup.Identifier, client)
	}

	err = client.DeleteConnection(testUserGroupConnection.Identifier)

	if err != nil {
		t.Errorf("Error %s deleting connection: %s with client %+v", err, testUserGroupConnection.Identifier, client)
	}

	err = client.DeleteConnectionGroup(testUserGroupConnectionGroup.Identifier)

	if err != nil {
		t.Errorf("Error %s deleting connection group: %s with client %+v", err, testUserGroupConnectionGroup.Identifier, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestDeleteUserGroup(t *testing.T) {
	client := New(userGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, userGroupsConfig)
	}

	err = client.DeleteUserGroup(testGroup.Identifier)
	if err != nil {
		t.Errorf("Error %s deleting user group with identifier: %s with client %+v", err, testGroup.Identifier, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}
