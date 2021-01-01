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

	createdConnection, err := client.CreateConnection(&newConnection)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%+v", createdConnection)
	}

	// Update Connection
	createdConnection.Properties.Hostname = "testing.example.com"
	createdConnection.Properties.Port = "22"
	createdConnection.Attributes.MaxConnections = "2"

	err = client.UpdateConnection(&createdConnection)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("update successful")
	}

	// Read Connection by identifier
	readConnection, err := client.ReadConnection(createdConnection.Identifier)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Read connection:\n%+v\n", readConnection)
	}

	// Read Connection by Path
	readConnection, err = client.ReadConnectionByPath("ROOT/Testing Connection")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Read connection:\n%+v\n", readConnection)
	}

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Read connection:\n%+v\n", readConnection)
	}

	err = client.DeleteConnection(createdConnection.Identifier)

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
