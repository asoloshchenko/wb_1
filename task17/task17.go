package main

import "fmt"

func BinSearch(slice []int, el int) (int, bool) {
	low, high := 0, len(slice)-1
	var mid int
	for low <= high {
		mid = (low + high) / 2
		if slice[mid] == el {
			return mid, true
		}
		if slice[mid] > el {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return mid, false

}

func main() {
	testCases := []struct {
		slice []int
		el    int
		want  int
		found bool
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2, true},
		{[]int{1, 2, 3, 4, 5}, 6, 4, false},
		{[]int{1, 2, 3, 4, 5}, 1, 0, true},
		{[]int{1, 2, 3, 4, 5}, 0, 0, false},
	}

	for _, tc := range testCases {
		got, found := BinSearch(tc.slice, tc.el)
		if got != tc.want || found != tc.found {
			fmt.Printf("BinSearch(%v, %d) = %d, %v; want %d, %v\n",
				tc.slice, tc.el, got, found, tc.want, tc.found)
		}
	}
}
