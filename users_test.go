// +build all unittests

package guacamole

import (
	"fmt"
	"os"
	"testing"

	"github.com/techBeck03/guacamole-api-client/types"
)

var (
	usersConfig = Config{
		URI:                    os.Getenv("GUACAMOLE_URL"),
		Username:               os.Getenv("GUACAMOLE_USERNAME"),
		Password:               os.Getenv("GUACAMOLE_PASSWORD"),
		DisableTLSVerification: true,
	}
	testUser                = types.GuacUser{Username: "testing"}
	testUserConnectionGroup = types.GuacConnectionGroup{
		Name:             "Testing Users Group",
		ParentIdentifier: "ROOT",
		Type:             "ORGANIZATIONAL",
	}
	testUserConnection = types.GuacConnection{
		Name:     "Testing User Connection",
		Protocol: "ssh",
	}
	testUserPermissionItems = []types.GuacPermissionItem{}
)

func TestListUsers(t *testing.T) {
	client := New(usersConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, usersConfig)
	}

	_, err = client.ListUsers()
	if err != nil {
		t.Errorf("Error %s listing users with client %+v", err, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestCreateUser(t *testing.T) {
	client := New(usersConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, usersConfig)
	}

	_, err = client.CreateUser(&testUser)
	if err != nil {
		t.Errorf("Error %s creating user: %s with client %+v", err, testUser.Username, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestReadUser(t *testing.T) {
	client := New(usersConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, usersConfig)
	}

	user, err := client.ReadUser(testUser.Username)
	if err != nil {
		t.Errorf("Error %s reading user: %s with client %+v", err, testUser.Username, client)
	}

	if user.Username != testUser.Username {
		t.Errorf("Expected username = %s read username = %s", testUser.Username, user.Username)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestUpdateUser(t *testing.T) {
	client := New(usersConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, usersConfig)
	}

	user, err := client.ReadUser(testUser.Username)
	if err != nil {
		t.Errorf("Error %s reading user: %s with client %+v", err, testUser.Username, client)
	}

	user.Attributes.GuacFullName = "Go Testing User"

	err = client.UpdateUser(&user)

	if err != nil {
		t.Errorf("Error %s updating user: %s with client %+v", err, testUser.Username, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestGetUserPermissions(t *testing.T) {
	client := New(usersConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, usersConfig)
	}

	_, err = client.GetUserPermissions(testUser.Username)
	if err != nil {
		t.Errorf("Error %s reading user permissions for: %s with client %+v", err, testUser.Username, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestAddUserConnectionPermissions(t *testing.T) {
	client := New(usersConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, usersConfig)
	}

	group, err := client.CreateConnectionGroup((&testUserConnectionGroup))

	testUserConnectionGroup = group

	if err != nil {
		t.Errorf("Error %s creating connection group: %s with client %+v", err, testUserConnectionGroup.Name, client)
	}

	testUserConnection.ParentIdentifier = group.Identifier

	connection, err := client.CreateConnection(&testUserConnection)
	if err != nil {
		t.Errorf("Error %s creating user connection: %s with client %+v", err, testUserConnection.Identifier, client)
	}

	testUserConnection = connection

	testUserPermissionItems = []types.GuacPermissionItem{
		{
			Op:    "add",
			Path:  fmt.Sprintf("%s/%s", connectionGroupPermissionsBasePath, testUserConnectionGroup.Identifier),
			Value: "READ",
		},
		{
			Op:    "add",
			Path:  fmt.Sprintf("%s/%s", connectionPermissionsBasePath, testUserConnection.Identifier),
			Value: "READ",
		},
	}
	err = client.SetUserConnectionPermissions(testUser.Username, &testUserPermissionItems)

	if err != nil {
		t.Errorf("Error %s adding user connection permissions for user: %s with client %+v", err, testUser.Username, client)
	}

	permissions, err := client.GetUserPermissions(testUser.Username)
	if err != nil {
		t.Errorf("Error %s reading user permissions for: %s with client %+v", err, testUser.Username, client)
	}

	_, ok := permissions.ConnectionGroupPermissions[group.Identifier]

	if !ok {
		t.Errorf("Error adding connection group: %s permissions for: %s with client %+v", group.Identifier, testUser.Username, client)
	}

	_, ok = permissions.ConnectionPermissions[connection.Identifier]

	if !ok {
		t.Errorf("Error adding connection: %s permissions for: %s with client %+v", connection.Identifier, testUser.Username, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestRemoveUserConnectionPermissions(t *testing.T) {
	client := New(usersConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, usersConfig)
	}

	testUserPermissionItems = []types.GuacPermissionItem{
		{
			Op:    "remove",
			Path:  fmt.Sprintf("%s/%s", connectionGroupPermissionsBasePath, testUserConnectionGroup.Identifier),
			Value: "READ",
		},
		{
			Op:    "remove",
			Path:  fmt.Sprintf("%s/%s", connectionPermissionsBasePath, testUserConnection.Identifier),
			Value: "READ",
		},
	}

	err = client.SetUserConnectionPermissions(testUser.Username, &testUserPermissionItems)

	if err != nil {
		t.Errorf("Error %s adding user connection permissions for user: %s with client %+v", err, testUser.Username, client)
	}

	permissions, err := client.GetUserPermissions(testUser.Username)
	if err != nil {
		t.Errorf("Error %s reading user permissions for: %s with client %+v", err, testUser.Username, client)
	}

	_, ok := permissions.ConnectionGroupPermissions[testUserConnectionGroup.Identifier]

	if ok {
		t.Errorf("Error %s deleting connection group permissions for: %s with client %+v", err, testUser.Username, client)
	}

	_, ok = permissions.ConnectionPermissions[testUserConnection.Identifier]

	if ok {
		t.Errorf("Error %s deleting connection permissions for: %s with client %+v", err, testUser.Username, client)
	}

	err = client.DeleteConnection(testUserConnection.Identifier)

	if err != nil {
		t.Errorf("Error %s deleting connection: %s with client %+v", err, testUserConnection.Identifier, client)
	}

	err = client.DeleteConnectionGroup(testUserConnectionGroup.Identifier)

	if err != nil {
		t.Errorf("Error %s deleting connection group: %s with client %+v", err, testUserConnectionGroup.Identifier, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestDeleteUser(t *testing.T) {
	client := New(usersConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, usersConfig)
	}

	err = client.DeleteUser(testUser.Username)
	if err != nil {
		t.Errorf("Error %s deleting user: %s with client %+v", err, testUser.Username, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}
