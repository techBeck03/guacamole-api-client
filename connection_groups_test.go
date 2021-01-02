// +build all unittests

package guacamole

import (
	"fmt"
	"os"
	"testing"

	"github.com/techBeck03/guacamole-api-client/types"
)

var (
	connectionGroupsConfig = Config{
		URL:                    os.Getenv("GUACAMOLE_URL"),
		Username:               os.Getenv("GUACAMOLE_USERNAME"),
		Password:               os.Getenv("GUACAMOLE_PASSWORD"),
		DisableTLSVerification: true,
	}
	testConnectionGroup = types.GuacConnectionGroup{
		Name:             "Testing Group",
		ParentIdentifier: "ROOT",
		Type:             "ORGANIZATIONAL",
	}
)

func TestListConnectionGroups(t *testing.T) {
	client := New(connectionGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, connectionGroupsConfig)
	}

	_, err = client.ListConnectionGroups()
	if err != nil {
		t.Errorf("Error %s listing connection group with client %+v", err, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestCreateConnectionGroup(t *testing.T) {
	client := New(connectionGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, connectionGroupsConfig)
	}

	err = client.CreateConnectionGroup(&testConnectionGroup)
	if err != nil {
		t.Errorf("Error %s creating connection group: %s with client %+v", err, testConnectionGroup.Name, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestReadConnectionGroup(t *testing.T) {
	client := New(connectionGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, connectionGroupsConfig)
	}

	connectionGroup, err := client.ReadConnectionGroup(testConnectionGroup.Identifier)
	if err != nil {
		t.Errorf("Error %s reading connection group: %s with client %+v", err, testConnectionGroup.Name, client)
	}

	if connectionGroup.Name != testConnectionGroup.Name {
		t.Errorf("Expected connection name = %s read connection name = %s", testConnectionGroup.Name, connectionGroup.Name)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestReadConnectionGroupByPath(t *testing.T) {
	client := New(connectionGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, connectionGroupsConfig)
	}

	connectionGroup, err := client.ReadConnectionGroupByPath(fmt.Sprintf("%s/%s", testConnectionGroup.ParentIdentifier, testConnectionGroup.Name))
	if err != nil {
		t.Errorf("Error %s reading connection by path: %s with client %+v", err, testConnectionGroup.Name, client)
	}

	if connectionGroup.Name != testConnectionGroup.Name {
		t.Errorf("Expected connection group name = %s read connection group name = %s", fmt.Sprintf("%s/%s", testConnectionGroup.ParentIdentifier, testConnectionGroup.Name), connectionGroup.Name)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestUpdateConnectionGroup(t *testing.T) {
	client := New(connectionGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, connectionGroupsConfig)
	}

	connectionGroup, err := client.ReadConnectionGroup(testConnectionGroup.Identifier)
	if err != nil {
		t.Errorf("Error %s reading connection group: %s with client %+v", err, testConnectionGroup.Identifier, client)
	}

	connectionGroup.Type = "BALANCING"

	err = client.UpdateConnectionGroup(&connectionGroup)

	if err != nil {
		t.Errorf("Error %s updating connection group: %s with client %+v", err, testConnectionGroup.Identifier, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestDeleteConnectionGroup(t *testing.T) {
	client := New(connectionGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, connectionGroupsConfig)
	}

	err = client.DeleteConnectionGroup(testConnectionGroup.Identifier)
	if err != nil {
		t.Errorf("Error %s deleting connection group: %s with client %+v", err, testConnectionGroup.Identifier, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}
