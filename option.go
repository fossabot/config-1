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

type option struct {
	filename string
}

// Option config option.
type Option func(*option)

// WithFilename assign YAML filename to load.
func WithFilename(filename string) Option {
	return func(o *option) {
		o.filename = filename
	}
}

func parseOption(opts []Option) *option {
	opt := &option{}
	for _, fn := range opts {
		fn(opt)
	}

	return opt
}
