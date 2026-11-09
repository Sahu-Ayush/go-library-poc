# Go Packages, Modules & Reusable Libraries — Learning & POC

![Go](https://img.shields.io/badge/Go-1.25.5-00ADD8?logo=go&logoColor=white)
![Version](https://img.shields.io/badge/Library-v1.1.0-blue)
![Status](https://img.shields.io/badge/Status-Completed-success)

A hands-on learning project to understand Go packages, modules, reusable libraries, dependency management, semantic versioning, Git tags, and how to publish and consume Go libraries from GitHub.

This repository demonstrates how to build a reusable Go library, publish versioned releases, and consume those releases from an independent Go application.

---

## Table of Contents

1. [Project Overview](#1-project-overview)
2. [Learning Objectives](#2-learning-objectives)
3. [Project Architecture](#3-project-architecture)
4. [Go Fundamentals](#4-go-fundamentals)
5. [Project Structure](#5-project-structure)
6. [Library Implementation](#6-library-implementation)
7. [Unit Testing](#7-unit-testing)
8. [Consumer Application](#8-consumer-application)
9. [Local Module Development](#9-local-module-development)
10. [Publishing a Go Library](#10-publishing-a-go-library)
11. [Consuming a Remote Go Module](#11-consuming-a-remote-go-module)
12. [Module Version Upgrade](#12-module-version-upgrade)
13. [Important Go Commands](#13-important-go-commands)
14. [Understanding go.mod and go.sum](#14-understanding-gomod-and-gosum)
15. [Git Tags and Go Module Versions](#15-git-tags-and-go-module-versions)
16. [Key Learnings](#16-key-learnings)
17. [Interview Preparation](#17-interview-preparation)
18. [Future Improvements](#18-future-improvements)
19. [References](#19-references)

---

## 1. Project Overview

The objective of this POC is to understand how Go code can be organized into reusable packages and distributed as versioned modules.

Instead of writing all functionality inside one application, we create a separate Go library that exposes reusable functions through a public API.

A second Go application imports and consumes that library.

We then demonstrate the complete library lifecycle:

- Create a reusable Go package.
- Initialize it as an independent Go module.
- Implement exported functions.
- Write unit tests.
- Create a separate consumer application.
- Consume the library locally using `replace`.
- Publish the library to GitHub using version tags.
- Download and consume the published module.
- Upgrade the consumer to a newer library version.
- Verify that new functionality is available after upgrading.

### Technologies

| Technology | Purpose |
|---|---|
| Go 1.25.5 | Programming language |
| Git | Source control |
| GitHub | Repository and library publishing |
| Go Modules | Dependency and version management |
| Go Testing | Unit testing |
| macOS Apple Silicon | Development environment |

---

## 2. Learning Objectives

By completing this POC, the following concepts were explored:

### Go fundamentals
- Packages and modules.
- Libraries versus applications.
- `package main` and `func main()`.
- Exported and unexported identifiers.
- Public APIs and abstraction.
- Package imports.

### Go module management
- `go.mod` and `go.sum`.
- Module paths.
- Dependency declarations.
- Local module replacement.
- Remote module consumption.
- Go module cache.
- Dependency version upgrades.

### Library publishing
- Git commits and tags.
- Semantic Versioning.
- Nested Go modules.
- Publishing versioned libraries.
- Consuming tagged releases.
- Backward-compatible library enhancements.

### Testing
- Unit testing with the `testing` package.
- Testing exported functions.
- Using `*testing.T`.
- Running tests without cached results.

---

## 3. Project Architecture

This repository contains two independent Go modules inside one Git repository.

```text
GitHub Repository
go-library-poc
        |
        |
        +-----------------------------+
        |                             |
        v                             v
  systeminfo/                  system-info-cli/
  Library Module               Consumer Module
        |                             |
        |                             |
  Reusable Functions            Imports Library
        |                             |
        |                             |
  Exported API  <--------------- Calls API
        |
        |
  Published Releases
  v1.0.0
  v1.1.0
```

### Architecture explanation

**Library module — systeminfo**

Responsible for providing reusable system information functionality.

It exposes functions such as:

- `GetOS()`
- `GetArchitecture()`
- `GetHostname()`
- `GetCPUCount()`
- `GetGoVersion()`
- `GetSystemSummary()`

**Consumer module — system-info-cli**

Responsible for consuming the library and displaying the system information.

It imports the library using its module/package path and calls the exported functions.

The two modules have independent `go.mod` files and can be managed independently.

---

## 4. Go Fundamentals

### 4.1 What is a Go package?

A package is a collection of Go source files in the same directory that share the same package declaration.

Example:

```go
package systeminfo
```

Packages help organize code into logical and reusable units.

### 4.2 What is a Go module?

A module is a collection of related Go packages managed together through a `go.mod` file.

Example:

```go
module github.com/sahu-ayush/go-library-poc/systeminfo

go 1.25.5
```

The module path identifies the module and is used when importing its packages.

### 4.3 Package vs Module vs Library

| Concept | Meaning |
|---|---|
| Package | Organizes related Go source files |
| Module | Defines a versioned collection of Go packages |
| Library | Reusable functionality provided to other applications |

A module can contain multiple packages.

A library is commonly distributed as a Go module.

### 4.4 Application vs Library

| Library | Application |
|---|---|
| Provides reusable functionality | Uses functionality to perform a task |
| Exposes APIs | Calls APIs |
| Usually does not contain an executable entry point | Commonly contains `package main` |
| Can be consumed by multiple applications | Produces an executable or service |

### 4.5 Exported and unexported identifiers

Go uses identifier capitalization to determine visibility across packages.

```go
func GetOS() string {
    return runtime.GOOS
}
```

`GetOS` is exported because it begins with an uppercase letter.

```go
func getOS() string {
    return runtime.GOOS
}
```

`getOS` is unexported because it begins with a lowercase letter.

Only exported identifiers can be accessed directly from another package.

---

## 5. Project Structure

```text
go-library-poc/
│
├── systeminfo/
│   ├── go.mod
│   ├── info.go
│   └── info_test.go
│
├── system-info-cli/
│   ├── go.mod
│   ├── go.sum
│   └── main.go
│
└── README.md
```

### Important observation

The Git repository root is not itself a Go module.

Instead, it contains two independent Go modules:

```text
systeminfo/go.mod
system-info-cli/go.mod
```

This allows the library and consumer application to have independent dependency configurations and versioning.

---

## 6. Library Implementation

### 6.1 Initialize the library module

Navigate to the library directory:

```bash
cd systeminfo
```

Initialize the module:

```bash
go mod init github.com/sahu-ayush/go-library-poc/systeminfo
```

The resulting `go.mod`:

```go
module github.com/sahu-ayush/go-library-poc/systeminfo

go 1.25.5
```

### 6.2 Implement reusable functions

File: `systeminfo/info.go`

```go
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
func GetGoVersion() string {
    return runtime.Version()
}

// GetSystemSummary returns a short system summary.
func GetSystemSummary() string {
    return runtime.GOOS + "/" + runtime.GOARCH
}
```

### Functions explained

| Function | Return type | Description |
|---|---|---|
| `GetOS()` | `string` | Returns operating system |
| `GetArchitecture()` | `string` | Returns CPU architecture |
| `GetHostname()` | `string, error` | Returns hostname |
| `GetCPUCount()` | `int` | Returns logical CPU count |
| `GetGoVersion()` | `string` | Returns Go runtime version |
| `GetSystemSummary()` | `string` | Returns OS and architecture combined |

---

## 7. Unit Testing

File: `systeminfo/info_test.go`

The library uses Go's standard `testing` package.

Example:

```go
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

func TestGetSystemSummary(t *testing.T) {
    got := GetSystemSummary()
    want := runtime.GOOS + "/" + runtime.GOARCH

    if got != want {
        t.Errorf("GetSystemSummary() = %v, want %v", got, want)
    }
}
```

### Run tests

```bash
go test ./...
```

Run tests with detailed output:

```bash
go test -v ./...
```

Force a fresh test execution:

```bash
go test -v -count=1 ./...
```

Expected result:

```text
=== RUN   TestGetOS
--- PASS: TestGetOS

=== RUN   TestGetArchitecture
--- PASS: TestGetArchitecture

=== RUN   TestGetHostname
--- PASS: TestGetHostname

=== RUN   TestGetCPUCount
--- PASS: TestGetCPUCount

=== RUN   TestGetGoVersion
--- PASS: TestGetGoVersion

=== RUN   TestGetSystemSummary
--- PASS: TestGetSystemSummary

PASS
```

All six library tests passed.

---

## 8. Consumer Application

The second module consumes the reusable library.

### 8.1 Initialize the consumer module

```bash
cd system-info-cli

go mod init github.com/ayush/system-info-cli
```

### 8.2 Import the library

File: `system-info-cli/main.go`

```go
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
        fmt.Println("Error getting hostname:", err)
        return
    }

    fmt.Println("Hostname:", hostname)
    fmt.Println("CPU Count:", systeminfo.GetCPUCount())
    fmt.Println("Go Version:", systeminfo.GetGoVersion())
    fmt.Println("System Summary:", systeminfo.GetSystemSummary())
}
```

### 8.3 Expected output

```text
System Information
====================
OS: darwin
Architecture: arm64
Hostname: <machine-hostname>
CPU Count: 11
Go Version: go1.25.5
System Summary: darwin/arm64
```

The CLI does not implement the system information functions itself.

It imports and consumes them from the reusable library.

---

## 9. Local Module Development

Before publishing the library, we tested local consumption using the `replace` directive.

### Why use replace?

During local development, you may want to consume a module directly from your local filesystem rather than downloading it from a remote repository.

Example `system-info-cli/go.mod`:

```go
module github.com/ayush/system-info-cli

go 1.25.5

require github.com/sahu-ayush/go-library-poc/systeminfo v0.0.0

replace github.com/sahu-ayush/go-library-poc/systeminfo => ../systeminfo
```

### Understand the configuration

| Directive | Purpose |
|---|---|
| `module` | Defines consumer module identity |
| `go` | Specifies Go language version |
| `require` | Declares the dependency |
| `replace` | Redirects dependency resolution to a local directory |

### Run the application

```bash
go test ./...
go run .
```

### Important distinction

With `replace`, Go resolves the library from:

```text
../systeminfo
```

It does not need to download the library from GitHub.

This is useful for local development and testing changes across modules.

---

## 10. Publishing a Go Library

After validating the library, we published it to GitHub.

### 10.1 Git repository

Repository:

https://github.com/Sahu-Ayush/go-library-poc

Both modules exist inside this Git repository.

### 10.2 Understand nested Go module versioning

The library is located inside the `systeminfo/` subdirectory.

Therefore, its release tags must include the module directory prefix.

For example:

```text
systeminfo/v1.0.0
systeminfo/v1.1.0
```

Why?

Go needs to identify which module inside the repository the version tag belongs to.

### 10.3 Create the first release

From the Git repository:

```bash
git tag systeminfo/v1.0.0
```

Verify:

```bash
git tag -l 'systeminfo/*'
```

Push the tag:

```bash
git push origin systeminfo/v1.0.0
```

The library is now published as version `v1.0.0`.

### Important concepts

**Git commit:** Records source code changes.

**Git tag:** Identifies a particular commit as a release.

**Go module version:** Allows consumers to request a specific library release.

A Git tag does not itself upload a compiled library binary. Go can retrieve the module source from the repository using its module version.

---

## 11. Consuming a Remote Go Module

After publishing `v1.0.0`, we changed the consumer application to use the remote release.

### 11.1 Remove the local replace directive

Navigate to the consumer module:

```bash
cd system-info-cli
```

Remove the local replacement:

```bash
go mod edit -dropreplace=github.com/sahu-ayush/go-library-poc/systeminfo
```

### 11.2 Require the released version

```bash
go get github.com/sahu-ayush/go-library-poc/systeminfo@v1.0.0
```

### 11.3 Verify dependency resolution

```bash
go list -m all
```

Expected output:

```text
github.com/ayush/system-info-cli
github.com/sahu-ayush/go-library-poc/systeminfo v1.0.0
```

### 11.4 Download the module explicitly

```bash
go mod download github.com/sahu-ayush/go-library-poc/systeminfo@v1.0.0
```

### 11.5 Verify module location

```bash
go env GOMODCACHE
```

Example output:

```text
/Users/ayushsahu/go/pkg/mod
```

Find the downloaded library:

```bash
go list -f '{{.Dir}}' -m github.com/sahu-ayush/go-library-poc/systeminfo
```

Example output:

```text
/Users/ayushsahu/go/pkg/mod/github.com/sahu-ayush/go-library-poc/systeminfo@v1.0.0
```

This confirms that Go is using the downloaded module rather than the local source directory.

### 11.6 Verify the application

```bash
go mod tidy
go test ./...
go run .
```

The CLI successfully executed using the remotely downloaded `v1.0.0` library.

---

## 12. Module Version Upgrade

The next objective was to understand how library consumers receive new functionality.

### 12.1 Add a new function

We added:

```go
func GetSystemSummary() string {
    return runtime.GOOS + "/" + runtime.GOARCH
}
```

We also added a corresponding unit test.

Run:

```bash
go test -v -count=1 ./...
```

All six tests passed.

### 12.2 Commit the library changes

```bash
git add info.go info_test.go

git commit -m "Add system summary function"

git push origin main
```

### 12.3 Create a new release

```bash
git tag systeminfo/v1.1.0
```

Verify:

```bash
git tag -l 'systeminfo/*'
```

Expected:

```text
systeminfo/v1.0.0
systeminfo/v1.1.0
```

Push the new tag:

```bash
git push origin systeminfo/v1.1.0
```

The new library release is now available.

### 12.4 Upgrade the consumer application

Navigate to the CLI module:

```bash
cd system-info-cli
```

Check the currently installed version:

```bash
go list -m github.com/sahu-ayush/go-library-poc/systeminfo
```

Before upgrading:

```text
github.com/sahu-ayush/go-library-poc/systeminfo v1.0.0
```

Upgrade to the new release:

```bash
go get github.com/sahu-ayush/go-library-poc/systeminfo@v1.1.0
```

Verify:

```bash
go list -m github.com/sahu-ayush/go-library-poc/systeminfo
```

After upgrading:

```text
github.com/sahu-ayush/go-library-poc/systeminfo v1.1.0
```

### 12.5 Consume the new function

Add this to the CLI:

```go
fmt.Println("System Summary:", systeminfo.GetSystemSummary())
```

Run:

```bash
gofmt -w main.go
go test ./...
go run .
```

Expected output:

```text
System Summary: darwin/arm64
```

The new functionality is now available to the consumer application.

### What did we prove?

The consumer application does not automatically receive new functionality just because the library publisher released a new version.

The consumer explicitly upgrades its dependency.

This is an important concept in dependency management and release engineering.

---

## 13. Important Go Commands

| Command | Purpose |
|---|---|
| `go version` | Displays installed Go version |
| `go env GOMODCACHE` | Displays module cache location |
| `go mod init` | Initializes a Go module |
| `go mod edit` | Edits module configuration |
| `go get` | Adds or upgrades dependencies |
| `go mod download` | Downloads module dependencies |
| `go mod tidy` | Synchronizes dependency requirements and checksums |
| `go mod verify` | Verifies downloaded module cache content |
| `go list -m all` | Lists modules in the dependency graph |
| `go test ./...` | Runs tests in all packages |
| `go test -v -count=1 ./...` | Runs tests with verbose output without reusing cached results |
| `go run .` | Builds and executes the current application |
| `go build ./...` | Builds packages in the current module |
| `go list -f '{{.Dir}}' -m MODULE` | Displays the resolved module directory |
| `gofmt -w .` | Formats Go source files recursively |
| `git tag` | Creates or lists Git tags |
| `git push origin TAG` | Pushes a release tag |

---

## 14. Understanding go.mod and go.sum

### go.mod

The `go.mod` file defines the module identity and its dependency requirements.

Example:

```go
module github.com/ayush/system-info-cli

go 1.25.5

require github.com/sahu-ayush/go-library-poc/systeminfo v1.1.0
```

Important directives:

- `module`: Identifies the module.
- `go`: Specifies the Go language version.
- `require`: Declares required dependencies.
- `replace`: Replaces a module with another module location or version.

### go.sum

The `go.sum` file records cryptographic checksums associated with downloaded module content and module `go.mod` files.

Example:

```text
github.com/sahu-ayush/go-library-poc/systeminfo v1.1.0 h1:...
github.com/sahu-ayush/go-library-poc/systeminfo v1.1.0/go.mod h1:...
```

There are two different checksums:

| Entry | Purpose |
|---|---|
| `module v1.1.0` | Verifies module source content |
| `module v1.1.0/go.mod` | Verifies the module's go.mod file |

### Key differences

| go.mod | go.sum |
|---|---|
| Declares module requirements | Records verification checksums |
| Defines dependency versions | Helps verify downloaded dependency content |
| Required for a module | Generated and maintained as needed |

Both files should generally be committed to Git for reproducible dependency management.

---

## 15. Git Tags and Go Module Versions

This POC introduced two library releases:

| Release | Functionality |
|---|---|
| `v1.0.0` | Initial system information library |
| `v1.1.0` | Added `GetSystemSummary()` |

Corresponding Git tags:

```text
systeminfo/v1.0.0
systeminfo/v1.1.0
```

### Semantic Versioning

Go modules commonly follow Semantic Versioning:

```text
MAJOR.MINOR.PATCH
```

| Component | Meaning | Example |
|---|---|---|
| MAJOR | Breaking changes | v2.0.0 |
| MINOR | Backward-compatible functionality | v1.1.0 |
| PATCH | Backward-compatible bug fixes | v1.1.1 |

Examples:

- `v1.0.0` → Initial stable release.
- `v1.1.0` → New backward-compatible functionality.
- `v1.1.1` → Bug fix.
- `v2.0.0` → Breaking API changes.

**Note:** Go has special module path requirements for major versions 2 and above. A module typically uses `/v2` in its module path when releasing version 2 or higher.

---

## 16. Key Learnings

The most important concepts from this POC:

1. A package organizes Go source files.
2. A module defines a versioned dependency boundary.
3. A reusable library exposes functionality through exported identifiers.
4. A consumer application imports and calls library APIs.
5. One Git repository can contain multiple independent Go modules.
6. The `replace` directive supports local module development.
7. Git tags identify specific library releases.
8. Nested modules require appropriately prefixed release tags.
9. `go.mod` declares dependency versions.
10. `go.sum` helps verify dependency integrity.
11. Go stores downloaded dependencies in its module cache.
12. Consumers control when they upgrade library versions.

### Connection to DevOps and Platform Engineering

These concepts are directly applicable to real-world engineering workflows.

For example, an organization may maintain reusable Go modules for:

- Common logging and observability.
- Cloud infrastructure utilities.
- Authentication