package arrays_and_slices

func Sum(a []int) (sum int) {
	for _, v := range a {
		sum += v
	}
	return
}

func SumAll(a ...[]int) (ret []int) {
	for _, v := range a {
		ret = append(ret, Sum(v))
	}
	return
}
func SumTails(a ...[]int) (ret []int) {
	for _, v := range a {
		if len(v) == 0 {
			ret = append(ret, 0)
			continue
		}
		tail := v[1:]
		ret = append(ret, Sum(tail))
	}
	return
}
