package utils

func GetOnlyEven(arr []int) []int {
	evens := []int{}
	for i := 0; i < len(arr); i++ {
		if CheckForEven(arr[i]) {
			evens = append(evens, arr[i])
		}
	}
	return evens
}
