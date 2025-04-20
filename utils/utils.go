package utils

/*
InArray Will return true or false if the value is in the array
needle = string
haystack = array
*/
func InArray(needle string, haystack []string) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}
	return false
}
