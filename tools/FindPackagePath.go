package tools

import (
	"path/filepath"
)

const packageJSON = "package.json"

// FindPackagePath find package.json path
func FindPackagePath(wd string) string {

	f := filepath.Join(wd, packageJSON)

	if FileExists(f) {
		return f
	}

	newWd := filepath.Join(wd, "../")

	if newWd == wd {
		return ""
	}

	return FindPackagePath(newWd)

}
