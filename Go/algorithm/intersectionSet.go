package algorithm

import (
	"reflect"
)

func IntersectionOfArray(firstArray, secondArray []int) (unicArray []int) {
	m := make(map[int]bool)

	for _, item := range firstArray {
		m[item] = true
	}

	for _, item := range secondArray {
		if _, ok := m[item]; ok {
			if !itemExists(unicArray,item){
			unicArray = append(unicArray, item)
			}
		}
	}
	return
}
func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)
	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}