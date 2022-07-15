package venn_test

import (
	"fmt"

	"github.com/ehrktia/venn"
)

func ExampleUnionString() {
	a1 := []string{"a", "b", "c"}
	a2 := []string{"a", "b", "c"}
	combine := func(a1, a2 []string) {
		result := venn.UnionString(a1, a2)
		fmt.Printf("%v\n", result)
	}
	combine(a1, a2)
	// Output
	// UnionString([]string{"a","b","c"},[]string{"a","b","c"}) =
	// []string{"a","b","c"}
}

func ExampleIntersectString() {
	a1 := []string{"a", "b", "c"}
	a2 := []string{"a", "b", "c"}
	join := func(a1, a2 []string) {
		result := venn.IntersectString(a1, a2)
		fmt.Printf("%v\n", result)
	}
	join(a1, a2)
	// Output
	// IntersectString([]string{"a","b","c"},[]string{"a","b","c"}) =
	// []string{"a","b","c"}
}
