// Copyright 2015 Dominik Pataky <mail@netdecorator.org>
// This file is part of Fingerprinter, for licence details see LICENCE

package web

import (
    "html/template"
    "log"
    "net/http"
    "strings"
)

type pageData struct {
    Title string
    Baseurl string
    Payload map[string]interface{}
}

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

// Configure the template paths and parse them. Returns a template which can
// then be executed.
func parse(files ...string) (*template.Template) {
    // Prepare a buffer array for the combination of base + files
    buff := make([]string, 0)

    // Insert the base template as first element!
    buff = append(buff, _templatedir + "base.html")

    // Append all other template names
    for _, f := range files {
        buff = append(buff, _templatedir + f)
    }

    // Now parse base template + all others
    t, err := template.ParseFiles(buff...)
    if err != nil {
        log.Panic(err)
    }
    return t
}

// Create a new PageData struct with a given page title and the configured
// baseurl. Also makes the Payload map.
func newPD(title string) (pageData) {
    return pageData{
        Title: title,
        Baseurl: _baseurl,
        Payload: make(map[string]interface{}),
    }
}

// Add a payload field to the PageData struct. Will be accessable via
// '.Payload.KEY' in the template.
func (pd *pageData) addPayload(key string, value interface{}) {
    pd.Payload[key] = value
}
