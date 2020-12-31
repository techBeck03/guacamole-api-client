// +build all unittests

package guacamole

import (
	"fmt"
	"os"
	"testing"

	"github.com/techBeck03/guacamole-api-client/types"
)

var (
	connectionsConfig = Config{
		URI:                    os.Getenv("GUACAMOLE_URL"),
		Username:               os.Getenv("GUACAMOLE_USERNAME"),
		Password:               os.Getenv("GUACAMOLE_PASSWORD"),
		DisableTLSVerification: true,
	}
	testConnection = types.GuacConnection{
		Name:             "Testing Connection",
		Protocol:         "ssh",
		ParentIdentifier: "ROOT",
	}
)

func TestListConnections(t *testing.T) {
	client := New(connectionsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, connectionsConfig)
	}

	_, err = client.ListConnections()
	if err != nil {
		t.Errorf("Error %s listing users with client %+v", err, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestCreateConnection(t *testing.T) {
	client := New(connectionsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, connectionsConfig)
	}

	createdConnection, err := client.CreateConnection(&testConnection)
	if err != nil {
		t.Errorf("Error %s creating connection: %s with client %+v", err, testConnection.Name, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}

	testConnection = createdConnection
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

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}

func TestReadConnectionByPath(t *testing.T) {
	client := New(connectionsConfig)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, connectionsConfig)
	}

	connection, err := client.ReadConnectionByPath(fmt.Sprintf("%s/%s", testConnection.ParentIdentifier, testConnection.Name))
	if err != nil {
		t.Errorf("Error %s reading connection by path: %s with client %+v", err, testConnection.Name, client)
	}

	if connection.Name != testConnection.Name {
		t.Errorf("Expected connection name = %s read connection name = %s", fmt.Sprintf("%s/%s", testConnection.ParentIdentifier, testConnection.Name), connection.Name)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
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

	connection.Properties.Hostname = "testing.example.com"
	connection.Properties.Port = "22"
	connection.Attributes.MaxConnections = "2"

	err = client.UpdateConnection(&connection)

	if err != nil {
		t.Errorf("Error %s updating connection: %s with client %+v", err, testConnection.Identifier, client)
	}

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
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

	err = client.Disconnect()

	if err != nil {
		t.Errorf("Disconnect errors: %s\n", err)
	}
}
