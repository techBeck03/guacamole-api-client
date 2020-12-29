// +build all unittests

package guacamole

import (
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
	testGroup = types.GuacUserGroup{Identifier: "(testing) Test Group"}
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
