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

	// Create new user group
	newGroup := types.GuacUserGroup{
		Identifier: "(testing) Test Group",
	}

	err = client.CreateUserGroup(&newGroup)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("New Group\n--------\n%+v\n", newGroup)
	}

	// Update user group
	newGroup.Attributes.Disabled = "true"

	err = client.UpdateUserGroup(&newGroup)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("update successful")
	}

	// Read group by identifier
	readGroup, err := client.ReadUserGroup(newGroup.Identifier)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Read user:\n%+v\n", readGroup)
	}

	// Add parent user group
	parentGroup := types.GuacUserGroup{
		Identifier: "(testing) Test Parent Group",
	}

	err = client.CreateUserGroup(&parentGroup)

	if err != nil {
		log.Fatal(err)
	}

	permissionItems := []types.GuacPermissionItem{
		client.NewAddGroupMemberPermission(parentGroup.Identifier),
	}

	err = client.SetUserGroupParentGroups(newGroup.Identifier, &permissionItems)

	// Add member user group
	memberGroup := types.GuacUserGroup{
		Identifier: "(testing) Test Member Group",
	}

	err = client.CreateUserGroup(&memberGroup)

	if err != nil {
		log.Fatal(err)
	}

	permissionItems = []types.GuacPermissionItem{
		client.NewAddGroupMemberPermission(memberGroup.Identifier),
	}

	err = client.SetUserGroupMemberGroups(newGroup.Identifier, &permissionItems)

	// Add permissions to group
	permissionItems = []types.GuacPermissionItem{
		client.NewAddSystemPermission(types.SystemPermissions{}.Administer()),
		client.NewAddSystemPermission(types.SystemPermissions{}.CreateUser()),
		client.NewAddSystemPermission(types.SystemPermissions{}.CreateConnection()),
		client.NewAddSystemPermission(types.SystemPermissions{}.CreateConnectionGroup()),
		client.NewAddSystemPermission(types.SystemPermissions{}.CreateSharingProfile()),
	}

	err = client.SetUserGroupPermissions(newGroup.Identifier, &permissionItems)

	if err != nil {
		log.Fatal(err)
	}

	// Read and verify user group system permissions
	permissions, err := client.GetUserGroupPermissions(newGroup.Identifier)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("System permissions = %s\n", strings.Join(permissions.SystemPermissions[:], ", "))

	// Delete user group
	err = client.DeleteUserGroup(newGroup.Identifier)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully deleted user group")
	}

	err = client.Disconnect()

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Disconnect successful")
	}
}
