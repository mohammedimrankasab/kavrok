package version

import "runtime"

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
