// +build all unittests

package guacamole

import (
	"os"
	"testing"

	"github.com/techBeck03/guacamole-api-client/types"
)

var (
	config = Config{
		URI:                    os.Getenv("GUACAMOLE_URL"),
		Username:               os.Getenv("GUACAMOLE_USERNAME"),
		Password:               os.Getenv("GUACAMOLE_PASSWORD"),
		DisableTLSVerification: true,
	}
	testUser = types.GuacUser{Username: "testing"}
)

func TestListUsers(t *testing.T) {
	client := New(config)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, config)
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
	client := New(config)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, config)
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
	client := New(config)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, config)
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
	client := New(config)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, config)
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

func TestDeleteUser(t *testing.T) {
	client := New(config)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, config)
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
