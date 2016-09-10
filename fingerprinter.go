// Copyright 2015, 2016 Dominik Pataky <dom@netdecorator.org>
// This file is part of Fingerprinter, for licence details see LICENCE

// Fingerprinter
// Use this tool to let others check if a fingerprint really belongs to one of your keys.

package fingerprinter

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
    Datafile    string
    Keydir      string
    Templatedir string
    Staticdir   string
}

func StartServer() {
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

    // Check needed paths for existence, eg the CSV file, template and static folders
    toCheck := []string{config.Datafile, config.Templatedir, config.Staticdir}
    for _, e := range toCheck {
        if ok, err := utils.PathExistsErr(e); !ok {
            log.Fatal(err)
        }
    }

    // If the value is not given in the config, it's an empty string.
    if config.Keydir != "" {
        // If not an empty string, check if it's a folder.
        if ok, _ := utils.IsDirectory(config.Keydir); ok {
            repo.ScanPGPKeys(config.Keydir)
        } else {
            log.Println("config.Keydir is given, but not a valid path.")
        }
    }

    log.Println("Starting server on port", config.Port)

    r := repo.GetRepo()
    r.ReadDatafile(config.Datafile)
    log.Println("The following entities were read:")
    r.Print()

    web.Run(config.Port, config.Baseurl, config.Templatedir, config.Staticdir)
}
