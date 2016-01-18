// Copyright 2016 Stefan Böhmann.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package envconf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepareKey(t *testing.T) {
	assert := assert.New(t)

	assert.Empty(prepareKey(""))
	assert.NotEmpty(prepareKey("ID"))

	assert.Equal("ID", prepareKey("ID"))
	assert.Equal("ID", prepareKey("id"))
	assert.Equal("ID", prepareKey("id "))
	assert.Equal("ID", prepareKey("  id "))

	assert.Equal("FOO_BAR", prepareKey("FOO_BAR"))
	assert.Equal("FOO_BAR", prepareKey("foo_bar"))
	assert.Equal("FOO_BAR", prepareKey("foo_bar "))
	assert.Equal("FOO_BAR", prepareKey("  foo_bar "))
	assert.Equal("FOO_BAR", prepareKey("Foo Bar"))
	assert.Equal("FOO_BAR", prepareKey("  FOO  BAR "))

	assert.NotEqual("Max Mustermann", prepareKey("Max Mustermann"))
}

func TestParseBool(t *testing.T) {
	assert := assert.New(t)

	for _, k := range []string{"", "1", "y", "Y", "true", "True", "TRUE", "yes", "Yes", "YES", "On", "on", "ON"} {
		value, okay := parseBool(k)
		assert.True(value, "Boolean value of \""+k+"\" is false but should be true")
		assert.True(okay, "Parsing \""+k+"\" as boolean failed")
	}

	for _, k := range []string{"0", "n", "N", "false", "False", "FALSE", "no", "No", "NO", "Off", "off", "OFF"} {
		value, okay := parseBool(k)
		assert.False(value, "Boolean value of \""+k+"\" is true but should be false")
		assert.True(okay, "Parsing \""+k+"\" as boolean failed")
	}
}

func TestPrefix(t *testing.T) {
	assert := assert.New(t)

	assert.Empty(GetPrefix())

	SetPrefix("foo")
	assert.NotEmpty(GetPrefix())

	SetPrefix("")
	assert.Empty(GetPrefix())
}
