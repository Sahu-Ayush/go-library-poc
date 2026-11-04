package systeminfo

import "runtime"

// GetOS returns the operating system.
func GetOS() string {
	return runtime.GOOS
}

// GetArchitecture returns the system architecture.
func GetArchitecture() string {
	return runtime.GOARCH
}
