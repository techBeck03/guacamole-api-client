//go:build all || unittests || specific
// +build all unittests specific

package guacamole

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/techBeck03/guacamole-api-client/types"
)

var (
	connectionsConfig = Config{
		URL:                    os.Getenv("GUACAMOLE_URL"),
		Username:               os.Getenv("GUACAMOLE_USERNAME"),
		Password:               os.Getenv("GUACAMOLE_PASSWORD"),
		Token:                  os.Getenv("GUACAMOLE_TOKEN"),
		DataSource:             os.Getenv("GUACAMOLE_DATA_SOURCE"),
		DisableTLSVerification: true,
	}
	testConnection = types.GuacConnection{
		Name:             "Test Connection",
		Protocol:         "ssh",
		ParentIdentifier: "1592",
		Path:             fmt.Sprintf("%s/Test Connection", os.Getenv("GUACAMOLE_CONNECTION_PATH")),
	}
)

func TestListConnections(t *testing.T) {
	if os.Getenv("GUACAMOLE_COOKIES") != "" {
		connectionsConfig.Cookies = make(map[string]string)
		for _, e := range strings.Split(os.Getenv("GUACAMOLE_COOKIES"), ",") {
			cookie_split := strings.Split(e, "=")
			connectionsConfig.Cookies[cookie_split[0]] = cookie_split[1]
		}
	}
	client := New(connectionsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, connectionsConfig)
	}

	_, err = client.ListConnections()
	if err != nil {
		t.Errorf("Error %s listing connections with client %+v", err, client)
	}
}

func TestCreateConnection(t *testing.T) {
	client := New(connectionsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, connectionsConfig)
	}

	err = client.CreateConnection(&testConnection)
	if err != nil {
		t.Errorf("Error %s creating connection: %s with client %+v", err, testConnection.Name, client)
	}

}

func TestReadConnection(t *testing.T) {
	client := New(connectionsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, connectionsConfig)
	}

	connection, err := client.ReadConnection(testConnection.Identifier)
	if err != nil {
		t.Errorf("Error %s reading connection: %s with client %+v", err, testConnection.Name, client)
	}

	if connection.Name != testConnection.Name {
		t.Errorf("Expected connection name = %s read connection name = %s", testConnection.Name, connection.Name)
	}

}

func TestReadConnectionByPath(t *testing.T) {
	client := New(connectionsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, connectionsConfig)
	}

	_, err = client.ReadConnectionByPath(fmt.Sprintf("%s", testConnection.Path))
	if err != nil {
		t.Errorf("Error %s reading connection by path: %s with client %+v", err, testConnection.Path, client)
	}

}

func TestUpdateConnection(t *testing.T) {
	client := New(connectionsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, connectionsConfig)
	}

	connection, err := client.ReadConnection(testConnection.Identifier)
	if err != nil {
		t.Errorf("Error %s reading connection: %s with client %+v", err, testConnection.Identifier, client)
	}

	connection.Parameters.Hostname = "testing.example.com"
	connection.Parameters.Port = "22"
	connection.Attributes.MaxConnections = "2"

	err = client.UpdateConnection(&connection)

	if err != nil {
		t.Errorf("Error %s updating connection: %s with client %+v", err, testConnection.Identifier, client)
	}

}

func TestDeleteConnection(t *testing.T) {
	client := New(connectionsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, connectionsConfig)
	}

	err = client.DeleteConnection(testConnection.Identifier)
	if err != nil {
		t.Errorf("Error %s deleting connection: %s with client %+v", err, testConnection.Identifier, client)
	}

}
