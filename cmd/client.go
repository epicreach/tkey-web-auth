package main

import (
	"chalmers/tkey-group22/internal/util"
	"fmt"
)

func main() {

	// Gets mode from user inputs and runs selected mode. Loops until program is told to exit.
	for {
		mode := util.SelectMode()

		switch mode {
		case 1:
			// Perform register
			util.CallRegister()
		case 2:
			// Perform login
			util.CallLogin()
		case 3:
			// Stop program
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
