# Go Run [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/micro/go-run?status.svg)](https://godoc.org/github.com/micro/go-run) [![Travis CI](https://api.travis-ci.org/micro/go-run.svg?branch=master)](https://travis-ci.org/micro/go-run) [![Go Report Card](https://goreportcard.com/badge/micro/go-run)](https://goreportcard.com/report/github.com/micro/go-run)

Go run manages the lifecycle of a service from source to process.

## Overview

A runtime fetches source code, builds a binary and executes it.

```
type Runtime interface {
	// Fetch source from url
	Fetch(url string, opts ...FetchOption) (*Source, error)
	// Build the binary from source
	Build(*Source) (*Binary, error)
	// Execute a binary
	Exec(*Binary) (*Process, error)
	// Kill a process
	Kill(*Process) error
	// Wait for a process to exit
	Wait(*Process) error
}
```

## Languages

- Go

