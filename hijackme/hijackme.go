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
	//First stop service
	println("Stopping the service")
	cmd1 := exec.Command("sc.exe", "stop", "WinDefend")
	err1 := cmd1.Run()

	if err1 != nil {
		log.Fatal(err1)
	}

	//Wait for the service to stop
	time.Sleep(5 * time.Second)

	println("disabling the service")

	//Now disabling the service
	cmd2 := exec.Command("sc.exe", "config", "WinDefend", "start=", "auto")
	err2 := cmd2.Run()

	if err2 != nil {
		log.Fatal(err2)
	}

	//Final sleep for debugging purpose
	time.Sleep(3 * time.Second)
}
