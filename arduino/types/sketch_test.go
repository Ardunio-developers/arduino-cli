// This file is part of arduino-cli.
//
// Copyright 2019 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to modify or
// otherwise use the software for commercial activities involving the Arduino
// software without disclosing the source code of your own applications. To purchase
// a commercial license, send an email to license@arduino.cc.

package types_test

import (
	"testing"

	"github.com/arduino/arduino-cli/arduino/types"
	paths "github.com/arduino/go-paths-helper"
	"github.com/stretchr/testify/assert"
)

func TestSketchFileSortByName(t *testing.T) {
	sketchSort := types.SketchFileSortByName{
		types.SketchFile{Name: paths.New("foo")},
		types.SketchFile{Name: paths.New("bar")},
	}

	assert.Len(t, sketchSort, 2)
	assert.False(t, sketchSort.Less(0, 1))
	assert.True(t, sketchSort.Less(1, 0))

	sketchSort.Swap(0, 1)
	assert.True(t, sketchSort.Less(0, 1))
	assert.False(t, sketchSort.Less(1, 0))
}
