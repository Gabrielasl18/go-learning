package codewars

func Pyramid(n int) [][]int {
	var slice2 [][]int
	if n != 0 {
		var subSlice []int
		for i := 1; i <= n; i++ {
			subSlice = append(subSlice, 1)
			slice2 = append(slice2, subSlice[:i])
		}
	} else {
		return [][]int{}
	}
	return slice2
}
