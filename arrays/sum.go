package main

func Sum(nums []int) int {
	res := 0
	for _, a := range nums {
		res += a
	}
	return res
}

func SumAll(nums ...[]int) []int {
	//l := len(nums)
	//res := make([]int, l)
	//for i, slice := range nums {
	//	res[i] = Sum(slice)
	//}
	res := []int{}
	for _, slice := range nums {
		res = append(res, Sum(slice))
	}
	return res
}

func SumAllTails(nums ...[]int) []int {
	l := len(nums)
	res := make([]int, l)
	for i, slice := range nums {
		res[i] = Sum(slice)
		if len(slice) > 0 {
			res[i] -= slice[0]
		}
	}
	return res
}
