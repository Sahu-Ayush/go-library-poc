package systeminfo

import (
	"os"
	"runtime"
)

// GetOS returns the operating system.
func GetOS() string {
	return runtime.GOOS
}

// GetArchitecture returns the system architecture.
func GetArchitecture() string {
	return runtime.GOARCH
}

// GetHostname returns the machine hostname.
func GetHostname() (string, error) {
	return os.Hostname()
}

// GetCPUCount returns the number of logical CPUs.
func GetCPUCount() int {
	return runtime.NumCPU()
}

// GetGoVersion returns the Go runtime version.
func GoGetVersion() string {
	return runtime.Version()
}
