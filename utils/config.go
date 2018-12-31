//          Specter  Copyright (C) 2018-2019  SpecterTeam
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package utils

import (
	"errors"
	"github.com/json-iterator/go"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Path     string
	Type     uint
	Defaults interface{}
}

const (
	YAML = iota
	JSON
)

var ErrUnexpectedType = errors.New("unexpected type")

func NewConfig(path string, Type uint, defaults interface{}) (c Config) {
	if Type != YAML && Type != JSON {
		DefaultLogger.Warn("Specified config type for the file", path, "is invalid, using YAML type instead.")
		Type = YAML
	}
	
	if !FileExists(path) {
		_, err := os.Create(path)
		PanicErr(err)

		if defaults != nil {
			c.Marshal(defaults)
		}
	}
	
	c.Type     = Type
	c.Path     = path
	c.Defaults = defaults
	
	return
}

func (c *Config) Unmarshal(out *interface{}) {
	bytes, err := ioutil.ReadFile(c.Path)
	PanicErr(err)

	switch c.Type {
	case YAML:
		PanicErr(yaml.Unmarshal(bytes, &out))
		break

	case JSON:
		PanicErr(jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(bytes, &out))
		break

	default:
		panic(ErrUnexpectedType)
	}
}

func (c *Config) Marshal(in interface{}) {
	var (
		bytes []byte
		err   error
	)

	switch c.Type {
	case YAML:
		bytes, err = yaml.Marshal(&in)
		PanicErr(err)
		break

	case JSON:
		bytes, err = jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(&in)
		PanicErr(err)
		break

	default:
		panic(ErrUnexpectedType)
	}

	PanicErr(ioutil.WriteFile(c.Path, bytes, 0777))
}