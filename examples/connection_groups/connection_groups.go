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

	// Create new connection group
	newConnectionGroup := types.GuacConnectionGroup{
		Name:             "Testing Group",
		ParentIdentifier: "ROOT",
		Type:             "ORGANIZATIONAL",
	}

	err = client.CreateConnectionGroup(&newConnectionGroup)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%+v", newConnectionGroup)
	}

	// Update connection group
	newConnectionGroup.Type = "BALANCING"

	err = client.UpdateConnectionGroup(&newConnectionGroup)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("update successful")
	}

	// Read connection group by identifier
	readConnectionGroup, err := client.ReadConnectionGroup(newConnectionGroup.Identifier)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Read connection group:\n%+v\n", readConnectionGroup)
	}

	// Read connection group by Path
	readConnectionGroup, err = client.ReadConnectionGroupByPath("ROOT/Testing Group")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Read connection group:\n%+v\n", readConnectionGroup)
	}

	// Delete connection group
	err = client.DeleteConnectionGroup(newConnectionGroup.Identifier)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully deleted connection")
	}

	err = client.Disconnect()

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Disconnect successful")
	}
}
