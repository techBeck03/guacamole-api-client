//go:build all || unittests
// +build all unittests

package guacamole

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/techBeck03/guacamole-api-client/types"
)

var (
	userGroupsConfig = Config{
		URL:                    os.Getenv("GUACAMOLE_URL"),
		Username:               os.Getenv("GUACAMOLE_USERNAME"),
		Password:               os.Getenv("GUACAMOLE_PASSWORD"),
		DisableTLSVerification: true,
	}
	testGroup                    = types.GuacUserGroup{Identifier: "(testing) Test Group"}
	testGroupMemberGroup         = types.GuacUserGroup{Identifier: "(testing) Test Child Group"}
	testGroupParentGroup         = types.GuacUserGroup{Identifier: "(testing) Test Parent Group"}
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
	testUserGroupUser            = types.GuacUser{Username: "testingUserGroups"}
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

	err = client.CreateUserGroup(&testGroup)
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

	err = client.CreateConnectionGroup(&testUserGroupConnectionGroup)

	if err != nil {
		t.Errorf("Error %s creating connection group: %s with client %+v", err, testUserGroupConnectionGroup.Name, client)
	}

	testUserGroupConnection.ParentIdentifier = testUserGroupConnectionGroup.Identifier

	err = client.CreateConnection(&testUserGroupConnection)
	if err != nil {
		t.Errorf("Error %s creating user connection: %s with client %+v", err, testUserGroupConnection.Identifier, client)
	}

	testUserGroupPermissionItems = []types.GuacPermissionItem{
		{
			Op:    "add",
			Path:  fmt.Sprintf("%s/%s", ConnectionGroupPermissionsBasePath, testUserGroupConnectionGroup.Identifier),
			Value: "READ",
		},
		{
			Op:    "add",
			Path:  fmt.Sprintf("%s/%s", ConnectionPermissionsBasePath, testUserGroupConnection.Identifier),
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

	_, ok := permissions.ConnectionGroupPermissions[testUserGroupConnectionGroup.Identifier]

	if !ok {
		t.Errorf("Error adding connection group: %s permissions for: %s with client %+v", testUserGroupConnectionGroup.Identifier, testGroup.Identifier, client)
	}

	_, ok = permissions.ConnectionPermissions[testUserGroupConnection.Identifier]

	if !ok {
		t.Errorf("Error adding connection: %s permissions for: %s with client %+v", testUserGroupConnection.Identifier, testGroup.Identifier, client)
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
			Path:  fmt.Sprintf("%s/%s", ConnectionGroupPermissionsBasePath, testUserGroupConnectionGroup.Identifier),
			Value: "READ",
		},
		{
			Op:    "remove",
			Path:  fmt.Sprintf("%s/%s", ConnectionPermissionsBasePath, testUserGroupConnection.Identifier),
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

func TestSetUserGroupMemberGroups(t *testing.T) {
	client := New(userGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, userGroupsConfig)
	}

	// Create child group
	err = client.CreateUserGroup(&testGroupMemberGroup)

	if err != nil {
		t.Errorf("Error creating user group: %s with client %+v\n", testGroupMemberGroup.Identifier, client)
	}

	permissionItems := []types.GuacPermissionItem{
		client.NewAddGroupMemberPermission(testGroupMemberGroup.Identifier),
	}

	err = client.SetUserGroupMemberGroups(testGroup.Identifier, &permissionItems)

	if err != nil {
		t.Errorf("Error adding group: %s to user group: %s with client %+v\n", testGroupMemberGroup.Identifier, testGroup.Identifier, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestGetUserGroupMemberGroups(t *testing.T) {
	client := New(userGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, userGroupsConfig)
	}

	// Get group member groups
	members, err := client.GetUserGroupMemberGroups(testGroup.Identifier)

	if members[0] != testGroupMemberGroup.Identifier {
		t.Errorf("Wrong member group found: %s expected %s", members[0], testGroupMemberGroup.Identifier)
	}

	// Remove group membership
	permissionItems := []types.GuacPermissionItem{
		client.NewRemoveGroupMemberPermission(testGroupMemberGroup.Identifier),
	}

	err = client.SetUserGroupMemberGroups(testGroup.Identifier, &permissionItems)

	if err != nil {
		t.Errorf("Error removing group: %s from user group: %s with client %+v\n", testGroupMemberGroup.Identifier, testGroup.Identifier, client)
	}

	if err != nil {
		t.Errorf("Error deleting user group: %s with client %+v\n", testGroupMemberGroup.Identifier, client)
	}

	// Remove group
	err = client.DeleteUserGroup(testGroupMemberGroup.Identifier)

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestSetUserGroupParentGroups(t *testing.T) {
	client := New(userGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, userGroupsConfig)
	}

	// Create parent group
	err = client.CreateUserGroup(&testGroupParentGroup)

	if err != nil {
		t.Errorf("Error creating user group: %s with client %+v\n", testGroupParentGroup.Identifier, client)
	}

	permissionItems := []types.GuacPermissionItem{
		client.NewAddGroupMemberPermission(testGroupParentGroup.Identifier),
	}

	err = client.SetUserGroupParentGroups(testGroup.Identifier, &permissionItems)

	if err != nil {
		t.Errorf("Error adding group: %s to user group: %s with client %+v\n", testGroupParentGroup.Identifier, testGroup.Identifier, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestGetUserGroupParentGroups(t *testing.T) {
	client := New(userGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, userGroupsConfig)
	}

	// Get group parent groups
	parents, err := client.GetUserGroupParentGroups(testGroup.Identifier)

	if parents[0] != testGroupParentGroup.Identifier {
		t.Errorf("Wrong member group found: %s expected %s", parents[0], testGroupParentGroup.Identifier)
	}

	// Remove group membership
	permissionItems := []types.GuacPermissionItem{
		client.NewRemoveGroupMemberPermission(testGroupParentGroup.Identifier),
	}

	err = client.SetUserGroupParentGroups(testGroup.Identifier, &permissionItems)

	if err != nil {
		t.Errorf("Error removing group: %s from user group: %s with client %+v\n", testGroupParentGroup.Identifier, testGroup.Identifier, client)
	}

	// Remove group
	err = client.DeleteUserGroup(testGroupParentGroup.Identifier)

	if err != nil {
		t.Errorf("Error deleting user group: %s with client %+v\n", testGroupParentGroup.Identifier, client)
	}

	// Remove group
	err = client.DeleteUserGroup(testGroupParentGroup.Identifier)

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestSetUserGroupPermissions(t *testing.T) {
	client := New(userGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, userGroupsConfig)
	}

	// Add permissions to group
	permissionItems := []types.GuacPermissionItem{
		client.NewAddSystemPermission(types.SystemPermissions{}.Administer()),
		client.NewAddSystemPermission(types.SystemPermissions{}.CreateUser()),
		client.NewAddSystemPermission(types.SystemPermissions{}.CreateConnection()),
		client.NewAddSystemPermission(types.SystemPermissions{}.CreateConnectionGroup()),
		client.NewAddSystemPermission(types.SystemPermissions{}.CreateSharingProfile()),
	}

	err = client.SetUserGroupPermissions(testGroup.Identifier, &permissionItems)

	if err != nil {
		t.Errorf("Error setting user group: %s permissions\n", testGroup.Identifier)
	}

	// Read and verify user group system permissions
	permissions, err := client.GetUserGroupPermissions(testGroup.Identifier)

	if err != nil {
		t.Error(err)
	}

	var missingPermissions []string
	for _, permission := range validSystemPermissions {
		if !types.ArrayContains(&permissions.SystemPermissions, permission) {
			missingPermissions = append(missingPermissions, permission)
		}
	}

	if len(missingPermissions) > 0 {
		t.Errorf("Error checking permissions.  Expected: %s but found: %s\n", strings.Join(validSystemPermissions[:], ", "), strings.Join(missingPermissions[:], ", "))
	}

	// Remove permissions
	permissionItems = []types.GuacPermissionItem{
		client.NewRemoveSystemPermission(types.SystemPermissions{}.Administer()),
		client.NewRemoveSystemPermission(types.SystemPermissions{}.CreateUser()),
		client.NewRemoveSystemPermission(types.SystemPermissions{}.CreateConnection()),
		client.NewRemoveSystemPermission(types.SystemPermissions{}.CreateConnectionGroup()),
		client.NewRemoveSystemPermission(types.SystemPermissions{}.CreateSharingProfile()),
	}

	err = client.SetUserGroupPermissions(testGroup.Identifier, &permissionItems)

	// Read and verify user group system permissions are empty
	permissions, err = client.GetUserGroupPermissions(testGroup.Identifier)

	if err != nil {
		t.Error(err)
	}

	if len(permissions.SystemPermissions) > 0 {
		t.Errorf("Error removing system permissions.  Expected none but found: %s", strings.Join(permissions.SystemPermissions[:], ", "))
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestSetUserGroupUsers(t *testing.T) {
	client := New(userGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, userGroupsConfig)
	}

	// Create test user
	err = client.CreateUser(&testUserGroupUser)
	if err != nil {
		t.Errorf("Error creating guac user %s with error: %s", testUserGroupUser.Username, err)
	}

	// Add user to group
	permissionItems := []types.GuacPermissionItem{
		client.NewAddGroupMemberPermission(testUserGroupUser.Username),
	}

	err = client.SetUserGroupUsers(testGroup.Identifier, &permissionItems)

	if err != nil {
		t.Errorf("Error adding user: %s to group: %s with error: %s\n", testUserGroupUser.Username, testGroup.Identifier, err)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestGettUserGroupUsers(t *testing.T) {
	client := New(userGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, userGroupsConfig)
	}

	// Read member users
	users, err := client.GetUserGroupUsers(testGroup.Identifier)

	if users[0] != testUserGroupUser.Username {
		t.Errorf("Expected member user: %s but found none\n", testUserGroupUser.Username)
	}

	// Remove user
	err = client.DeleteUser(testUserGroupUser.Username)
	if err != nil {
		t.Errorf("Error deleting test user: %s with error: %s\n", testUserGroupUser.Username, err)
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
