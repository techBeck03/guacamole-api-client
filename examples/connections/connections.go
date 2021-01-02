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

	// Create new connection
	newConnection := types.GuacConnection{
		Name:             "Testing Connection",
		Protocol:         "ssh",
		ParentIdentifier: "ROOT",
	}

	err = client.CreateConnection(&newConnection)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%+v", newConnection)
	}

	// Update connection
	newConnection.Properties.Hostname = "testing.example.com"
	newConnection.Properties.Port = "22"
	newConnection.Attributes.MaxConnections = "2"

	err = client.UpdateConnection(&newConnection)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("update successful")
	}

	// Read connection by identifier
	readConnection, err := client.ReadConnection(newConnection.Identifier)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Read connection:\n%+v\n", readConnection)
	}

	// Read connection by Path
	readConnection, err = client.ReadConnectionByPath("ROOT/Testing Connection")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Read connection:\n%+v\n", readConnection)
	}

	// Delete connection
	err = client.DeleteConnection(newConnection.Identifier)

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
