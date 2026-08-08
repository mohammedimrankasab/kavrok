// Package version provides build and runtime version information.
package version

import "runtime"

var (
	version   = "dev"
	commit    = "none"
	buildDate = "unknown"
)

// Info contains build and runtime metadata.
type Info struct {
	Version   string
	Commit    string
	BuildDate string
	GoVersion string
	Platform  string
	TreeState string
}

// Get returns the current application version information.
func Get() Info {
	return Info{
		Version:   version,
		Commit:    commit,
		BuildDate: buildDate,
		GoVersion: runtime.Version(),
		Platform:  runtime.GOOS + "/" + runtime.GOARCH,
		TreeState: "unknown",
	}
}
