# golog Package

## Overview
This is a simple Go logging package that provides different log levels (DEBUG, INFO, WARN, ERROR). It allows you to set the log level at runtime and logs messages accordingly.

## Installation
To install the package, run:

```sh
go get github.com/luytbq/golog
```

## Usage

### Import the Package
```go
import "github.com/luytbq/golog"
```

### Set Log Level
Before logging messages, set the desired log level:
```go
golog.SetLogLevel(golog.DEBUG)
```

### Logging Messages
```go
golog.Debug("This is a debug message")
golog.Info("This is an info message")
golog.Warn("This is a warning message")
golog.Error("This is an error message")
```

## Log Levels
The available log levels are:
- `golog.DEBUG`
- `golog.INFO`
- `golog.WARN`
- `golog.ERROR`

## Example
Here is a complete example of how to use the golog package:

```go
package main

import (
    "github.com/luytbq/golog"
)

func main() {
    golog.SetLogLevel(golog.DEBUG)
    
    golog.Debug("This is a debug message")
    golog.Info("This is an info message")
    golog.Warn("This is a warning message")
    golog.Error("This is an error message")
}
```

## License
This package is open-source and available under the MIT License.

