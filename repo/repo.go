// Copyright 2015 Dominik Pataky <mail@netdecorator.org>
// This file is part of Fingerprinter, for licence details see LICENCE

package repo

import (
    "encoding/csv"
    "log"
    "os"
    "sync"
)

type RepoEntry struct {
    Fingerprint string
    Created string
    Identity string
    Service string
    Comment string
}

// Repo holding all collected RepoEntry.
// Provides functions like Add().
type Repo struct {
    Entries []*RepoEntry
    Lock sync.Mutex
}

// Print details of a RepoEntry to log
func (e *RepoEntry) Print() {
    log.Printf("Fingerprint %s for ID %s on service %s\n", e.Fingerprint, e.Identity, e.Service)
}

// Print all RepoEntry to the log
func (r *Repo) Print() {
    for _, e := range r.Entries {
        e.Print()
    }
}

// Read the datafile and parse as CSV.
// For each record a new RepoEntry is created and added to the repo
func (r *Repo) ReadDatafile(datafile string) {
    csvfile, err := os.Open(datafile)
    if err != nil {
        log.Println(err)
        return
    }
    defer csvfile.Close()

    reader := csv.NewReader(csvfile)
    reader.FieldsPerRecord = 5
    reader.TrimLeadingSpace = true

    data, err := reader.ReadAll()
    if err != nil {
        log.Println(err)
        return
    }

    for _, e := range data {
        log.Println(e)
        r.Add(newRepoEntry(e[0], e[1], e[2], e[3], e[4]))
    }
}

// Add a RepoEntry to a Repo
func (r *Repo) Add(entry *RepoEntry) {
    r.Lock.Lock()
    defer r.Lock.Unlock()

    appendFunc := func(newKey *RepoEntry) {
        r.Entries = append(r.Entries, newKey)
    }

    // If repo is empty, don't check for duplicates
    if len(r.Entries) == 0 {
        appendFunc(entry)
        return
    }

    // Avoid duplicates
    for _, v := range r.Entries {
        if v.Fingerprint == entry.Fingerprint {
            log.Printf("Key %s already in repo.\n", v.Fingerprint)
            return
        }
    }

    // Key not in repo, append
    appendFunc(entry)

    return
}

// Create a new RepoEntry and return the pointer
func newRepoEntry(fp, created, id, service, comment string) *RepoEntry {
    e := new(RepoEntry)
    e.Fingerprint = fp
    e.Created = created
    e.Identity = id
    e.Service = service
    e.Comment = comment

    return e
}

