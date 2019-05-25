package tools

import (
	"path/filepath"
)

const packageJSON = "package.json"

// FindPackagePath find package.json path
func FindPackagePath(pwd string) string {

	f := filepath.Join(pwd, packageJSON)

	if FileExists(f) {
		return f
	}

	newPwd := filepath.Join(pwd, "../")

	if newPwd == pwd {
		return ""
	}

	return FindPackagePath(newPwd)

}
