# venn

![GitHub](https://img.shields.io/github/license/ehrktia/venn)
[![star this repo](https://githubbadges.com/star.svg?user=ehrktia&repo=venn&style=default)](https://github.com/ehrktia/venn)
[![fork this repo](https://githubbadges.com/fork.svg?user=ehrktia&repo=venn&style=default)](https://github.com/ehrktia/venn/fork)
![GitHub last commit](https://img.shields.io/github/last-commit/ehrktia/venn)
![GitHub contributors](https://img.shields.io/github/contributors/ehrktia/venn)
![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/ehrktia/venn/main)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/ehrktia/venn)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/ehrktia/venn/Test)
[![Go Reference](https://pkg.go.dev/badge/github.com/ehrktia/venn.svg)](https://pkg.go.dev/github.com/ehrktia/venn)




Venn follows unix principle and does only one thing i.e - union and intersection operations on data sets.

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
 - float64

## contributions 

This is a open source library , all pull request and feature request are welcome.


Thank you in advance for your time and efforts 



