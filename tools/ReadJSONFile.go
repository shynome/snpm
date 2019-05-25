package tools

import (
	"encoding/json"
	"io/ioutil"
)

// ReadJSONFile read json file to map[string]interface{}
func ReadJSONFile(f string) map[string]interface{} {
	file, _ := ioutil.ReadFile(f)
	data := map[string]interface{}{}
	json.Unmarshal(file, &data)
	return data
}
