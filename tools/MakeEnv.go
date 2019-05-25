package tools

import (
	"fmt"
	"os"
)

type object map[string]interface{}

// MakeEnv generate env
func MakeEnv(data object, prefix string, env *(map[string]string)) {

	for i, v := range data {

		// 有些不需要转成环境变量的就直接设为 nil
		if v == nil {
			continue
		}

		index := prefix + "_" + i

		if v := os.Getenv(index); v != "" {
			continue
		}

		if t, ok := v.(map[string]interface{}); ok {
			MakeEnv(t, index, env)
			continue
		}

		(*env)[index] = fmt.Sprint(v)

	}

}
