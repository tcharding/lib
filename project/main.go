// Copyright 2016 Tobin Harding. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// project: utilities for managing projects

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	template  = "/home/tobin/build/go/src/github.com/tcharding/project/template"
	filePerms = 0644
	dirPerms  = 0755
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: project <name>\n")
		os.Exit(1)
	}
	projectName := os.Args[1]

	log.SetFlags(0)
	log.SetPrefix("project:")

	err := os.Mkdir(projectName, dirPerms)
	if err != nil {
		log.Fatal(err)
	}

	err = createTest(projectName)
	if err != nil {
		log.Fatal(err)
	}

	err = createFile(projectName)
	if err != nil {
		log.Fatal(err)
	}

}

func createTest(projectName string) error {
	s, err := os.Open(template)
	if err != nil {
		s.Close()
		return fmt.Errorf("failed to open template file: %s", template)
	}
	defer s.Close()

	dstName := projectName + "/" + projectName + "_test.go"
	d, err := os.Create(dstName)
	if err != nil {
		return fmt.Errorf("failed to create file: %s", dstName)
	}
	defer d.Close()

	if err = os.Chmod(dstName, filePerms); err != nil {
		os.Remove(dstName)
		return fmt.Errorf("failed to chmod: %s %v", dstName, filePerms)
	}

	d.WriteString("package " + projectName + "\n")
	d.WriteString("\n")

	if _, err = io.Copy(d, s); err != nil {
		os.Remove(dstName)
		return fmt.Errorf("failed to copy to: %s", dstName)
	}
	return nil
}

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
	defer d.Close()
	err = os.Chmod(dst, filePerms)
	if err != nil {
		os.Remove(dst)
		return err
	}
	if _, err := io.Copy(d, s); err != nil {
		os.Remove(dst)
		return err
	}
	return nil
}

func createFile(projectName string) error {
	dstName := projectName + "/" + projectName + ".go"
	f, err := os.Create(dstName)
	if err != nil {
		return fmt.Errorf("failed to create file: %s", dstName)
	}
	defer f.Close()

	err = os.Chmod(dstName, filePerms)
	if err != nil {
		os.Remove(dstName)
		return fmt.Errorf("failed to chmod: %s %v", dstName, filePerms)
	}

	f.WriteString("// " + projectName + "\n")
	f.WriteString("package " + projectName + "\n")

	return nil
}
