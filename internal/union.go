package internal

import (
	"sync"
)

// UnionAll combines two input slices a1 and a2
// retains original data from input
func UnionAll(a1, a2 []interface{}) []interface{} {
	var r []interface{}
	r = append(r, a1...)
	r = append(r, a2...)
	return r
}

// Union combines two different input slices a1 and a2
// provides unique result
func UnionString(a1, a2 []string) []string {
	r := []string{}
	r = append(r, a1...)
	r = append(r, a2...)
	return deDuplicateString(r)
}

func deDuplicateString(a1 []string) []string {
	r := []string{}
	wg := new(sync.WaitGroup)
	inpOutCh := make(chan string)
	lookup := make(stringLookUp, 0)
	// send data
	wg.Add(1)
	go func() {
		defer func() {
			close(inpOutCh)
			wg.Done()
		}()
		for _, data := range a1 {
			inpOutCh <- data
		}
	}()
	// remove duplicates
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		for v := range inpOutCh {
			if _, ok := lookup[v]; !ok {
				lookup[v] = true
				r = append(r, v)
			}
		}

	}()
	wg.Wait()
	// read inp
	return r
}

// UnionInt combines two different input slices a1 and a2 of type int
// provides unique result
func UnionInt(a1, a2 []int) []int {
	r := []int{}
	r = append(r, a1...)
	r = append(r, a2...)
	return deDuplicateint(r)
}

func deDuplicateint(a1 []int) []int {
	r := []int{}
	lookUp := map[int]bool{}
	wg := new(sync.WaitGroup)
	inCh := make(chan int)
	wg.Add(1)
	go func() {
		defer func() {
			close(inCh)
			wg.Done()
		}()
		for _, data := range a1 {
			inCh <- data
		}
	}()
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		for v := range inCh {
			if ok := lookUp[v]; !ok {
				lookUp[v] = true
				r = append(r, v)
			}
		}
	}()
	wg.Wait()
	return r
}

// UnionFloat64 combines two different slices a1 and a2 of type float64
// provides a unique result
func UnionFloat64(a1, a2 []float64) []float64 {
	r := []float64{}
	r = append(r, a1...)
	r = append(r, a2...)
	return deDuplicatefloat64(r)
}

func deDuplicatefloat64(a1 []float64) []float64 {
	r := []float64{}
	lookUp := map[float64]bool{}
	wg := new(sync.WaitGroup)
	inCh := make(chan float64)
	wg.Add(1)
	go func() {
		defer func() {
			close(inCh)
			wg.Done()
		}()
		for _, data := range a1 {
			inCh <- data
		}
	}()
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		for v := range inCh {
			if ok := lookUp[v]; !ok {
				lookUp[v] = true
				r = append(r, v)
			}
		}
	}()
	wg.Wait()
	return r
}
