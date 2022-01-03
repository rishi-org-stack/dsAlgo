package array

func InsertionSort(ar []int) {
	for i := 1; i < len(ar); i++ {
		for j := i; j < len(ar)-1; j++ {
			if ar[i] > ar[j+1] {
				ar[i], ar[j+1] = ar[j+1], ar[i]
			}
		}
	}
}
func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		iMin := i
		for j := i + 1; j < len(arr); j++ {
			if arr[iMin] > arr[j] {
				iMin = j
			}
		}

		arr[i], arr[iMin] = arr[iMin], arr[i]
	}
}
func BubbleSort(arr []int) {
	swapped := 1
	for i := 0; i < len(arr) && swapped == 1; i++ {
		swapped = 0
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped++
			}
		}

	}
}

func QuickSort(arr []int, f, l int) {
	if f >= l {
		return
	}

	partionIndex := f
	pivot := l - 1

	for i := f; i < l; i++ {
		if arr[i] < arr[pivot] {
			arr[partionIndex], arr[i] = arr[i], arr[partionIndex]
			partionIndex++

		}
	}
	arr[partionIndex], arr[pivot] = arr[pivot], arr[partionIndex]
	QuickSort(arr, 0, partionIndex-1)
	QuickSort(arr, partionIndex+1, l)
}
func mergeSorted(a []int, b []int) []int {
	aLen := len(a)
	bLen := len(b)
	aStart := 0
	bStart := 0
	temp := make([]int, 0)
	for aStart < aLen && bStart < bLen {
		if a[aStart] < b[bStart] {
			temp = append(temp, a[aStart])
			aStart++
		} else if a[aStart] > b[bStart] {
			temp = append(temp, b[bStart])
			bStart++
		}
	}
	if aStart != aLen-1 {
		temp = append(temp, a[aStart:]...)
	}

	if bStart != bLen-1 {
		temp = append(temp, b[bStart:]...)
	}
	return temp
}
