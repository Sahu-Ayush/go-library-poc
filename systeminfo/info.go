package systeminfo

import "runtime"

// GetSystemInfo returns information about the system.
func GetOS() string{
	return runtime.GOOS
}

func GetArchitecture() string (){
	return runtime.GOARCH
}

