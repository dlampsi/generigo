package gogen

// StringInSlice Detect if string in slice of strings.
func StringInSlice(str string, slice []string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// NumInSlice Detect if number in slice of int.
func NumInSlice(number int, slice []int) bool {
	for _, s := range slice {
		if number == s {
			return true
		}
	}
	return false
}

// InterfaceInSlice Detect if interface in slice of interfaces.
func InterfaceInSlice(inf interface{}, slice []interface{}) bool {
	for _, s := range slice {
		if s == inf {
			return true
		}
	}
	return false
}

// CompareStringSlices Compares two string slices on equivalence.
// Elements do not need to be in the same position in the list.
func CompareStringSlices(sl1, sl2 []string) bool {
	if (sl1 == nil) != (sl2 == nil) {
		return false
	}
	if len(sl1) != len(sl2) {
		return false
	}
	for _, s := range sl1 {
		if !StringInSlice(s, sl2) {
			return false
		}
	}
	return true
}

// FullCompareStringSlices Compares two string slices on full equivalence.
// Element of the same should be in the same position in the list.
func FullCompareStringSlices(sl1, sl2 []string) bool {
	if (sl1 == nil) != (sl2 == nil) {
		return false
	}
	if len(sl1) != len(sl2) {
		return false
	}
	for i := range sl1 {
		if sl1[i] != sl2[i] {
			return false
		}
	}
	return true
}

// SliceInStringSlices Check if string slice in slice of string slices.
func SliceInStringSlices(slice []string, slices [][]string) bool {
	for _, s := range slices {
		if CompareStringSlices(slice, s) {
			return true
		}
	}
	return false
}

// IsOdd Check if number is odd.
func IsOdd(number int) bool {
	return !(number%2 == 0)
}
