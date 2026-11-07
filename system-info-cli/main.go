package main

import (
	"fmt"

	"github.com/sahu-ayush/go-library-poc/systeminfo"
)

func main() {
	fmt.Println("System Information")
	fmt.Println("====================")

	fmt.Println("OS:", systeminfo.GetOS())
	fmt.Println("Architecture:", systeminfo.GetArchitecture())

	hostname, err := systeminfo.GetHostname()
	if err != nil {
		fmt.Println("Error getting Hostname:", err)
		return
	}

	fmt.Println("Hostname:", hostname)
	fmt.Println("CPU Count:", systeminfo.GetCPUCount())
	fmt.Println("Go Version:", systeminfo.GetGoVersion())
	fmt.Println("System Summary:", systeminfo.GetSystemSummary())
}
