// Copyright 2015 Dominik Pataky <mail@netdecorator.org>
// This file is part of Fingerprinter, for licence details see LICENCE

package web

import (
    "net/http"
)

// Handler for requests to "/"
func IndexHandler(w http.ResponseWriter, r *http.Request) {
    data := newPD("Index")
    data.addPayload("FormUrl", "check")
    t := parse("index.html")
    t.Execute(w, data)
}

// Handler for requests to "/check"
func CheckHandler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, _baseurl, http.StatusFound)
}
