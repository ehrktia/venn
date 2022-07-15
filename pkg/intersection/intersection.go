package intersection

import "github.com/ehrktia/venn/internal"

// Intersect combines two different input slices a1 and a2
// provides unique result
func IntersectString(a1, a2 []string) []string {
	r := make([]string, 0, len(a1)+len(a2))
	uniqueA2 := internal.DeDuplicateString(a2)
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
	uniqA2 := internal.DeDuplicateint(a2)
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
	uniqA2 := internal.DeDuplicatefloat64(a2)
	lookUp := floatconvertToMap(a1)
	for _, data := range uniqA2 {
		if _, ok := lookUp[data]; ok {
			r = append(r, data)
		}
	}
	return r
}

func stringconvertToMap(a1 []string) internal.StringLookUp {
	r := make(internal.StringLookUp, len(a1))
	uniqA1 := internal.DeDuplicateString(a1)
	for _, data := range uniqA1 {
		r[data] = struct{}{}
	}
	return r

}

func intconvertToMap(a1 []int) internal.IntLookUp {
	r := make(internal.IntLookUp, len(a1))
	uniqA1 := internal.DeDuplicateint(a1)
	for _, data := range uniqA1 {
		r[data] = true
	}
	return r
}

func floatconvertToMap(a1 []float64) internal.FloatLookUp {
	r := make(internal.FloatLookUp, len(a1))
	uniqA1 := internal.DeDuplicatefloat64(a1)
	for _, data := range uniqA1 {
		r[data] = true
	}
	return r
}
