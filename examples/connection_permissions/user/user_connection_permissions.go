package main

import (
	"fmt"
	"log"

	guac "github.com/techBeck03/guacamole-api-client"
	"github.com/techBeck03/guacamole-api-client/types"
)

func main() {
	// Change with your values
	client := guac.New(guac.Config{
		URI:                    "https://guac.example.com",
		Username:               "guacadmin",
		Password:               "guacadmin",
		DisableTLSVerification: true,
	})

	// Create login session
	err := client.Connect()

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection successful")
	}

	// Create new user
	newUser := types.GuacUser{
		Username: "testing",
	}

	err = client.CreateUser(&newUser)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("New User\n-------\n%+v", newUser)
	}

	// Create connection group
	userConnectionGroup := types.GuacConnectionGroup{
		Name:             "Testing Users Group",
		ParentIdentifier: "ROOT",
		Type:             "ORGANIZATIONAL",
	}
	err = client.CreateConnectionGroup(&userConnectionGroup)

	if err != nil {
		log.Fatal(err)
	}

	// Create connection
	userConnection := types.GuacConnection{
		Name:             "Testing User Connection",
		Protocol:         "ssh",
		ParentIdentifier: userConnectionGroup.Identifier,
	}
	err = client.CreateConnection(&userConnection)

	if err != nil {
		log.Fatal(err)
	}

	// Create permissions
	userPermissionItems := []types.GuacPermissionItem{
		{
			Op:    "add",
			Path:  fmt.Sprintf("%s/%s", guac.ConnectionGroupPermissionsBasePath, userConnectionGroup.Identifier),
			Value: "READ",
		},
		{
			Op:    "add",
			Path:  fmt.Sprintf("%s/%s", guac.ConnectionPermissionsBasePath, userConnection.Identifier),
			Value: "READ",
		},
	}
	err = client.SetUserConnectionPermissions(newUser.Username, &userPermissionItems)

	if err != nil {
		log.Fatal(err)
	}

	// Read permissions
	permissions, err := client.GetUserPermissions(newUser.Username)
	if err != nil {
		log.Fatal(err)
	}

	// Verify permissions
	_, ok := permissions.ConnectionGroupPermissions[userConnectionGroup.Identifier]

	if !ok {
		log.Fatalf("Error adding connection group: %s permissions for: %s with client %+v", userConnectionGroup.Identifier, newUser.Username, client)
	}

	_, ok = permissions.ConnectionPermissions[userConnection.Identifier]

	if !ok {
		log.Fatalf("Error adding connection: %s permissions for: %s with client %+v", userConnection.Identifier, newUser.Username, client)
	}

	// Remove permissions
	userPermissionItems = []types.GuacPermissionItem{
		{
			Op:    "remove",
			Path:  fmt.Sprintf("%s/%s", guac.ConnectionGroupPermissionsBasePath, userConnectionGroup.Identifier),
			Value: "READ",
		},
		{
			Op:    "remove",
			Path:  fmt.Sprintf("%s/%s", guac.ConnectionPermissionsBasePath, userConnection.Identifier),
			Value: "READ",
		},
	}

	err = client.SetUserConnectionPermissions(newUser.Username, &userPermissionItems)

	if err != nil {
		log.Fatalf("Error %s adding user connection permissions for user: %s with client %+v", err, newUser.Username, client)
	}

	// Verify permissions
	permissions, err = client.GetUserPermissions(newUser.Username)
	if err != nil {
		log.Fatalf("Error %s reading user permissions for: %s with client %+v", err, newUser.Username, client)
	}

	_, ok = permissions.ConnectionGroupPermissions[userConnectionGroup.Identifier]

	if ok {
		log.Fatalf("Error %s deleting connection group permissions for: %s with client %+v", err, newUser.Username, client)
	}

	_, ok = permissions.ConnectionPermissions[userConnection.Identifier]

	if ok {
		log.Fatalf("Error %s deleting connection permissions for: %s with client %+v", err, newUser.Username, client)
	}

	// Delete connection group and connection
	err = client.DeleteConnection(userConnection.Identifier)

	if err != nil {
		log.Fatalf("Error %s deleting connection: %s with client %+v", err, userConnection.Identifier, client)
	}

	err = client.DeleteConnectionGroup(userConnectionGroup.Identifier)

	if err != nil {
		log.Fatalf("Error %s deleting connection group: %s with client %+v", err, userConnectionGroup.Identifier, client)
	}

	// Delete user
	err = client.DeleteUser(newUser.Username)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully deleted user")
	}

	err = client.Disconnect()

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Disconnect successful")
	}
}
