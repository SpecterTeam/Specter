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

package server

import (
	"bufio"
	"errors"
	"github.com/SpecterTeam/Specter/utils"
	"os"
	"strconv"
)

type (
	Server struct {
		Running    bool
		Path       string
		Logger     utils.Logger
		Debug      bool
		Config     utils.Config
		Parameters ParametersConfig
	}

	ParametersConfig struct {
		Name       string
		MaxPlayers uint
		Port       int
	}
)


var (
	ErrServerAlreadyRunning = errors.New(
		"server is already running, you can't start it twice on the same instance")
)

func NewServer() *Server {
	return new(Server)
}

func (s *Server) Start(path string, interactive, debug bool) {
	if s.Running == true {
		panic(ErrServerAlreadyRunning)
	}

	s.Running = true
	s.Path = path
	s.Logger = *utils.NewLogger("[Specter]")

	if !utils.FileExists(s.Path + "server.yml") {
		if interactive {
			s.Config = utils.NewConfig(s.Path + "server.yml", utils.YAML, nil)
			s.Logger.Debug(s.Debug,
				"Couldn't find the server configuration file. Please follow the instructions below to run your server.",
				)

			s.Logger.Info("What do you want to name your server (Enter a name)")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			s.Parameters.Name = scanner.Text()

			s.Logger.Info("What should be the maximum number of players allowed (Enter a number)")
			scanner.Scan()
			number, err := strconv.Atoi(scanner.Text())
			utils.PanicErr(err)
			s.Parameters.MaxPlayers = uint(number)

			s.Logger.Info("What should be the maximum number of players allowed (Enter a number)")
			scanner.Scan()
			number, err = strconv.Atoi(scanner.Text())
			utils.PanicErr(err)
			s.Parameters.Port = number
		} else {
			s.Logger.Warn(
				"Couldn't find the server configuration file. Creating a new one with default parameters.")
			s.Config = utils.NewConfig(s.Path + "server.yml", utils.YAML, ParametersConfig{
				Name: "This server is running on Specter.",
				MaxPlayers: 1000,
				Port: 19132,
			})
		}
	}

	s.loadParameters()


}

func (s *Server) loadParameters() {
	var out interface{}
	s.Config.Unmarshal(&out)
	s.Parameters = out.(ParametersConfig)
}
