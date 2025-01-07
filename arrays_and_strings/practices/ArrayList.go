package practices

import (
	"errors"
	"fmt"
)

// strings := []string{"Go", "is", "awesome"}
var strings = []string{"Go", "is", "awesome"}

func main() {
	printSlice(strings)
	fmt.Println(searchSlice(strings, "Go"))

	strings, removedBool, removedError := remove(strings, "Go")
	fmt.Printf("Remove was called and returned: %v, %v, %v", strings, removedBool, removedError)

}

func printSlice(inSlice []string) {
	for _, str := range inSlice {
		fmt.Println(str)
	}
}

func searchSlice(inSlice []string, key string) (int, error) {
	for index, searchKey := range inSlice {
		if searchKey == key {
			return index, nil
		}
	}
	return -1, errors.New("Key it not here")
}

func remove(inSlice []string, key string) ([]string, bool, error) {
	removed := false
	outSlice := []string{}

	for index, searchKey := range inSlice {
		if searchKey == key {
			outSlice = append(inSlice[:index], inSlice[index+1:]...)
			removed = true
			return outSlice, removed, nil
		}
	}
	return inSlice, removed, errors.New("String not found in slice")
}
