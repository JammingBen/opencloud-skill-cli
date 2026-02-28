package version

var (
	// Version is the version of the OpenCloud CLI, which can be set at build time via ldflags.
	Version = "0.0.1"
)

// GetVersion returns the current version of the OpenCloud CLI
func GetVersion() string {
	return Version
}

// GetVersionWithoutSuffix returns the version without any suffix (e.g. "-beta", "-rc1")
func GetVersionWithoutSuffix() string {
	for i, c := range Version {
		if c == '-' {
			return Version[:i]
		}
	}
	return Version
}
