package tools

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindPackagePath(t *testing.T) {

	pwd, _ := os.Getwd()

	testCase := [][2]string{
		{"../test/p/", "../test/p/"},
		{"../test/p/a", "../test/p/a"},
		{"../test/p/b", "../test/p/"},
		{"../test/", ""},
	}

	for _, v := range testCase {

		p1 := FindPackagePath(filepath.Join(pwd, v[0]))
		p2 := filepath.Join(pwd, v[1], packageJSON)
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
