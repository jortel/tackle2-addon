package main

import (
	"github.com/konveyor/tackle2-addon/command"
	"github.com/konveyor/tackle2-addon/scm"
	"github.com/konveyor/tackle2-addon/ssh"
	"github.com/konveyor/tackle2-hub/shared/api"
)

func main() {
	_ = command.New("")
	_ = ssh.Agent{}
	_, _ = scm.New("", api.Repository{}, nil)
}
