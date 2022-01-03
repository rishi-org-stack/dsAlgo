package array

import "fmt"

//find f from  array ar
func BinarySearch(ar []int, f, start, end int) (bool, int) {
	if start > end {
		return false, 0
	}

	mid := (start + end) / 2
	// fmt.Println(mid)
	if ar[mid] == f {
		return true, mid
	} else if f > ar[mid] {
		return BinarySearch(ar, f, mid+1, end)

	} else {
		return BinarySearch(ar, f, start, mid-1)

	}
}

// func CustomizedBinarySearch(ar []int, f func(int)bool , start, end int) (bool, int) {
// if start > end {
// return false, 0
// }
//
// mid := (start + end) / 2
// fmt.Println(mid)
// if ar[mid] == f {
// return true, mid
// } else if f > ar[mid] {
// return BinarySearch(ar, f, mid+1, end)
//
// } else {
// return BinarySearch(ar, f, start, mid-1)
//
// }
// }
//O(n)
func FirstRepeated(ar []int) int {
	res := 0
	occurence := make(map[int]int)
	for _, val := range ar {
		if occurence[val] == 1 {
			res = val
		} else {
			occurence[val]++
		}
	}
	fmt.Println(occurence)
	return res
}

func RepeatedValues(ar []int) []int {
	res := make([]int, 0)
	occurence := make(map[int]int)
	for _, val := range ar {
		if occurence[val] == 1 {
			// occurence[val]++
			res = append(res, val)
		} else {
			occurence[val]++
		}
	}
	fmt.Println(occurence)
	return res
}

func MostRepeatedValues(ar []int) int {
	repeatedValues := make([]int, 0)
	occurence := make(map[int]int)
	for _, val := range ar {
		if occurence[val] == 1 {
			// occurence[val]++
			repeatedValues = append(repeatedValues, val)
		} else {
			occurence[val]++
		}
	}

	maxtimes := 0
	res := 0
	for _, val := range repeatedValues {
		if occurence[val] > maxtimes {
			maxtimes = occurence[val]
			res = val
		}
	}
	return res
}
