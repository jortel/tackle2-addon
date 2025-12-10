package addon

import (
	"github.com/konveyor/tackle2-hub/addon"
)

// Environment.
const (
	EnvSharedDir = addon.EnvSharedPath
	EnvCacheDir  = addon.EnvCachePath
	EnvToken     = addon.EnvHubToken
	EnvTask      = addon.EnvTask
)

var Addon addon.Addon
