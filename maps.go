package generigo

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
