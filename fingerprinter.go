// Copyright 2015 Dominik Pataky <mail@netdecorator.org>
// This file is part of Fingerprinter, for licence details see LICENCE

// Fingerprinter 
// Use this tool to let others check if a fingerprint really belongs to one of your keys.

package main

import (
    "encoding/json"
    "log"
    "os"

    "fingerprinter/utils"
    "fingerprinter/web"
)


// Config with all needed fields
type Config struct {
    Baseurl     string
    Keydir      string
    Datafile    string
    Templatedir string
}

func main() {
    config := Config{}

    fh, err := os.Open("config.json")
    if err != nil {
        log.Println(err)
    }
    dec := json.NewDecoder(fh)
    err = dec.Decode(&config)
    if err != nil {
        log.Println(err)
    }

    toCheck := []string{config.Keydir, config.Datafile, config.Templatedir}
    for _, e := range toCheck {
        if !utils.PathExists(e) {
            log.Printf("Path '%s' does not exist.\n", e)
        }
    }

    web.Run()
}
