// +build all integrationtests

package guacamole

import (
	"fmt"
	"os"
	"testing"

	guac "github.com/techBeck03/guacamole-api-client"
	"github.com/techBeck03/guacamole-api-client/types"
)

var (
	config = guac.Config{
		URI:                    os.Getenv("GUACAMOLE_URL"),
		Username:               os.Getenv("GUACAMOLE_USERNAME"),
		Password:               os.Getenv("GUACAMOLE_PASSWORD"),
		DisableTLSVerification: true,
	}
	testUser = types.GuacUser{Username: "testing"}
)

func TestListUsers(t *testing.T) {
	client := guac.New(config)

	fmt.Printf("URL = %s\n", config.URI)

	err := client.Connect()
	if err != nil {
		t.Errorf("Error %s connecting to guacamole with config %+v", err, config)
	}
	_, err = client.ListUsers()
	if err != nil {
		t.Errorf("Error %s listing users with client %+v", err, client)
	}
}
