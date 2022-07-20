package venn

import (
	"sync"
)

func deDuplicateString(a1 []string) []string {
	r := make([]string, 0, len(a1))
	wg := new(sync.WaitGroup)
	inpOutCh := make(chan string)
	lookup := make(stringLookUp)
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
		// TODO: most expensive operation find alternative
		for v := range inpOutCh {
			if _, ok := lookup[v]; !ok {
				lookup[v] = struct{}{}
				r = append(r, v)
			}
		}

	}()
	wg.Wait()
	// read inp
	return r
}

func deDuplicateint(a1 []int) []int {
	r := make([]int, 0, len(a1))
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

func deDuplicatefloat64(a1 []float64) []float64 {
	r := make([]float64, 0, len(a1))
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
