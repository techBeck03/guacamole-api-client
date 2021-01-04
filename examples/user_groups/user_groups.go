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
