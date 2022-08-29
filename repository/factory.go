package repository

import (
	hub "github.com/konveyor/tackle2-hub/addon"
	"github.com/konveyor/tackle2-hub/api"
	"os"
)

var (
	addon   = hub.Addon
	HomeDir = ""
)

func init() {
	HomeDir, _ = os.UserHomeDir()
}

type SoftError = hub.SoftError

//
// New SCM repository factory.
func New(destDir string, application *api.Application) (r Repository, err error) {
	kind := application.Repository.Kind
	switch kind {
	case "subversion":
		r = &Subversion{}
	default:
		r = &Git{}
	}
	r.With(destDir, application)
	err = r.Validate()
	return
}

//
// Repository interface.
type Repository interface {
	//
	// With constructor.
	With(path string, application *api.Application)
	//
	// Validate settings.
	Validate() (err error)
	//
	// Fetch the repository.
	Fetch() (err error)
	//
	// Add a file.
	Add(path string) (err error)
	//
	// Delete a file.
	Delete(path string) (err error)
	//
	// CreateBranch creates a branch.
	CreateBranch(name string) (err error)
	//
	// DeleteBranch deletes a branch.
	DeleteBranch(name string) (err error)
	//
	// Commit changes.
	Commit() (err error)
}

//
// SCM - source code manager.
type SCM struct {
	Application *api.Application
	Path        string
}

//
// With settings.
func (r *SCM) With(path string, application *api.Application) {
	r.Application = application
	r.Path = path
}
