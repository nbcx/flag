// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package flag_test

import (
	"fmt"

	"github.com/nbcx/flag"
)

func ExampleShorthandLookup() {
	name := "verbose"
	short := name[:1]

	flag.BoolP(name, short, false, "verbose output")

	// len(short) must be == 1
	flag := flag.ShorthandLookup(short)

	fmt.Println(flag.Name)
}

func ExampleFlagSet_ShorthandLookup() {
	name := "verbose"
	short := name[:1]

	fs := flag.NewFlagSet("Example", flag.ContinueOnError)
	fs.BoolP(name, short, false, "verbose output")

	// len(short) must be == 1
	flag := fs.ShorthandLookup(short)

	fmt.Println(flag.Name)
}
