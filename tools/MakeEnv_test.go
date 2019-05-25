package tools

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMakeEnv(t *testing.T) {

	wd, _ := os.Getwd()

	f := FindPackagePath(filepath.Join(wd, "../test/p"))

	data := ReadJSONFile(f)

	env := map[string]string{}

	MakeEnv(data, "npm_package", &env)

	checkIsNotSet := func(name string) bool {
		return env[name] == ""
	}

	for _, v := range []string{"npm_package_name", "npm_package_config_port"} {
		if checkIsNotSet(v) {
			t.Errorf("env %v not set", v)
		}
	}

	return

}
