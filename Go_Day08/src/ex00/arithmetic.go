package arithmetic

import (
	"fmt"
	"unsafe"
)

func getElement(arr []int, idx int) (int, error) {
	if len(arr) == 0 {
		err := fmt.Errorf("slice is empty")
		return 0, err
	}
	if idx > len(arr)-1 {
		err := fmt.Errorf("index %d - out of range", idx)
		return 0, err
	}
	if idx < 0 {
		err := fmt.Errorf("index %d - negative", idx)
		return 0, err
	}

	for range idx {
		arr = arr[1:]
	}

	answer := arr[0]

	return answer, nil
}

func getElementUnsafePtr(arr []int, idx int) (int, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("slice is empty")
	}
	if idx < 0 {
		return 0, fmt.Errorf("index %d - negative", idx)
	}
	if idx >= len(arr) {
		return 0, fmt.Errorf("index %d - out of range", idx)
	}

	base := unsafe.Pointer(&arr[0])
	size := unsafe.Sizeof(arr[0])
	
	elementPtr := unsafe.Pointer(uintptr(base) + uintptr(idx)*size)
	value := *(*int)(elementPtr)

	return value, nil
}
