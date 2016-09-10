// Copyright 2015, 2016 Dominik Pataky <dom@netdecorator.org>
// This file is part of Fingerprinter, for licence details see LICENCE

package repo

import (
    "bytes"
    "encoding/hex"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"

    "golang.org/x/crypto/openpgp"

    "fingerprinter/utils"
)

// Collect all files in `rootpath` and parse them as OpenPGP ASCII armor files.
func ScanPGPKeys(rootpath string) {
    // Walk through all directories and fetch pubkeys
    err := filepath.Walk(rootpath, func (path string, f os.FileInfo, err error) error {
        if err != nil {
            panic(err)
        }

        if dir, err := utils.IsDirectory(path); dir && err == nil {
            // Don't scan a directory for ascii blocks
            return nil
        }

        // Pass a file path to parser function
        ParseArmorFile(path)

        return nil
    })

    if err != nil {
        log.Printf("filepath.Walk() returned %v\n", err)
    }
}

// Parse a single ASCII armored PGP pubkey file and add it to the repo.
func ParseArmorFile(keyPath string) {
    ascii, err := ioutil.ReadFile(keyPath)
    if err != nil {
        panic(err)
    }

    asciiReader := bytes.NewReader([]byte(ascii))

    // An armored key ring can contain one or more ascii key blobs.
    entityList, errReadArm := openpgp.ReadArmoredKeyRing(asciiReader)
    if errReadArm != nil {
        log.Println("Reading Pubkey ", errReadArm.Error())
        return
    }

    for _, pubKeyEntity := range entityList {
        if pubKeyEntity.PrimaryKey != nil {
            pubKey := *pubKeyEntity.PrimaryKey
            fingerprint := hex.EncodeToString(pubKey.Fingerprint[:])

            GetRepo().Add(newRepoEntry(
                strings.ToUpper(fingerprint),
                pubKey.CreationTime.String(),
                "0x" + pubKey.KeyIdString(),
                "email",
                "",
                string(ascii[:]),
            ))
        }
    }
    return
}
