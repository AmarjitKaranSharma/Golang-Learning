# GoLang Project

This is a Go-based project using SQLite via `github.com/mattn/go-sqlite3` (which requires CGO). The project contains standard Go code patterns, including closures, and provides a clean development setup for Windows, macOS, and Linux.

---

## 🚀 Table of Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [Environment & CGO (Required for SQLite)](#environment--cgo-required-for-sqlite)
- [Common Commands](#common-commands)
- [Troubleshooting](#troubleshooting)

---

## Overview

This project demonstrates Go development using modular project structure (`cmd/*`), SQLite interaction, and CGO-enabled builds. Includes example patterns like closures and best practices for Windows development with MinGW or MSYS2.

---

## Prerequisites

- Go **1.20+**
- Git
- SQLite driver requiring CGO:
  - `github.com/mattn/go-sqlite3`
- **A C compiler (mandatory for SQLite):**
  - **Windows:** MinGW-w64 or MSYS2
  - **Linux/macOS:** gcc or clang

---

## Install Dependency

```
go mod tidy
```

## Run the application

### Windows (PowerShell)

```
$env:CGO_ENABLED = "1"
$env:CC = "gcc"   # or full path to MinGW gcc e.g. C:\msys64\mingw64\bin\gcc.exe

go run ./cmd/students-api/main.go --config ./config/local.yaml
```

### Linux/macOS

```
export CGO_ENABLED=1
export CC=gcc

go run ./cmd/students-api/main.go --config ./config/local.yaml
```

## Environment & CGO (Required for SQLite)

### If CGO is disabled, you will receive this error:

```
Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub
exit status 1
```

### ✔ How to enable CGO on Windows

#### 1. Install MSYS2 or MinGW-w64

#### 2. Ensure gcc is accessible:

```
where.exe gcc
```

#### 3. Set environment variables before build or run:

```
$env:CGO_ENABLED = "1"
$env:CC = "gcc"
```

### ✔ Building a production binary

```
CGO_ENABLED=1 CC=gcc go build -o bin/app ./cmd/students-api
```

## Common Commands

| Purpose          | Command                                                       |
| ---------------- | ------------------------------------------------------------- |
| Run app          | `go run ./cmd/students-api/main.go`                           |
| Build binary     | `CGO_ENABLED=1 CC=gcc go build -o bin/app ./cmd/students-api` |
| Tidy modules     | `go mod tidy`                                                 |
| Run tests        | `go test ./...`                                               |
| Check Go version | `go version`                                                  |


## Troubleshooting

### ❌ Binary was compiled with 'CGO_ENABLED=0'

#### Enable CGO and set compiler:

```
$env:CGO_ENABLED = "1"
$env:CC = "gcc"
```

### ❌ go : The term 'go' is not recognized

#### Add Go to PATH (Windows default):

```
C:\Go\bin
```

#### Restart terminal.

### ❌ gcc not found

#### Install using MSYS2:

```
pacman -S --needed base-devel mingw-w64-x86_64-toolchain
```
