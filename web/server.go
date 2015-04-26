// Copyright 2015 Dominik Pataky <mail@netdecorator.org>
// This file is part of Fingerprinter, for licence details see LICENCE

package web

import (
    "io"
    "log"
    "net/http"
    "strings"
)

var _baseurl string

// Helper function, wraps http.HandleFunc and inserts the baseurl
// into all endpoints
func registerEndpoint(endpoint string,
        handler func(http.ResponseWriter, *http.Request)) {
    // Remove leading slashes from endpoints
    if strings.HasPrefix(endpoint, "/") {
        endpoint = endpoint[1:]
    }

    http.HandleFunc(_baseurl + endpoint, handler)
}

// Helper function, wraps a logger around http.Handler
func Logger(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
        handler.ServeHTTP(w, r)
    })
}

// Handler for requests to "/"
func IndexHandler(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "hello\n")
}

// Handler for requests to "/check"
func CheckHandler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, _baseurl, http.StatusFound)
}

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
