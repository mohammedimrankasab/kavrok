// Package version provides build and runtime version information.
package version

import "runtime"

// Info contains build and runtime metadata.
type Info struct {
	Version   string
	Commit    string
	BuildDate string

	GoVersion string
	OS        string
	Arch      string
}

var (
	version   = "dev"
	commit    = "none"
	buildDate = "unknown"
)

// Get returns the current application version information.
func Get() Info {
	return Info{
		Version:   version,
		Commit:    commit,
		BuildDate: buildDate,

		GoVersion: runtime.Version(),
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}
}
