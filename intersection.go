// package venn provides user with option to combine and join two or more datasets
// in to a unified unique data set
package venn

// Intersect combines two different input slices a1 and a2
// provides unique result
func IntersectString(a1, a2 []string) []string {
	r := make([]string, 0, len(a1)+len(a2))
	uniqueA2 := deDuplicateString(a2)
	lookup := stringconvertToMap(a1)
	for _, data := range uniqueA2 {
		if _, ok := lookup[data]; ok {
			r = append(r, data)
		}
	}
	return r
}

// IntersectInt gives the list of common elements between two data set of int
// provides unique result
func IntersectInt(a1, a2 []int) []int {
	r := make([]int, 0, len(a1)+len(a2))
	uniqA2 := deDuplicateint(a2)
	lookUp := intconvertToMap(a1)
	for _, data := range uniqA2 {
		if _, ok := lookUp[data]; ok {
			r = append(r, data)
		}
	}
	return r
}

// IntersectFloat64 gives list of common elements between two data set of
// float64 provides a unique result
func IntersectFloat64(a1, a2 []float64) []float64 {
	r := make([]float64, 0, len(a1)+len(a2))
	uniqA2 := deDuplicatefloat64(a2)
	lookUp := floatconvertToMap(a1)
	for _, data := range uniqA2 {
		if _, ok := lookUp[data]; ok {
			r = append(r, data)
		}
	}
	return r
}

func stringconvertToMap(a1 []string) stringLookUp {
	r := make(stringLookUp, len(a1))
	uniqA1 := deDuplicateString(a1)
	for _, data := range uniqA1 {
		r[data] = struct{}{}
	}
	return r

}

func intconvertToMap(a1 []int) intLookUp {
	r := make(intLookUp, len(a1))
	uniqA1 := deDuplicateint(a1)
	for _, data := range uniqA1 {
		r[data] = true
	}
	return r
}

func floatconvertToMap(a1 []float64) floatLookUp {
	r := make(floatLookUp, len(a1))
	uniqA1 := deDuplicatefloat64(a1)
	for _, data := range uniqA1 {
		r[data] = true
	}
	return r
}
