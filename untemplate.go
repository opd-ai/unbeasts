package untemplate

import "fmt"

// LibraryName returns the package name used by generated projects.
func LibraryName() string {
	return "untemplate"
}

// Version returns a static template version.
func Version() string {
	return "0.1.0"
}

// Summary provides a one-line summary suitable for CLIs.
func Summary() string {
	return fmt.Sprintf("%s template %s", LibraryName(), Version())
}
