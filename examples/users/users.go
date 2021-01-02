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
