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
	"fmt"
	"time"
)

var DefaultLogger = NewLogger("[Specter]")

type Logger struct {
	Prefix string
}

func NewLogger(prefix string) *Logger {
	return &Logger{Prefix: prefix}
}

func (l *Logger) Info(text ...interface{}) {
	fmt.Println(l.Prefix, "[" + time.Now().String() + "]", "[Info]", fmt.Sprint(text))
}

func (l *Logger) Warn(text ...interface{}) {
	fmt.Println(l.Prefix, "[" + time.Now().String() + "]", "[WARNING]", fmt.Sprint(text))
}

func (l *Logger) Fatal(text ...interface{}) {
	fmt.Println(l.Prefix, "[" + time.Now().String() + "]", "[FATAL]", fmt.Sprint(text))
}

func (l *Logger) Debug(active bool, text ...interface{}) {
	if active {
		fmt.Println(l.Prefix, "[" + time.Now().String() + "]", "[Debug]", fmt.Sprint(text))
	}
}
