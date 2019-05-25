package tools

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindPackagePath(t *testing.T) {

	wd, _ := os.Getwd()

	testCase := [][2]string{
		{"../test/p/", "../test/p/"},
		{"../test/p/a", "../test/p/a"},
		{"../test/p/b", "../test/p/"},
		{"../test/", ""},
	}

	for _, v := range testCase {

		p1 := FindPackagePath(filepath.Join(wd, v[0]))
		p2 := filepath.Join(wd, v[1], packageJSON)
		if v[1] == "" {
			p2 = ""
		}

		if p1 != p2 {
			t.Errorf("get pacakge.json path fail when path is %v , has get %v", v[0], p1)
			return
		}

	}

	return

}
