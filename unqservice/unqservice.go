package main

import (
	"fmt"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		// dont care...
	}
	username := user.Username
	fmt.Printf("Username: %s\n", username)
}
