// Copyright 2015 Dominik Pataky <mail@netdecorator.org>
// This file is part of Fingerprinter, for licence details see LICENCE

// Fingerprinter 
// Use this tool to let others check if a fingerprint really belongs to one of your keys.

package main

import (
    "encoding/json"
    "log"
    "os"

    "fingerprinter/repo"
    "fingerprinter/utils"
    "fingerprinter/web"
)


// Config with all needed fields
type Config struct {
    Port        string
    Baseurl     string
    Keydir      string
    Datafile    string
    Templatedir string
    Staticdir   string
}

func main() {
    config := Config{}

    fh, err := os.Open("config.json")
    if err != nil {
        log.Fatal(err)
    }
    dec := json.NewDecoder(fh)
    err = dec.Decode(&config)
    if err != nil {
        log.Println("Error parsing config.json")
        log.Fatal(err)
    }
    log.Println("Config okay, checking paths")

    toCheck := []string{config.Keydir, config.Datafile, config.Templatedir, config.Staticdir}
    for _, e := range toCheck {
        if ok, err := utils.PathExistsErr(e); !ok {
            log.Fatal(err)
        }
    }

    log.Println("Starting server on port", config.Port)

    r := repo.GetRepo()
    r.ReadDatafile(config.Datafile)
    log.Println("The following records were read from the data file:")
    r.Print()

    web.Run(config.Port, config.Baseurl, config.Templatedir, config.Staticdir)
}
