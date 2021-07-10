package generigo

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Copy map into new map.
func CopyMap(sourceMap map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})
	for k, v := range sourceMap {
		vm, ok := v.(map[string]interface{})
		if ok {
			newMap[k] = CopyMap(vm)
		} else {
			newMap[k] = v
		}
	}
	return newMap
}

// Copy map of string into new map.
func CopyMapString(sourceMap map[string]string) map[string]string {
	newMap := make(map[string]string)
	for k, v := range sourceMap {
		newMap[k] = v
	}
	return newMap
}

// Converts map of string into string in json format.
func JsonMapToString(m map[string]string) (string, error) {
	var result string
	if len(m) > 0 {
		resBytes, err := json.Marshal(m)
		if err != nil {
			return "", err
		}
		result = string(resBytes)
	}
	return result, nil
}

// Converts string (in json format) to map of strings.
func StringToJsonMap(s string) (map[string]string, error) {
	var result map[string]string
	if s != "" {
		if err := json.Unmarshal([]byte(s), &result); err != nil {
			return nil, err
		}
	}
	if result == nil {
		result = map[string]string{}
	}
	return result, nil
}
