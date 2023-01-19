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
	connectionGroupsConfig = Config{
		URL:                    os.Getenv("GUACAMOLE_URL"),
		Username:               os.Getenv("GUACAMOLE_USERNAME"),
		Password:               os.Getenv("GUACAMOLE_PASSWORD"),
		Token:                  os.Getenv("GUACAMOLE_TOKEN"),
		DataSource:             os.Getenv("GUACAMOLE_DATA_SOURCE"),
		DisableTLSVerification: true,
	}
	testConnectionGroup = types.GuacConnectionGroup{
		Name: "Testing Group",
		Type: "ORGANIZATIONAL",
	}
)

func TestListConnectionGroups(t *testing.T) {
	if os.Getenv("GUACAMOLE_COOKIES") != "" {
		connectionGroupsConfig.Cookies = make(map[string]string)
		for _, e := range strings.Split(os.Getenv("GUACAMOLE_COOKIES"), ",") {
			cookie_split := strings.Split(e, "=")
			connectionGroupsConfig.Cookies[cookie_split[0]] = cookie_split[1]
		}
	}
	client := New(connectionGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, connectionGroupsConfig)
	}

	_, err = client.ListConnectionGroups()
	if err != nil {
		t.Errorf("Error %s listing connection group with client %+v", err, client)
	}
}

func TestCreateConnectionGroup(t *testing.T) {
	client := New(connectionGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, connectionGroupsConfig)
	}

	grp, err := client.ReadConnectionGroupByPath(os.Getenv("GUACAMOLE_CONNECTION_PATH"))
	if err != nil {
		t.Errorf("Error unable to find parent group with path: %s", os.Getenv("GUACAMOLE_CONNECTION_PATH"))
	}

	testConnectionGroup.ParentIdentifier = grp.Identifier
	testConnectionGroup.Path = fmt.Sprintf("%s/%s", grp.Path, testConnectionGroup.Name)

	err = client.CreateConnectionGroup(&testConnectionGroup)
	if err != nil {
		t.Errorf("Error %s creating connection group: %s with client %+v", err, testConnectionGroup.Name, client)
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
}

func TestReadConnectionGroupByPath(t *testing.T) {
	client := New(connectionGroupsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, connectionGroupsConfig)
	}

	_, err = client.ReadConnectionGroupByPath(testConnectionGroup.Path)
	if err != nil {
		t.Errorf("Error %s reading connection by path: %s with client %+v", err, testConnectionGroup.Path, client)
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
}
