// Copyright The Wuenak Contributors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config_test

import (
	"fmt"
	"os"

	"github.com/pk60/config"
)

func ExampleLoad() {
	_ = os.Setenv("KEY_2", "from env")
	defer os.Unsetenv("KEY_2")

	type conf struct {
		Val  string
		Key1 string
		Key2 string `env:"KEY_2"`
		Key3 string `envDefault:"default value"`
	}

	c := &conf{
		Val: "value",
	}

	_ = config.Load(c, config.WithFilename("./testdata/valid.yml"))
	fmt.Printf("%#v", c)
	// Output:
	// &config_test.conf{Val:"value from yaml", Key1:"yaml value 1", Key2:"from env", Key3:"default value"}
}

func ExampleLoad_readYAMLError() {
	type conf struct{}
	c := &conf{}

	err := config.Load(c, config.WithFilename("./non_existing_config.yml"))
	fmt.Print(err.Error())
	// Output:
	// failed to read YAML file: open ./non_existing_config.yml: no such file or directory
}

func ExampleLoad_parseYAMLError() {
	type conf struct{}
	c := &conf{}

	err := config.Load(c, config.WithFilename("./testdata/invalid.yml"))
	fmt.Print(err.Error())
	// Output:
	// failed to unmarshal YAML file: ./testdata/invalid.yml: yaml: line 2: mapping values are not allowed in this context
}

func ExampleLoad_parseEnvError() {
	_ = os.Setenv("INVALID_DATA", "{1+2+3}")
	defer os.Unsetenv("INVALID_DATA")

	type conf struct {
		InvalidData *[]string `env:"INVALID_DATA"`
	}

	c := &conf{}

	err := config.Load(c)
	fmt.Printf("%v", err)
	// Output:
	// failed to parse environment variables: env: no parser found for field "InvalidData" of type "*[]string"
}
