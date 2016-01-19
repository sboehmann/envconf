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

	SetPrefix("foo")
	assert.Equal("FOO_FOO_BAR", prepareKey("FOO_BAR"))

	SetPrefix("")
	assert.Equal("FOO_BAR", prepareKey("FOO_BAR"))
}

func TestParseBool(t *testing.T) {
	assert := assert.New(t)

	// Test true
	for _, k := range []string{"", "1", "y", "Y", "true", "True", "TRUE", "yes", "Yes", "YES", "On", "on", "ON"} {
		value, okay := parseBool(k)
		assert.True(value, "Boolean value of \""+k+"\" is false but should be true")
		assert.True(okay, "Parsing \""+k+"\" as boolean failed")
	}

	// Test false
	for _, k := range []string{"0", "n", "N", "false", "False", "FALSE", "no", "No", "NO", "Off", "off", "OFF"} {
		value, okay := parseBool(k)
		assert.False(value, "Boolean value of \""+k+"\" is true but should be false")
		assert.True(okay, "Parsing \""+k+"\" as boolean failed")
	}

	// Test invalid
	for _, k := range []string{"foo", "Disabled", "Enabled", "12", "-1", "2"} {
		value, okay := parseBool(k)
		assert.False(value, "Boolean value of \""+k+"\" is true but should be false")
		assert.False(okay, "Parsing \""+k+"\" as boolean should fail!")
	}
}

func TestPrefix(t *testing.T) {
	assert := assert.New(t)

	assert.Empty(GetPrefix())

	SetPrefix("foo")
	assert.NotEmpty(GetPrefix())
	assert.Equal("FOO_", GetPrefix())

	SetPrefix(" BAR  ")
	assert.NotEmpty(GetPrefix())
	assert.Equal("BAR_", GetPrefix())

	SetPrefix(" FOO  BAR  ")
	assert.NotEmpty(GetPrefix())
	assert.Equal("FOO_BAR_", GetPrefix())

	SetPrefix("")
	assert.Empty(GetPrefix())
	SetPrefix("   	")
	assert.Empty(GetPrefix())
}

func TestString(t *testing.T) {
	assert := assert.New(t)

	UnsetKey("envconf_test_0815")
	str, ok := GetString("envconf_test_0815")
	assert.False(ok)
	assert.Empty(str)

	SetString("envconf_test_0815", "Foo Bar 42")
	str, ok = GetString("envconf_test_0815")
	assert.True(ok)
	assert.NotEmpty(str)
	assert.Equal("Foo Bar 42", str)

	SetDefaultString("envconf_test_0815", "Bar Foo 98")
	str, ok = GetString("envconf_test_0815")
	assert.True(ok)
	assert.NotEmpty(str)
	assert.Equal("Foo Bar 42", str)

	UnsetKey("envconf_test_0815")
	str, ok = GetString("envconf_test_0815")
	assert.False(ok)
	assert.Empty(str)

	SetDefaultString("envconf_test_3345", "34")
	str, ok = GetString("envconf_test_3345")
	assert.True(ok)
	assert.NotEmpty(str)
	assert.Equal("34", str)

	SetString("envconf_test_3345", "23")
	str = MustGetString("envconf_test_3345")
	assert.NotEmpty(str)
	assert.Equal("23", str)

	SetString("envconf_test_3345", "")
	str, ok = GetString("envconf_test_3345")
	assert.True(ok)
	assert.Empty(str)
	assert.Equal("", str)

	assert.True(IssetKey("envconf_test_3345"))
	UnsetKey("envconf_test_3345")
	assert.False(IssetKey("envconf_test_3345"))

	assert.Panics(func() { MustGetString("envconf_test_3345") })
}
