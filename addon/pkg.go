package addon

import (
	"github.com/konveyor/tackle2-hub/addon"
)

// Environment.
const (
	EnvSharedDir = addon.EnvSharedDir
	EnvCacheDir  = addon.EnvCacheDir
	EnvToken     = addon.EnvToken
	EnvTask      = addon.EnvTask
)

// Addon adapter.
var Addon = addon.Addon
