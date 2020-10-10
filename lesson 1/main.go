package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	res := GetMapValuesSortedByKey(map[int]string{
		9:  "Сентябрь",
		1:  "Январь",
		2:  "Февраль",
		10: "Октябрь",
		5:  "Май",
		7:  "Июль",
		8:  "Август",
		12: "Декарь",
		3:  "Март",
		6:  "Июнь",
		4:  "Апрель",
		11: "Ноябрь",
	})
	fmt.Println(res)
}

// ReturnInt return int
func ReturnInt() int {
	var res int
	return res
}

//ReturnFloat return float
func ReturnFloat() float32 {
	var res float32
	return res
}

//ReturnIntArray return array
func ReturnIntArray() [3]int {
	arr := [3]int{1, 3, 4}
	return arr
}

//ReturnIntSlice return slice
func ReturnIntSlice() []int {
	expected := []int{1, 2, 3}
	return expected
}

//IntSliceToString convert int slice to string
func IntSliceToString(slice []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(slice), " ", "", -1), "[]")
}

//MergeSlices merge slices
func MergeSlices(slice1 []float32, slice2 []int32) []int {

	var slice3 []int

	for i := range slice1 {
		slice3 = append(slice3, int(slice1[i]))
	}

	for i := range slice2 {
		slice3 = append(slice3, int(slice2[i]))
	}

	return slice3
}

//GetMapValuesSortedByKey sort structure
func GetMapValuesSortedByKey(inputSruct map[int]string) []string {
	keys := make([]int, 0, len(inputSruct))
	for k := range inputSruct {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	string := make([]string, 0, len(inputSruct))

	for _, k := range keys {
		string = append(string, inputSruct[k])
	}

	return string
}
