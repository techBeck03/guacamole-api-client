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
}
