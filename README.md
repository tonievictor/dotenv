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
or
```go
package main

import (
	"fmt"

	"github.com/tonievictor/dotenv"
)

func main() {


    // dotenv.Config() returns a map of all the environment variables provided in the .env file and an err if any.
	envVars, err := dotenv.Config()

	if err != nil {
		fmt.Println(err)
	}
    
	for _, v := range envVars {
		fmt.Println(v)
	}
}
```

2. Providing a value for either the logger or the filename. (If one is provided the other is assigned to the default.)
```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tonievictor/dotenv"
)

func main() {

	
	// We provide a string here and this is assigned to the filename while the logger uses the dafault.
	dotenv.Config("env")


	// To assign the logger to a value instead, we do:
	logger := log.New(os.Stdout, "Example logger", log.LstdFlags)

	dotenv.Config(logger)

	
	// You can also assign logger to nil by doing. This indicates that you dont want a logger.
	dotenv.Config(nil)

    // prints the value of the MY_ENV_KEY provided in the env file
	fmt.Println(os.Getenv("MY_ENV_KEY"))
}
```

3. Providing a logger and a filename.
```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tonievictor/dotenv"
)

func main() {

	filename := "sample.env"

	logger := log.New(os.Stdout, "Example logger", log.LstdFlags)

	dotenv.Config(filename, logger)

	// You can also assign logger to nil by doing. This indicates that you don't want a logger.
	dotenv.Config(filename, nil)

	fmt.Println(os.Getenv("MY_ENV_KEY"))
}
```

## Contibutions
Contributions to this project are welcome, please refer to the [contribution](https://github.com/tonievictor/dotenv/blob/main/CONTRIBUTING.md) guide.

## Contributors
[Tonie Victor](https://tonie.me) - Author
