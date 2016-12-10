// Copyright 2016 Stefan Böhmann.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
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
