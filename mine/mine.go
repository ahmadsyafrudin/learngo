package mine

import "runtime"

// Version is copy of runtime.Version
func Version() string {
	return runtime.Version()
}
