package util

import (
	"chalmers/tkey-group22/internal/auth"
	"fmt"
	"log"
	"os"
)

var le = log.New(os.Stderr, "Error: ", 0)

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

func CallLogin() {
	err := auth.Login()
	if err != nil {
		le.Println(err)
	}
}

func CallRegister() {
	err := auth.Register()
	if err != nil {
		le.Println(err)
	}
}
