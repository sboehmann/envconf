envconf
=======

A Go (Golang) library for managing configuration data from environment variables which is used by my Twelve-Factor Apps.

[![GoDoc](https://godoc.org/github.com/spf13/hugo?status.svg)](https://godoc.org/github.com/spf13/hugo)&nbsp;[![Travis CI Status](https://travis-ci.org/sboehmann/envconf.svg?branch=master)](https://travis-ci.org/sboehmann/envconf)&nbsp;[![Coverage Status](https://coveralls.io/repos/github/sboehmann/envconf/badge.svg?branch=master)](https://coveralls.io/github/sboehmann/envconf?branch=master)&nbsp;[![Go Report Card](https://goreportcard.com/badge/github.com/sboehmann/envconf)](https://goreportcard.com/report/github.com/sboehmann/envconf)

Installation
------------

Download and install it:

```sh
go get -u github.com/sboehmann/envconf
```

Import it in your code:

```go
import "github.com/sboehmann/envconf"
```

Quick start
-----------

First you need a few environment variables:

```sh
export MY_PORT=8080
export MY_MESSAGE="Hello World!"
```

And now a simple program that uses these variables:

```go
package main

import (
    "github.com/sboehmann/envconf"
    "io"
    "net/http"
)

func message(w http.ResponseWriter, r *http.Request) {
    value, _ := envconf.GetString("MESSAGE")
    io.WriteString(w, value)
}

func main() {
    envconf.SetPrefix("MY_")

    http.HandleFunc("/", message)
    http.ListenAndServe(":"+envconf.MustGetString("PORT"), nil)
}
```
