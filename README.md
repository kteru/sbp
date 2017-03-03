sbp
===

**CURRENTLY IN DEVELOPMENT**

Parser/Serializer for the Swift Navigation Binary Protocol (SBP). This is compatible with SBP protocol specification v0.52.4 (stable messages).

[![GoDoc](https://godoc.org/github.com/kteru/sbp?status.svg)](https://godoc.org/github.com/kteru/sbp)

Installation
------------

```
$ go get -u github.com/kteru/sbp
```

Example
-------

```
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/kteru/sbp"
)

func main() {
	frd := sbp.NewFrameReader(os.Stdin)

	for {
		fr, err := frd.Next()
		if err == io.EOF {
			os.Exit(0)
		}
		if err != nil {
			continue
		}

		msg, _ := fr.Msg()

		switch m := msg.(type) {
		case *sbp.MsgPosLlh:
			fmt.Printf("%.9f,%.9f,%.3f,%d\n", m.Lat, m.Lon, m.Height, m.FixMode)
		}
	}
}
```
