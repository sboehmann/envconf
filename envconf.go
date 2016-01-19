// Copyright 2016 Stefan Böhmann.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package envconf

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var prefix = ""

func prepareKey(key string) string {
	key = strings.TrimSpace(strings.ToUpper(key))
	for strings.Contains(key, "  ") {
		key = strings.Replace(key, "  ", " ", -1)
	}
	key = strings.Replace(key, " ", "_", -1)

	if key != "" && prefix != "" {
		key = prefix + key
	}

	return key
}

func parseBool(str string) (value bool, ok bool) {
	switch strings.TrimSpace(strings.ToLower(str)) {
	case "1", "y", "true", "yes", "on", "":
		return true, true
	case "0", "n", "false", "no", "off":
		return false, true
	}

	return false, false
}

// GetPrefix ...
func GetPrefix() string {
	return prefix
}

// SetPrefix ...
func SetPrefix(p string) {
	p = strings.TrimSpace(strings.ToUpper(p))
	if p != "" {
		for strings.Contains(p, "  ") {
			p = strings.Replace(p, "  ", " ", -1)
		}
		p = strings.Replace(p, " ", "_", -1)

		if !strings.HasSuffix(p, "_") {
			p += "_"
		}
	}

	prefix = p
}

// UnsetKey ...
func UnsetKey(key string) {
	os.Unsetenv(prepareKey(key))
}

// IssetKey ...
func IssetKey(key string) bool {
	_, ok := os.LookupEnv(prepareKey(key))
	return ok
}

// SetDefaultString ...
func SetDefaultString(key string, value string) {
	if !IssetKey(key) {
		SetString(key, value)
	}
}

// SetString ...
func SetString(key string, value string) {
	os.Setenv(prepareKey(key), value)
}

// GetString ...
func GetString(key string) (value string, ok bool) {
	key = prepareKey(key)

	if v, ok := os.LookupEnv(key); ok {
		return v, true
	}

	return "", false
}

// MustGetString ...
func MustGetString(key string) (value string) {
	value, ok := GetString(key)
	if !ok {
		panic("Environment variable \"" + prepareKey(key) + "\" not found")
	}

	return value
}

// SetDefaultBool ...
func SetDefaultBool(key string, value bool) {
	if !IssetKey(key) {
		SetBool(key, value)
	}
}

// SetBool ...
func SetBool(key string, value bool) {
	SetString(key, strconv.FormatBool(value))
}

// GetBool ...
func GetBool(key string) (value bool, ok bool) {
	str, ok := GetString(key)
	if ok {
		return parseBool(str)
	}

	return false, true
}

// MustGetBool ...
func MustGetBool(key string) (value bool) {
	str, ok := GetString(key)
	if ok {
		if value, ok := parseBool(str); ok {
			return value
		}

		panic("Can not convert environment variable \"" +
			prepareKey(key) + "\" to type boolean")
	}

	return false
}

//SetDefaultDuration ...
func SetDefaultDuration(key string, value time.Duration) {
	if !IssetKey(key) {
		SetDuration(key, value)
	}
}

// SetDuration ...
func SetDuration(key string, value time.Duration) {
	SetString(key, value.String())
}

// GetDuration ...
func GetDuration(key string) (value time.Duration, ok bool) {
	str, ok := GetString(key)

	if ok {
		v, err := time.ParseDuration(str)
		if err == nil {
			return v, true
		}

		log.Println(err)
	}

	return 0, false
}

// MustGetDuration  ...
func MustGetDuration(key string) time.Duration {
	str, ok := GetString(key)

	if ok {
		v, err := time.ParseDuration(str)
		if err == nil {
			return v
		}

		panic("Failed to parse duration from environment variable \"" + prepareKey(key) + "\" with error: " + err.Error())
	}

	panic("Environment variable \"" + prepareKey(key) + "\" not found")
}

// SetDefaultFloat64 ...
func SetDefaultFloat64(key string, value float64) {
	if !IssetKey(key) {
		SetFloat64(key, value)
	}
}

// SetFloat64 ...
func SetFloat64(key string, value float64) {
	SetString(key, strconv.FormatFloat(value, 'f', -1, 64))
}

// GetFloat64 ...
func GetFloat64(key string) (value float64, ok bool) {
	str, ok := GetString(key)

	if ok {
		v, err := strconv.ParseFloat(str, 64)
		if err == nil {
			return v, true
		}

		log.Println(err)
	}

	return 0, false
}

// MustGetFloat64 ...
func MustGetFloat64(key string) float64 {
	str, ok := GetString(key)

	if ok {
		v, err := strconv.ParseFloat(str, 64)
		if err == nil {
			return v
		}

		panic("Failed to parse float64 from environment variable \"" + prepareKey(key) + "\" with error: " + err.Error())
	}

	panic("Environment variable \"" + prepareKey(key) + "\" not found")
}

// SetDefaultInt ...
func SetDefaultInt(key string, value int) {
	if !IssetKey(key) {
		SetInt(key, value)
	}
}

// SetInt ...
func SetInt(key string, value int) {
	SetString(key, strconv.FormatInt(int64(value), 10))
}

// GetInt ...
func GetInt(key string) (value int, ok bool) {
	str, ok := GetString(key)

	if ok {
		v, err := strconv.ParseInt(str, 10, 32)
		if err == nil {
			return int(v), true
		}

		log.Println(err)
	}

	return 0, false
}

// MustGetInt ...
func MustGetInt(key string) (value int) {
	str, ok := GetString(key)

	if ok {
		v, err := strconv.ParseInt(str, 10, 32)
		if err == nil {
			return int(v)
		}

		panic("Failed to parse int from environment variable \"" + prepareKey(key) + "\" with error: " + err.Error())
	}

	panic("Environment variable \"" + prepareKey(key) + "\" not found")
}

// SetDefaultInt64 ...
func SetDefaultInt64(key string, value int64) {
	if !IssetKey(key) {
		SetInt64(key, value)
	}
}

// SetInt64 ...
func SetInt64(key string, value int64) {
	SetString(key, strconv.FormatInt(value, 10))
}

// GetInt64 ...
func GetInt64(key string) (value int64, ok bool) {
	str, ok := GetString(key)

	if ok {
		v, err := strconv.ParseInt(str, 10, 64)
		if err == nil {
			return v, true
		}

		log.Println(err)
	}

	return 0, false
}

// MustGetInt64 ...
func MustGetInt64(key string) (value int64) {
	str, ok := GetString(key)

	if ok {
		v, err := strconv.ParseInt(str, 10, 64)
		if err == nil {
			return v
		}

		panic("Failed to parse int64 from environment variable \"" + prepareKey(key) + "\" with error: " + err.Error())
	}

	panic("Environment variable \"" + prepareKey(key) + "\" not found")
}

// SetDefaultUInt ...
func SetDefaultUInt(key string, value uint) {
	if !IssetKey(key) {
		SetUInt(key, value)
	}
}

// SetUInt ...
func SetUInt(key string, value uint) {
	SetString(key, strconv.FormatUint(uint64(value), 10))
}

// GetUInt ...
func GetUInt(key string) (value uint, ok bool) {
	str, ok := GetString(key)

	if ok {
		v, err := strconv.ParseUint(str, 10, 32)
		if err == nil {
			return uint(v), true
		}

		log.Println(err)
	}

	return 0, false
}

// MustGetUInt ...
func MustGetUInt(key string) (value uint) {
	str, ok := GetString(key)

	if ok {
		v, err := strconv.ParseUint(str, 10, 32)
		if err == nil {
			return uint(v)
		}

		panic("Failed to parse uint from environment variable \"" + prepareKey(key) + "\" with error: " + err.Error())
	}

	panic("Environment variable \"" + prepareKey(key) + "\" not found")
}

// SetDefaultUInt64 ...
func SetDefaultUInt64(key string, value uint64) {
	if !IssetKey(key) {
		SetUInt64(key, value)
	}
}

// SetUInt64 ...
func SetUInt64(key string, value uint64) {
	SetString(key, strconv.FormatUint(value, 10))
}

// GetUInt64 ...
func GetUInt64(key string) (value uint64, ok bool) {
	str, ok := GetString(key)

	if ok {
		v, err := strconv.ParseUint(str, 10, 64)
		if err == nil {
			return v, true
		}

		log.Println(err)
	}

	return 0, false
}

// MustGetUInt64 ...
func MustGetUInt64(key string) (value uint64) {
	str, ok := GetString(key)

	if ok {
		v, err := strconv.ParseUint(str, 10, 64)
		if err == nil {
			return v
		}

		panic("Failed to parse uint64 from environment variable \"" + prepareKey(key) + "\" with error: " + err.Error())
	}

	panic("Environment variable \"" + prepareKey(key) + "\" not found")
}
