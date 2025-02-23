package auth

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUsername() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter username: ")
	username, _ := reader.ReadString('\n')
	return strings.TrimSpace(username)
}
