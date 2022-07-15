package venn

/*Venn follows unix principle and does only one thing i.e - union and intersection operations on data sets.

## Union

Combines two data set together similar to the mathematical union operation `AUB` . It provides unique results by eliminating duplicates

## Intersection

compares two data sets and extracts the common elements between the two sets as result which is similar to mathematical set operation `ANB`

## example usage


```
package main

import (
	"fmt"

	"github.com/ehrktia/venn/pkg/intersection"
	"github.com/ehrktia/venn/pkg/union"
)

func main() {
	a := []string{"a", "b", "c"}
	b := []string{"a", "b", "c"}
	v := union.UnionString(a, b)
	fmt.Printf("%v\n", v)
	i:=intersection.IntersectString(a,b)
	fmt.Printf("%v\n", i)

}


```
There are dedicated union and intersection operations available for the following data types
 - string
 - Int
 - float64 */
