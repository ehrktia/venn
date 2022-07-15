package venn

// UnionAll combines two input slices a1 and a2
// retains original data from input
func UnionAll(a1, a2 []interface{}) []interface{} {
	r := make([]interface{}, 0, len(a1)+len(a2))
	r = append(r, a1...)
	r = append(r, a2...)
	return r
}

// Union combines two different input slices a1 and a2
// provides unique result
func UnionString(a1, a2 []string) []string {
	r := make([]string, 0, len(a1)+len(a2))
	r = append(append(r, a1...), a2...)
	return deDuplicateString(r)
}

// UnionInt combines two different input slices a1 and a2 of type int
// provides unique result
func UnionInt(a1, a2 []int) []int {
	r := make([]int, 0, len(a1)+len(a2))
	r = append(r, a1...)
	r = append(r, a2...)
	return deDuplicateint(r)
}

// UnionFloat64 combines two different slices a1 and a2 of type float64
// provides a unique result
func UnionFloat64(a1, a2 []float64) []float64 {
	r := make([]float64, 0, len(a1)+len(a2))
	r = append(r, a1...)
	r = append(r, a2...)
	return deDuplicatefloat64(r)
}
