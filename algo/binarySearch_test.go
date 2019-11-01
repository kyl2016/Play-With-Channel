package algo

import(
	"testing"
	"fmt"
	"errors"
)

func TestBinarySearch(t *testing.T){
	index, err := BinarySearch([]int{1,2,3,4,5}, 5)
	if err != nil {
		panic(err)
	}
if index != 4 {
	panic(err)
}

	fmt.Println(index)
}

func BinarySearch(s []int, target int) (index int, err error) {
	len := len(s)

	if len == 0 {
		return -1, errors.New("Not found")
	}

	return binarySearchInternally(s, 0, len-1, target)
}

func binarySearchInternally(s []int, left, right, target int) (index int, err error) {
	if left >= right {
		if s[left] == target {
			return left,nil
		}
		return -1, errors.New("Not found")
	}

	mid := left + (right-left)/2
	// mid := left + ((right-left)>>2)
	if s[mid] == target {
		return mid, nil
	}

	if s[mid] > target {
		return binarySearchInternally(s, left, mid-1, target)
	} 

	return binarySearchInternally(s, mid+1, right, target)
}

func BinarySearch2(s []int, target int)int{
	low := 0
	high := len(s) -1

	for {
		if low <= high {
			mid := low + ((high-low)>>2)
			if s[mid] == target{
				return mid
			}
			if s[mid]>target{
				high = mid-1
			} else {
				low = mid+1
			}
		}
	}

	return -1
}