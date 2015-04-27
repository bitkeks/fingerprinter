// Copyright 2015 Dominik Pataky <mail@netdecorator.org>
// This file is part of Fingerprinter, for licence details see LICENCE

package web

import (
    "log"
    "net/http"
    "strings"
)

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
