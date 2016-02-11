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
    entries []*RepoEntry
    l sync.Mutex
}

var _repo *Repo

// Print details of a RepoEntry to log
func (e *RepoEntry) Print() {
    log.Printf("Fingerprint %s for ID %s on service %s\n", e.Fingerprint, e.Identity, e.Service)
}

// Print the Entrys fingerprint with seperators after every 4 characters.
func (e *RepoEntry) PrintSep() string {
    fp := e.Fingerprint
    res := ""
    for idx, e := range fp {
        if idx != 0 && idx % 4 == 0 {
            res = res + " "
        }
        res = res + string(e)
    }
    return res
}

// Print all RepoEntry to the log
func (r *Repo) Print() {
    for _, e := range r.GetEntries() {
        e.Print()
    }
}

// Get a list of all RepoEntry in this Repo
func (r *Repo) GetEntries() ([]*RepoEntry) {
    return r.entries
}

// Get a RepoEntry by its fingerprint.
func (r *Repo) GetEntry(fp string) (bool, *RepoEntry) {
    for _, e := range r.entries {
        if e.Fingerprint == fp {
            return true, e
        }
    }
    return false, nil
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
        r.Add(newRepoEntry(e[0], e[1], e[2], e[3], e[4]))
    }
}

// Add a RepoEntry to a Repo
func (r *Repo) Add(entry *RepoEntry) {
    r.lock()
    defer r.unlock()

    appendFunc := func(newKey *RepoEntry) {
        r.entries = append(r.entries, newKey)
    }

    // If repo is empty, don't check for duplicates
    if len(r.entries) == 0 {
        appendFunc(entry)
        return
    }

    // Avoid duplicates
    for _, v := range r.entries {
        if v.Fingerprint == entry.Fingerprint {
            log.Printf("WARNING: Fingerprint %s has multiple records, only the first one is used!\n", v.Fingerprint)
            return
        }
    }

    // Key not in repo, append
    appendFunc(entry)

    return
}

func (r *Repo) lock() {
    r.l.Lock()
}

func (r *Repo) unlock() {
    r.l.Unlock()
}

// Singleton Repo object
func GetRepo() (*Repo) {
    if _repo == nil {
        _repo = new(Repo)
    }
    return _repo
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

