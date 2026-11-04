package main

import (
	"fmt"

	"github.com/ayush/systeminfo"
)

func main() {
	fmt.Println("System Information")
	fmt.Println("====================")

	fmt.Println("OS:", systeminfo.GetOS())
	fmt.Println("Architecture:", systeminfo.GetArchitecture())
}
