package snpm

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewPackage(t *testing.T) {

	wd, _ := os.Getwd()

	wd = filepath.Join(wd, "./test/p")

	pkg, err := NewPackage(wd)

	if err != nil {
		t.Error(err)
		return
	}

	if len(pkg.Scripts) != 3 {
		t.Error("pkg script length is not right, the current scripts is: ", pkg.Scripts)
	}

	return

}
