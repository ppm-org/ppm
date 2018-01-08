package utils

func ArrayHas(array []interface{}, index int, value interface{}) bool {
	return len(array) > index && array[index] == value
}
