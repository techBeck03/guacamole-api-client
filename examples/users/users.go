package main

import (
	"fmt"
	"log"
	"strings"

	guac "github.com/techBeck03/guacamole-api-client"
	"github.com/techBeck03/guacamole-api-client/types"
)

func main() {
	// Change with your values
	client := guac.New(guac.Config{
		URL:                    "https://guac.example.com",
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
		fmt.Printf("%+v", newUser)
	}

	// Update user
	newUser.Attributes.GuacFullName = "Go Testing User"

	err = client.UpdateUser(&newUser)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("update successful")
	}

	// Read user by username
	readUser, err := client.ReadUser(newUser.Username)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Read user:\n%+v\n", readUser)
	}

	// Add user system permissions
	permissionItems := []types.GuacPermissionItem{
		client.NewAddSystemPermission(types.SystemPermissions{}.Administer()),
		client.NewAddSystemPermission(types.SystemPermissions{}.CreateUser()),
		client.NewAddSystemPermission(types.SystemPermissions{}.CreateConnection()),
		client.NewAddSystemPermission(types.SystemPermissions{}.CreateConnectionGroup()),
		client.NewAddSystemPermission(types.SystemPermissions{}.CreateSharingProfile()),
	}

	err = client.SetUserPermissions(newUser.Username, &permissionItems)

	if err != nil {
		log.Fatal(err)
	}

	// Read and verify user group system permissions
	permissions, err := client.GetUserPermissions(newUser.Username)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("System permissions = %s\n", strings.Join(permissions.SystemPermissions[:], ", "))

	// Add user to group
	// Create new user group
	newGroup := types.GuacUserGroup{
		Identifier: "(testing) Test Group",
	}

	permissionItems = []types.GuacPermissionItem{
		client.NewAddGroupMemberPermission(newGroup.Identifier),
	}

	err = client.SetUserGroupMembership(newUser.Username, &permissionItems)

	if err != nil {
		log.Fatal(err)
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
