# Dotenv
A lightweight Go package designed to simplify the configuration of your Go applications by loading environment variables from a `.env` file (or any file of your choice)

**Note:** The default `env` file is a ".env" file at the root of the project and the default `logger` writes to `os.Stdout`

## Installation
```bash
go get -u github.com/tonievictor/dotenv
```

## Usage
There are various ways to use this package.
1. Using the default values for the filename and logger.

```go
package main

import (
	"fmt"
	"os"

	"github.com/tonievictor/dotenv"
)

func main() {
    
    // loads the env variables using the defaults while ignoring the return values.
    dotenv.Config()

    envVariable := os.Getenv("MY_ENV_KEY")
    
    // This prints the value assigned to the MY_ENV_KEY
    fmt.Println(envVariable)
}
```
2. You can also provide an env file and a logger.
```go
package main

import (
	"fmt"
	"os"

	"github.com/tonievictor/dotenv"
)

func main() {
    filename := dotenv.WithFilename("env file")
    logger := log.New(os.Stdout, "Example logger", log.LstdFlags)

    dotenv.Config(filename, dotenv.WithLogger(logger))
    // you can also choose to provide just the filename or the logger
    // and use the default for the other.
    // dotenv.Config(filename)
    // or
    // dotenv.Config(dotenv.WithLogger(logger))

    envVariable := os.Getenv("MY_ENV_KEY")
    
    // This prints the value assigned to the MY_ENV_KEY
    fmt.Println(envVariable)
}
```
> Note: `dotenv.Config()` returns a map of the environment variables and an error

## Contributors
[Tonie Victor](https://tonie.me) - Author
