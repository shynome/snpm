package tools

import (
	"os"
)

// FileExists check file is exists
func FileExists(f string) bool {

	info, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()

}
