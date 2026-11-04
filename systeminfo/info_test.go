package systeminfo

import (
	"runtime"
	"testing"
)

func TestGetOS(t *testing.T) {
	got := GetOS()

	if got != runtime.GOOS {
		t.Errorf("GetOS() = %v, want %v", got, runtime.GOOS)
	}
}

func TestGetArchitecture(t *testing.T) {
	got := GetArchitecture()

	if got != runtime.GOARCH {
		t.Errorf("GetArchitecture() = %v, want %v", got, runtime.GOARCH)
	}
}

func TestGetHostname(t *testing.T) {
	got, err := GetHostname()

	if err != nil {
		t.Fatalf("GetHostname() returned error: %v", err)
	}

	if got == "" {
		t.Error("GetHostname() returned an empty hostname")
	}
}

func TestGetCPUCount(t *testing.T) {
	got := GetCPUCount()

	if got <= 0 {
		t.Errorf("GetCPUCount() = %v, want a positive number", got)
	}
}

func TestGetGoVersion(t *testing.T) {
	got := GetGoVersion()

	if got != runtime.Version() {
		t.Errorf("GetGoVersion() = %v, want %v", got, runtime.Version())
	}
}
