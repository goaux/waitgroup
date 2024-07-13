# waitgroup
Package waitgroup provides a wrapper around sync.WaitGroup to simplify launching goroutines and waiting for their completion.

[![Go Reference](https://pkg.go.dev/badge/github.com/goaux/waitgroup.svg)](https://pkg.go.dev/github.com/goaux/waitgroup)
[![Go Report Card](https://goreportcard.com/badge/github.com/goaux/waitgroup)](https://goreportcard.com/report/github.com/goaux/waitgroup)

## Features

- Simple wrapper around `sync.WaitGroup`
- Handle goroutine launch and completion waiting in a single method
- Keep your concurrent code cleaner

## Installation

Install the package with:

    go get github.com/goaux/waitgroup

## Usage

Here's a basic example of how to use the package:

    package main

    import (
        "fmt"
        "time"

        "github.com/goaux/waitgroup"
    )

    func main() {
        var sy waitgroup.Sync
        results := make(chan string, 3)

        sy.Go(func() {
            time.Sleep(100 * time.Millisecond)
            results <- "First task done"
        })

        sy.Go(func() {
            time.Sleep(200 * time.Millisecond)
            results <- "Second task done"
        })

        sy.Go(func() {
            time.Sleep(50 * time.Millisecond)
            results <- "Third task done"
        })

        // Wait for all goroutines to complete
        sy.Wait()
        close(results)

        // Print results in the order they were added to the channel
        for result := range results {
            fmt.Println(result)
        }
    }
