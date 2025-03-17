package util

import (
	"bufio"
	"chalmers/tkey-group22/client/internal/auth"
	"fmt"
	"log"
	"os"
	"strings"
)

var le = log.New(os.Stderr, "Error: ", 0)
var appurl = "http://localhost:8080"

// SelectMode prompts the user to select a mode of operation
//
// Returns:
// - int: The selected mode of operation
func SelectMode() int {
	fmt.Println("\nSelect Mode:")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Exit")

	var choice int
	fmt.Print("Enter choice (1/2/3): ")
	fmt.Scanln(&choice)
	return choice
}

// CallLogin retrieves the username and attempts to log in the user using the provided app URL
// If an error occurs during the login process, it prints the error
func CallLogin() {
	username := getUsername()
	_, _, err := auth.GetAndSign(appurl, username)
	if err != nil {
		le.Println(err)
	} else {
		fmt.Printf("User '%s' has been successfully logged in!\n", username)
	}
}

// CallRegister retrieves the username and label, and attempts to register it with the authentication service
// If an error occurs during the registration process, it prints the error
func CallRegister() {
	username := getUsername()
	label := getLabel()
	err := auth.Register(appurl, username, label)
	if err != nil {
		le.Println(err)
	}
}

// getUsername gets the username from the user
//
// Returns:
// - string: The username entered by the user
func getUsername() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter username: ")
	username, _ := reader.ReadString('\n')
	return strings.TrimSpace(username)
}

// getLabel gets the label for the public key from the user
//
// Returns:
// - string: The label entered by the user
func getLabel() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter label for the public key: ")
	label, _ := reader.ReadString('\n')
	return strings.TrimSpace(label)
}
