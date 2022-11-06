package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
)

func main() {

	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	username := user.Username

	fmt.Printf("Username: %s\n", username)

	file, err := os.Create("C:\\Program Files\\UNQ Service\\common.txt")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString(("Username: " + username))
}
