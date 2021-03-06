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

import "os"


// FileExists checks if a file exists or not on the specified system path.
func FileExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

// PanicErr checks if the given error is a nil and if it's not, it panic it.
func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}