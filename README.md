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
- [Closures in Go (Reference)](#closures-in-go-reference)
- [Project Structure](#project-structure)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

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

