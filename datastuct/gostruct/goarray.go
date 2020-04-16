package gostruct

func CopyArray(array []interface{}) []interface{} {
	copyArray := make([]interface{}, len(array))
	copy(copyArray, array)
	return copyArray
}

func CopyIntArray(array []int) []int {
	copyArray := make([]int, len(array))
	copy(copyArray, array)
	return copyArray
}