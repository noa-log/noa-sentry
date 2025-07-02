# Noa Sentry
A Sentry integration module for Noa Log that allows quick integration of Noa with Sentry. It automatically sends errors to Sentry when print errors.

## Installation
```bash
go get -u github.com/noa-log/noa-sentry
```

## Quick Start
```go
package main

import (
    "errors"
    "os"
    "time"

    "github.com/getsentry/sentry-go"
    "github.com/noa-log/noa"
    "github.com/noa-log/noa-sentry"
)

func main() {
    // Get Dsn from environment variable
    Dsn := os.Getenv("SENTRY_DSN")

    // Initialize Sentry
    err := sentry.Init(sentry.ClientOptions{
        Dsn:   Dsn,
        Debug: true,
    })
    if err != nil {
        panic(err)
    }
    defer sentry.Flush(2 * time.Second)

    // Create a new logger instance
    logger := noa.NewLog()

    // Register the Sentry Handler
    logger.AddAfterHandle(noasentry.CaptureHandler)

    // Log an error
    err = errors.New("An example error")
    logger.Error("Test", err)
}
```

## License
This project is open-sourced under the [Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0). Please comply with the terms when using it.