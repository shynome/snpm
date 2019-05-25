package tools

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func readJSONFile(f string) map[string]interface{} {
	file, _ := ioutil.ReadFile(f)
	data := map[string]interface{}{}
	json.Unmarshal(file, &data)
	return data
}

func TestMakeEnv(t *testing.T) {

	wd, _ := os.Getwd()

	f := FindPackagePath(filepath.Join(wd, "../test/p"))

	data := readJSONFile(f)

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
