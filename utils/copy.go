// Copyright 2016 Tobin Harding. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"io"
	"log"
	"os"
)

const mode = 0644

func Cp(dst, src string) error {
	s, err := os.Open(src)
	if err != nil {
		return err
	}
	// no need to check errors on read only file, we already got everything
	// we need from the filesystem, so nothing can go wrong now.
	defer s.Close()
	d, err := os.Create(dst)
	if err != nil {
		return err
	}
	err = os.Chmod(dst, mode)
	if err != nil {
		os.Remove(dst)
		return err
	}
	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		os.Remove(dst)
		return err
	}
	return d.Close()
}

// mustCopy copy src to dst
// ref: The Go Programming Language p221
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
