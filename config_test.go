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

package config_test

import (
	"os"
	"testing"

	"github.com/pk60/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	type cfg struct {
		Val string
	}

	testCases := []struct {
		options   []config.Option
		assertErr assert.ErrorAssertionFunc
		want      interface{}
		name      string
	}{
		{
			[]config.Option{},
			assert.NoError,
			cfg{},
			"EmptyFile",
		},
		{
			[]config.Option{
				config.WithFilename("./non_existing_config.yml"),
			},
			assert.Error,
			cfg{},
			"FileNotExist",
		},
		{
			[]config.Option{
				config.WithFilename("./testdata/invalid.yml"),
			},
			assert.Error,
			cfg{},
			"InvalidFile",
		},
		{
			[]config.Option{
				config.WithFilename("./testdata/valid.yml"),
			},
			assert.NoError,
			cfg{
				Val: "value from yaml",
			},
			"ValidFile",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := cfg{}
			tc.assertErr(t, config.Load(&actual, tc.options...))
			assert.Equal(t, tc.want, actual)
		})
	}
}

func TestLoad_Precedence(t *testing.T) {
	_ = os.Setenv("KEY2", "env value 2")
	defer os.Unsetenv("KEY2")

	type cfg struct {
		Key0 string
		Key1 string
		Key2 string `env:"KEY2"`
		Key3 string `envDefault:"defaulted to value 3"`
	}

	actual := cfg{
		Key0: "value zero",
	}

	require.NoError(t, config.Load(&actual, config.WithFilename("./testdata/valid.yml")))

	want := cfg{
		Key0: "value zero",
		Key1: "yaml value 1",
		Key2: "env value 2",
		Key3: "defaulted to value 3",
	}

	assert.Equal(t, want, actual)
}
