// Copyright 2015 Dominik Pataky <mail@netdecorator.org>
// This file is part of Fingerprinter, for licence details see LICENCE

package web

import (
    "net/http"

    "fingerprinter/repo"
    "fingerprinter/utils"
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
    inputFingerprint := r.FormValue("fingerprint")
    if inputFingerprint != "" {
        utils.Sanitizer(&inputFingerprint, ":", " ")

        data := newPD("Check result")
        data.addPayload("fp", inputFingerprint)

        r := repo.GetRepo()
        if ok, e := r.GetEntry(inputFingerprint); ok {
            data.addPayload("Entry", e)
        }

        t := parse("result.html")
        t.Execute(w, data)

        return
    }
    http.Redirect(w, r, _baseurl, http.StatusFound)
}
