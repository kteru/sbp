sbp
===

Parser/Serializer for the Swift Navigation Binary Protocol (SBP). This is compatible with SBP protocol specification v0.52.4 (stable messages).

[![GoDoc](https://godoc.org/github.com/kteru/sbp?status.svg)](https://godoc.org/github.com/kteru/sbp)
[![Build Status](https://travis-ci.org/kteru/sbp.svg?branch=master)](https://travis-ci.org/kteru/sbp)
[![Coverage Status](https://coveralls.io/repos/github/kteru/sbp/badge.svg?branch=master)](https://coveralls.io/github/kteru/sbp?branch=master)

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
		fr, err := frd.ReadFrame()
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
