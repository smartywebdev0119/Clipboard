// Copyright 2013 @atotto. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin

package clipboard

import (
	"os/exec"
)

var (
	pbpasteCmd = exec.Command("pbpaste")
	pbcopyCmd  = exec.Command("pbcopy")
)

func readAll() string {
	out, err := pbpasteCmd.Output()
	if err != nil {
		panic(err)
	}
	return string(out)
}

func writeAll(text string) {
	in, err := pbcopyCmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	pbcopyCmd.Start()
	in.Write([]byte(text))
	in.Close()
	pbcopyCmd.Wait()
}
