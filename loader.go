package dotenv

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func load(filename string) (map[string]string, error) {
	var key, value string
	envVars := make(map[string]string)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, "=")

		if len(tokens) != 2 {
			continue
		} else if string(tokens[0][0]) == "#" || string(tokens[0][0]) == "//" {
			continue
		} else if string(tokens[0]) == "" || string(tokens[1]) == "" {
			continue
		} else {
			value = strings.TrimSpace(tokens[1])
			key = strings.TrimSpace(tokens[0])

			keysize := strings.Split(key, " ")
			valuesize := strings.Split(value, " ")

			if len(keysize) != 1 {
				continue
			}

			if len(valuesize) > 1 {
				if string(valuesize[1][0]) == "#" || string(valuesize[1][:2]) == "//" {
					value = valuesize[0]
				}
			}

			if value[0] == '"' && value[len(value)-1] == '"' {
				value = strings.Trim(value, "\"")
			}

			envVars[key] = value
		}
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return envVars, nil
}

func Config(params ...interface{}) (map[string]string, error) {
	var (
		reterr   error
		logger   *log.Logger
		filename string
	)

	if len(params) > 2 {
		return nil, fmt.Errorf("Config accepts up to two parameters ie Config(string, *log.Logger)")
	}

	if len(params) == 1 {
		if v, ok := params[0].(string); ok {
			filename = v
			log.New(os.Stderr, "Dotenv logger", log.LstdFlags)
		} else if v, ok := params[0].(*log.Logger); ok {
			logger = v
			filename = ".env"
		} else if params[0] == nil {
			logger = nil
			filename = ".env"
		} else {
			return nil, fmt.Errorf("Invalid parameters. Please use this format: Config(string)\n or Config(*log.Logger)\n or Config(nil)")
		}
	}

	if len(params) == 2 {
		if v, ok := params[0].(string); !ok {
			return nil, fmt.Errorf("%v represents a filename and must be a string", v)
		} else {
			filename = v
		}

		if params[1] == nil {
			logger = nil
		} else if v, ok := params[1].(*log.Logger); ok {
			logger = v
		} else {
			return nil, fmt.Errorf("%v represents a logger and must be of type *log.Logger", v)
		}
	}

	if len(params) == 0 {
		filename = ".env"
		logger = log.New(os.Stderr, "Dotenv logger", log.LstdFlags)
	}

	envVars, err := load(filename)
	if err != nil {
		if logger != nil {
			logger.Println(err)
		}
		return nil, err
	}

	for key, value := range envVars {
		err = os.Setenv(key, value)
		if err != nil {
			if logger != nil {
				logger.Println(err)
			}
			reterr = err
			break
		}
	}

	return envVars, reterr
}
