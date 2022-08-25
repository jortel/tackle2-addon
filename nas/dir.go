/*
Package nas provides support for efficiently working with
network attached storage (NAS).
*/
package nas

import (
	"errors"
	liberr "github.com/konveyor/controller/pkg/error"
	"github.com/konveyor/tackle2-addon/command"
	"os"
)

//
// RmDir deletes the directory.
func RmDir(path string) (err error) {
	cmd := command.Command{Path: "/usr/bin/rm"}
	cmd.Options.Add("-rf", path)
	err = cmd.Run()
	return
}

//
// HasDir return if the path exists.
func HasDir(path string) (found bool, err error) {
	found, err = Exists(path)
	return
}

//
// MkDir ensures the directory exists.
func MkDir(path string, mode os.FileMode) (err error) {
	err = os.MkdirAll(path, mode)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			err = nil
		} else {
			err = liberr.Wrap(
				err,
				"path",
				path)
		}
	}
	return
}

//
// Exists return if the path exists.
func Exists(path string) (found bool, err error) {
	_, err = os.Stat(path)
	if err == nil {
		found = true
		return
	}
	if !os.IsNotExist(err) {
		err = liberr.Wrap(
			err,
			"path",
			path)
	} else {
		err = nil
	}
	return
}
