// Copyright The pk60 Contributors.
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

package config

import (
	"io/ioutil"

	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Load v with data from environment variables or YAML file (if defined).
func Load(v interface{}, opts ...Option) error {
	opt := parseOption(opts)

	if opt.filename != "" {
		in, err := ioutil.ReadFile(opt.filename)
		if err != nil {
			return errors.Wrap(err, "failed to read YAML file")
		}

		if err := yaml.Unmarshal(in, v); err != nil {
			return errors.Wrapf(err, "failed to unmarshal YAML file: %s", opt.filename)
		}
	}

	if err := env.Parse(v); err != nil {
		return errors.Wrap(err, "failed to parse environment variables")
	}

	return nil
}
