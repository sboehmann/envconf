// Copyright 2016 Stefan Böhmann.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package envconf

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const EPSILON = float64(0.00000001)

func floatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}

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

	assert.False(IssetKey("envconf_test_0815"))
	SetString("envconf_test_0815", "")
	assert.True(IssetKey("envconf_test_0815"))
	assert.Equal("", MustGetString("envconf_test_0815"))

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

func TestBool(t *testing.T) {
	assert := assert.New(t)

	UnsetKey("envconf_test_bool_1")
	str, ok := GetString("envconf_test_bool_1")
	assert.False(ok)
	assert.Empty(str)

	UnsetKey("envconf_test_bool_2")
	str, ok = GetString("envconf_test_bool_2")
	assert.False(ok)
	assert.Empty(str)

	// Not existing environment variable equals false
	v, ok := GetBool("envconf_test_bool_1")
	assert.False(v)
	assert.True(ok)
	assert.False(MustGetBool("envconf_test_bool_1"))

	// A empty variable is considered to be true
	SetString("envconf_test_bool_1", "")
	v, ok = GetBool("envconf_test_bool_1")
	assert.True(v)
	assert.True(ok)
	assert.True(MustGetBool("envconf_test_bool_1"))

	SetString("envconf_test_bool_1", "1")
	v, ok = GetBool("envconf_test_bool_1")
	assert.True(v)
	assert.True(ok)
	assert.True(MustGetBool("envconf_test_bool_1"))

	SetString("envconf_test_bool_1", "0")
	v, ok = GetBool("envconf_test_bool_1")
	assert.False(v)
	assert.True(ok)
	assert.False(MustGetBool("envconf_test_bool_1"))

	SetString("envconf_test_bool_1", "blah")
	v, ok = GetBool("envconf_test_bool_1")
	assert.False(v)
	assert.False(ok)
	assert.Panics(func() { MustGetBool("envconf_test_bool_1") })

	UnsetKey("envconf_test_bool_1")
	UnsetKey("envconf_test_bool_2")

	SetBool("envconf_test_bool_1", true)
	assert.True(MustGetBool("envconf_test_bool_1"))
	assert.False(MustGetBool("envconf_test_bool_2"))

	SetDefaultBool("envconf_test_bool_1", false)
	assert.True(MustGetBool("envconf_test_bool_1"))
	assert.False(MustGetBool("envconf_test_bool_2"))

	SetDefaultBool("envconf_test_bool_2", true)
	assert.True(MustGetBool("envconf_test_bool_1"))
	assert.True(MustGetBool("envconf_test_bool_2"))

	SetBool("envconf_test_bool_2", false)
	SetDefaultBool("envconf_test_bool_1", false)
	SetDefaultBool("envconf_test_bool_2", false)
	assert.True(MustGetBool("envconf_test_bool_1"))
	assert.False(MustGetBool("envconf_test_bool_2"))
}

func TestDuration(t *testing.T) {
	assert := assert.New(t)

	UnsetKey("envconf_test1")
	UnsetKey("envconf_test2")

	v, ok := GetDuration("envconf_test1")
	assert.False(ok)
	assert.Zero(v)
	assert.Panics(func() { MustGetDuration("envconf_test1") })

	SetString("envconf_test1", "blahBlah")
	v, ok = GetDuration("envconf_test1")
	assert.False(ok)
	assert.Zero(v)
	assert.Panics(func() { MustGetDuration("envconf_test1") })

	SetString("envconf_test1", "15s")
	v, ok = GetDuration("envconf_test1")
	assert.True(ok)
	assert.Equal(15*time.Second, v)
	assert.Equal(15*time.Second, MustGetDuration("envconf_test1"))

	SetDefaultDuration("envconf_test1", 12*time.Millisecond)
	v, ok = GetDuration("envconf_test1")
	assert.True(ok)
	assert.Equal(15*time.Second, v)
	assert.Equal(15*time.Second, MustGetDuration("envconf_test1"))

	SetDefaultDuration("envconf_test2", 33*time.Millisecond)
	v, ok = GetDuration("envconf_test2")
	assert.True(ok)
	assert.Equal(33*time.Millisecond, v)
	assert.Equal(33*time.Millisecond, MustGetDuration("envconf_test2"))
}

func TestFloat64(t *testing.T) {
	assert := assert.New(t)

	UnsetKey("envconf_test1")
	UnsetKey("envconf_test2")

	v, ok := GetFloat64("envconf_test1")
	assert.False(ok)
	assert.Zero(v)
	assert.Panics(func() { MustGetFloat64("envconf_test1") })

	SetString("envconf_test1", "blahBlah")
	v, ok = GetFloat64("envconf_test1")
	assert.False(ok)
	assert.Zero(v)
	assert.Panics(func() { MustGetFloat64("envconf_test1") })

	SetString("envconf_test1", "83.3")
	v, ok = GetFloat64("envconf_test1")
	assert.True(ok)
	assert.True(floatEquals(83.3, v))
	assert.True(floatEquals(83.3, MustGetFloat64("envconf_test1")))

	SetDefaultFloat64("envconf_test1", -434.43202)
	v, ok = GetFloat64("envconf_test1")
	assert.True(ok)
	assert.True(floatEquals(83.3, v))
	assert.True(floatEquals(83.3, MustGetFloat64("envconf_test1")))

	SetDefaultFloat64("envconf_test2", -0.3424562)
	v, ok = GetFloat64("envconf_test2")
	assert.True(ok)
	assert.True(floatEquals(-0.3424562, v))
	assert.True(floatEquals(-0.3424562, MustGetFloat64("envconf_test2")))
}

func TestInt(t *testing.T) {
	assert := assert.New(t)

	UnsetKey("envconf_test1")
	UnsetKey("envconf_test2")

	v, ok := GetInt("envconf_test1")
	assert.False(ok)
	assert.Zero(v)
	assert.Panics(func() { MustGetInt("envconf_test1") })

	SetString("envconf_test1", "blahBlah")
	v, ok = GetInt("envconf_test1")
	assert.False(ok)
	assert.Zero(v)
	assert.Panics(func() { MustGetInt("envconf_test1") })

	SetString("envconf_test1", "42")
	v, ok = GetInt("envconf_test1")
	assert.True(ok)
	assert.Equal(42, v)
	assert.Equal(42, MustGetInt("envconf_test1"))

	SetDefaultInt("envconf_test1", -5)
	v, ok = GetInt("envconf_test1")
	assert.True(ok)
	assert.Equal(42, v)
	assert.Equal(42, MustGetInt("envconf_test1"))

	SetDefaultInt("envconf_test2", -33)
	v, ok = GetInt("envconf_test2")
	assert.True(ok)
	assert.Equal(-33, v)
	assert.Equal(-33, MustGetInt("envconf_test2"))
}

func TestInt64(t *testing.T) {
	assert := assert.New(t)

	UnsetKey("envconf_test1")
	UnsetKey("envconf_test2")

	v, ok := GetInt64("envconf_test1")
	assert.False(ok)
	assert.Zero(v)
	assert.Panics(func() { MustGetInt64("envconf_test1") })

	SetString("envconf_test1", "blahBlah")
	v, ok = GetInt64("envconf_test1")
	assert.False(ok)
	assert.Zero(v)
	assert.Panics(func() { MustGetInt64("envconf_test1") })

	SetString("envconf_test1", "42")
	v, ok = GetInt64("envconf_test1")
	assert.True(ok)
	assert.Equal(int64(42), v)
	assert.Equal(int64(42), MustGetInt64("envconf_test1"))

	SetDefaultInt64("envconf_test1", int64(-5))
	v, ok = GetInt64("envconf_test1")
	assert.True(ok)
	assert.Equal(int64(42), v)
	assert.Equal(int64(42), MustGetInt64("envconf_test1"))

	SetDefaultInt64("envconf_test2", int64(-33))
	v, ok = GetInt64("envconf_test2")
	assert.True(ok)
	assert.Equal(int64(-33), v)
	assert.Equal(int64(-33), MustGetInt64("envconf_test2"))

	SetInt64("envconf_test2", int64(-9223372036854775807))
	v, ok = GetInt64("envconf_test2")
	assert.True(ok)
	assert.Equal(int64(-9223372036854775807), v)
	assert.Equal(int64(-9223372036854775807), MustGetInt64("envconf_test2"))
}

func TestUInt(t *testing.T) {
	assert := assert.New(t)

	UnsetKey("envconf_test1")
	UnsetKey("envconf_test2")

	v, ok := GetUInt("envconf_test1")
	assert.False(ok)
	assert.Zero(v)
	assert.Panics(func() { MustGetUInt("envconf_test1") })

	SetString("envconf_test1", "blahBlah")
	v, ok = GetUInt("envconf_test1")
	assert.False(ok)
	assert.Zero(v)
	assert.Panics(func() { MustGetUInt("envconf_test1") })

	SetString("envconf_test1", "42")
	v, ok = GetUInt("envconf_test1")
	assert.True(ok)
	assert.Equal(uint(42), v)
	assert.Equal(uint(42), MustGetUInt("envconf_test1"))

	SetDefaultUInt("envconf_test1", uint(5))
	v, ok = GetUInt("envconf_test1")
	assert.True(ok)
	assert.Equal(uint(42), v)
	assert.Equal(uint(42), MustGetUInt("envconf_test1"))

	SetDefaultUInt("envconf_test2", uint(33))
	v, ok = GetUInt("envconf_test2")
	assert.True(ok)
	assert.Equal(uint(33), v)
	assert.Equal(uint(33), MustGetUInt("envconf_test2"))

	SetInt("envconf_test2", -4)
	v, ok = GetUInt("envconf_test2")
	assert.False(ok)
	assert.Equal(uint(0), v)
	assert.Panics(func() { MustGetUInt("envconf_test2") })
}

func TestUInt64(t *testing.T) {
	assert := assert.New(t)

	UnsetKey("envconf_test1")
	UnsetKey("envconf_test2")

	v, ok := GetUInt64("envconf_test1")
	assert.False(ok)
	assert.Zero(v)
	assert.Panics(func() { MustGetUInt64("envconf_test1") })

	SetString("envconf_test1", "blahBlah")
	v, ok = GetUInt64("envconf_test1")
	assert.False(ok)
	assert.Zero(v)
	assert.Panics(func() { MustGetUInt64("envconf_test1") })

	SetString("envconf_test1", "42")
	v, ok = GetUInt64("envconf_test1")
	assert.True(ok)
	assert.Equal(uint64(42), v)
	assert.Equal(uint64(42), MustGetUInt64("envconf_test1"))

	SetDefaultUInt64("envconf_test1", uint64(5))
	v, ok = GetUInt64("envconf_test1")
	assert.True(ok)
	assert.Equal(uint64(42), v)
	assert.Equal(uint64(42), MustGetUInt64("envconf_test1"))

	SetDefaultUInt64("envconf_test2", uint64(9223372036854775809))
	v, ok = GetUInt64("envconf_test2")
	assert.True(ok)
	assert.Equal(uint64(9223372036854775809), v)
	assert.Equal(uint64(9223372036854775809), MustGetUInt64("envconf_test2"))

	SetInt("envconf_test2", -4)
	v, ok = GetUInt64("envconf_test2")
	assert.False(ok)
	assert.Equal(uint64(0), v)
	assert.Panics(func() { MustGetUInt64("envconf_test2") })
}
