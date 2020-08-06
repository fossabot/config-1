# config

[![CircleCI](https://img.shields.io/circleci/build/github/pk60/config/master?token=1154f9ffcdc4c1bd95f320a7777f920a3e4ca94d)](https://circleci.com/gh/pk60/config)
[![Go Report Card](https://goreportcard.com/badge/github.com/pk60/config)](https://goreportcard.com/report/github.com/pk60/config)
[![codecov](https://codecov.io/gh/pk60/config/branch/master/graph/badge.svg)](https://codecov.io/gh/pk60/config)

> Yet another golang configuration library.

### Why?

I wanted to load environment variables and YAML file to my configuration struct that's it.

### Feature

- Load configuration from environment variables (using [caarlos0/env](https://github.com/caarlos0/env))
- Optionally load configuration from YAML file (using [go-yaml/yaml](https://github.com/go-yaml/yaml))

### Installation

Go version 1.14+

```sh
go get github.com/pk60/config
```

### Example

```go
package main

import (
    "log"

    "github.com/pk60/config"
)

type MyConfig struct {
    Port int `yaml:"port" env:"PORT" envDefault:"3000"`
}

func main()  {
    c := &MyConfig{}
    if err := config.Load(c, config.WithFilename("./config.yml")); err != nil {
        log.Fatalf("failed to load config: %v", err)
    }    
}
```

### Documentation

Read here: [pkg.go.dev](https://pkg.go.dev/github.com/pk60/config?tab=doc)

### License

[Apache-2.0 License](https://github.com/pk60/config/blob/master/LICENSE)
