// Copyright 2015, 2016 Dominik Pataky <dom@netdecorator.org>
// This file is part of Fingerprinter, for licence details see LICENCE

package utils

import (
    "os"
    "strings"
)


// If path does not end with '/' append trailing slash
func AppendSlash(path *string) {
    if !strings.HasSuffix(*path, "/") {
        *path = *path + "/"
    }
}

// Remove characters from a string
func Sanitizer(input *string, filters ...string) {
    for _, filter := range filters {
        if strings.ContainsAny(*input, filter) {
            *input = strings.Replace(*input, filter, "", -1)
        }
    }
}

// Check if a path exists
func PathExists(path string) bool {
    _, err := os.Stat(path)
    if os.IsNotExist(err) {
        return false
    }
    return true
}

// Check if a path exists and return bool with os.Stat error
func PathExistsErr(path string) (bool, error) {
    _, err := os.Stat(path)
    if os.IsNotExist(err) {
        return false, err
    }
    return true, err
}

// Check if a path is a directory
func IsDirectory(path string) (bool, error) {
    if ex, err := PathExistsErr(path); !ex {
        return false, err
    }

    fileInfo, err := os.Stat(path)
    return fileInfo.IsDir(), err
}
