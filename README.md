# Go Dotenv

## Description:

**Go Dotenv** is a lightweight Go package designed to simplify the configuration of your Go applications by loading environment variables from a `.env` file (or any file of your choice). It provides a flexible structure, allowing you to customize the configuration process based on your application's needs.

**Note:** The default `env` file is a ".env" file at the root of the project and the default `logger` writes to `os.std.Out`

## Features:
- **Simple Configuration:** Easily load environment variables from a specified .env file or use default values.
- **Logger Integration:** Seamlessly integrate with your existing logging infrastructure by accepting a `*log.Logger` as a parameter.
- **Error Handling:** Provides clear error messages for quick debugging.
- **Customization:** Customize the configuration process by specifying a filename, logger, or both, depending on your application requirements.

## Installation
```bash
go get -u github.com/tonie-ng/go-dotenv
```

## Usage
There are various ways to use this package.
1. Using the default values for the filename and logger. Please ensure to have a `.env` file at the root directory of your project

```go
package main

import (
	"fmt"
	"os"

	"github.com/tonie-ng/go-dotenv"
)

func main() {
    dotenv.Config()

    // This is equivalent to _, _ := dotenv.Config()

    envVariable := os.Getenv("MY_ENV_KEY")
    /*
    "MY_ENV_KEY" represents an env eky that was provided in the env file
    ie., MY_ENV_KEY=value
    */
    
    fmt.Println(envVariable)
    // This prints the value assigned to the MY_ENV_KEY
}
```
or
```go
package main

import (
	"fmt"

	"github.com/tonie-ng/go-dotenv"
)

func main() {

	envVars, err := dotenv.Config()

	/*
		dotenv.Config() returns a map of all the environment variables provided in the .env file and an err if any.
		These can be ignored by using the _ operator.
	*/
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

	"github.com/tonie-ng/go-dotenv"
)

func main() {

	dotenv.Config("env")
	/*
	We proviide a string here and this is assigned to the filename while the logger uses the dafault.
	To assign the logger to a value instead, we do:
	*/

	logger := log.New(os.Stdout, "Example logger", log.LstdFlags)

	dotenv.Config(logger)

	
	// You can also assign logger to nil by doing. This indicates that you dont want a logger.
	dotenv.Config(nil)

	fmt.Println(os.Getenv("MY_ENV_KEY"))
  // prints the value of the MY_ENV_KEY provided in the env file
}
```

3. Providing a logger and a filename.
```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tonie-ng/go-dotenv"
)

func main() {

	filename := "sample.env"

	logger := log.New(os.Stdout, "Example logger", log.LstdFlags)

	dotenv.Config(filename, logger)

	// You can also assign logger to nil by doing. This indicates that you dont want a logger.
	logger = nil
	dotenv.Config(filename, logger)

	fmt.Println(os.Getenv("MY_ENV_KEY"))
}
```

## Contibutions
Contributions to this project are welcome, plese refer to the [contibution](https://github.com/tonie-ng/go-dotenv/blob/main/CONTRIBUTING.md) guide.
