// Copyright 2015 Dominik Pataky <mail@netdecorator.org>
// This file is part of Fingerprinter, for licence details see LICENCE

package web

import (
    "net/http"
    "strings"
)

var _baseurl string

// Establish endpoints and static paths and run the server
func Run(port, baseurl, templatedir, staticdir string) {
    // If not given, append an appending slash
    if !strings.HasSuffix(baseurl, "/") {
        baseurl = baseurl + "/"
    }
    _baseurl = baseurl

    registerEndpoint("/", IndexHandler)
    registerEndpoint("/check", CheckHandler)

    http.Handle(_baseurl + "static/", http.StripPrefix(_baseurl + "static/",
        http.FileServer(http.Dir(staticdir))))

    http.ListenAndServe(":" + port, Logger(http.DefaultServeMux))
}
