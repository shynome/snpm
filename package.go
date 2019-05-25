package snpm

import (
	"fmt"

	"path/filepath"

	"github.com/shynome/snpm/tools"
)

// Package package.json
type Package struct {
	RAW     map[string]interface{} // JSON 内容
	DIR     string                 // the working dir
	ENV     map[string]string      // 要设置的环境变量
	Scripts map[string]string      // 可执行的脚本
}

// NewPackage new package
func NewPackage(wd string) (Package, error) {

	f := tools.FindPackagePath(wd)

	data := tools.ReadJSONFile(f)

	env := map[string]string{}
	tools.MakeEnv(data, "npm_package", &env)

	scripts := map[string]string{}
	_scripts, ok := data["scripts"].(map[string]interface{})
	notFoundScriptsError := fmt.Errorf("can't any npm scripts")
	if !ok {
		return Package{}, notFoundScriptsError
	}
	for i, v := range _scripts {
		if v2, ok := v.(string); ok {
			scripts[i] = v2
		}
	}
	if len(scripts) == 0 {
		return Package{}, notFoundScriptsError
	}

	pkg := Package{
		RAW:     data,
		DIR:     filepath.Dir(f),
		ENV:     env,
		Scripts: scripts,
	}

	return pkg, nil
}

func (pkg Package) getEnv() []string {
	env := []string{}
	for i, v := range pkg.ENV {
		env = append(env, fmt.Sprintf("%v='%v'", i, v))
	}
	return env
}
