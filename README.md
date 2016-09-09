# Fingerprinter

With Fingerprinter you can run a webinterface which let's others check
fingerprints against a configured list of your own key fingerprints.

Simply compile, customize the config, add all your fingerprints to the
fingerprints.dat CSV file and start the service. It automatically parses the CSV
and creates the data entries.

## Usage

To use the fingerprinter library, create a `main.go` program like this:

```golang
package main

import "fingerprinter"

func main() {
    fingerprinter.StartServer()
}
```

This could for example be placed in `$GOPATH/src/fingerprinter-bin/main.go`.

Then go into `$GOPATH/src/` and `git pull github.com/cooox/fingerprinter`. You can then build a binary with `go install fingerprinter-bin` which will be created at `$GOPATH/bin/fingerprinter-bin`.

## Screenshots

Pretty simple interface.

<img src="input.png" style="max-width: 400px; border: 1px solid gray" >

<img src="output.png" style="max-width: 400px; border: 1px solid gray" >

## Licence

This software is licenced under the GNU AGPL licence. See the LICENCE file for
more information.
