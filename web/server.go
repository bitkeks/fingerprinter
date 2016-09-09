// Copyright 2015, 2016 Dominik Pataky <dom@netdecorator.org>
// This file is part of Fingerprinter, for licence details see LICENCE

package web

import (
    "net/http"

    "fingerprinter/utils"
)

var (
    _baseurl,
    _templatedir string
)

// Establish endpoints and static paths and run the server
func Run(port, baseurl, templatedir, staticdir string) {
    utils.AppendSlash(&baseurl)
    _baseurl = baseurl

    utils.AppendSlash(&templatedir)
    _templatedir = templatedir

    registerEndpoint("/", IndexHandler)
    registerEndpoint("/check", CheckHandler)

    http.Handle(_baseurl + "static/", http.StripPrefix(_baseurl + "static/",
        http.FileServer(http.Dir(staticdir))))

    http.ListenAndServe(":" + port, Logger(http.DefaultServeMux))
}
