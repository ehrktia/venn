package union

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/ehrktia/venn/internal"
)

func Test_Union(t *testing.T) {
	t.Run("union string successfully", func(t *testing.T) {
		tt := []struct {
			Name   string
			i1     []string
			i2     []string
			expLen int
			exp    []string
		}{
			{
				Name:   "equal string list",
				i1:     []string{"1", "2"},
				i2:     []string{"1", "2"},
				expLen: 2,
				exp:    []string{"1", "2"},
			},
			{
				Name:   "non eq str list",
				i1:     []string{"1", "2", "3"},
				i2:     []string{"1", "2"},
				expLen: 3,
				exp:    []string{"1", "2", "3"},
			},
			{
				Name:   "non eq str list",
				i1:     []string{},
				i2:     []string{},
				expLen: 0,
				exp:    []string{},
			},
		}
		for _, tc := range tt {
			t.Run(tc.Name, func(t *testing.T) {
				got := UnionString(tc.i1, tc.i2)
				t.Logf("got:%v\n", got)
				if len(got) != tc.expLen {
					t.Fatalf("exp:%v,got:%v", tc.expLen, len(got))
				}
				if !reflect.DeepEqual(got, tc.exp) {
					t.Fatalf("exp:%v,got:%v", tc.exp, got)
				}
			})
		}
	})
	t.Run("union int successfully", func(t *testing.T) {
		tt := []struct {
			Name   string
			i1     []int
			i2     []int
			expLen int
			exp    []int
		}{
			{
				Name:   "equal int list",
				i1:     []int{1, 2},
				i2:     []int{1, 2},
				expLen: 2,
				exp:    []int{1, 2},
			},
			{
				Name:   "non eq int list",
				i1:     []int{1, 2, 3},
				i2:     []int{1, 2},
				expLen: 3,
				exp:    []int{1, 2, 3},
			},
			{
				Name:   "non eq int list",
				i1:     []int{},
				i2:     []int{},
				expLen: 0,
				exp:    []int{},
			},
		}
		for _, tc := range tt {
			t.Run(tc.Name, func(t *testing.T) {
				got := UnionInt(tc.i1, tc.i2)
				if len(got) != tc.expLen {
					t.Fatalf("exp:%v,got:%v", tc.expLen, len(got))
				}
				if !reflect.DeepEqual(got, tc.exp) {
					t.Fatalf("exp:%v,got:%v", tc.exp, got)
				}
			})
		}
	})
	t.Run("union float64 successfully", func(t *testing.T) {
		tt := []struct {
			Name   string
			i1     []float64
			i2     []float64
			expLen int
			exp    []float64
		}{
			{
				Name:   "equal float list",
				i1:     []float64{1, 2},
				i2:     []float64{1, 2},
				expLen: 2,
				exp:    []float64{1, 2},
			},
			{
				Name:   "non eq float list",
				i1:     []float64{1, 2, 3},
				i2:     []float64{1, 2},
				expLen: 3,
				exp:    []float64{1, 2, 3},
			},
			{
				Name:   "empty float list",
				i1:     []float64{},
				i2:     []float64{},
				expLen: 0,
				exp:    []float64{},
			},
		}
		for _, tc := range tt {
			t.Run(tc.Name, func(t *testing.T) {
				got := UnionFloat64(tc.i1, tc.i2)
				if len(got) != tc.expLen {
					t.Fatalf("exp:%v,got:%v", tc.expLen, len(got))
				}
				if !reflect.DeepEqual(got, tc.exp) {
					t.Fatalf("exp:%v,got:%v", tc.exp, got)

				}
			})
		}
	})
	t.Run("union interface successfully", func(t *testing.T) {
		tt := []struct {
			Name   string
			i1     []interface{}
			i2     []interface{}
			expLen int
		}{
			{
				Name:   "eq interface diff data types",
				i1:     []interface{}{"1", 2, "3"},
				i2:     []interface{}{"1", 2, "3"},
				expLen: 6,
			},
			{
				Name:   "non interface diff data types",
				i1:     []interface{}{"1", 2, "3"},
				i2:     []interface{}{2, "3"},
				expLen: 5,
			},
		}
		for _, tc := range tt {
			t.Run(tc.Name, func(t *testing.T) {
				got := UnionAll(tc.i1, tc.i2)
				if len(got) != tc.expLen {
					t.Fatalf("exp:%v,got:%v", tc.expLen, len(got))
				}
			})
		}
	})
}

func Test_dedupString(t *testing.T) {
	exp := []string{"1", "2"}
	got := internal.DeDuplicateString(exp)
	if !reflect.DeepEqual(got, exp) {
		t.Fatalf("got:%v,exp:%v", got, exp)
	}
}

func Test_dedupInt(t *testing.T) {
	exp := []int{1, 2}
	got := internal.DeDuplicateint(exp)
	if !reflect.DeepEqual(got, exp) {
		t.Fatalf("exp:%v,got:%v", exp, got)
	}

}

func Test_dedupFloat(t *testing.T) {
	exp := []float64{1, 2}
	got := internal.DeDuplicatefloat64(exp)
	if !reflect.DeepEqual(got, exp) {
		t.Fatalf("exp:%v,got:%v", exp, got)
	}
}

// two inp elements
// goos: linux
// goarch: amd64
// pkg: github.com/ehrktia/venn/internal
// cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
// BenchmarkDeDupString-8            360528              7814 ns/op

// 50 inp elements
// goos: linux
// goarch: amd64
// pkg: github.com/ehrktia/venn/internal
// cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
// BenchmarkDeDupString-8             15994             63850 ns/op

// 100 inp elements
// goos: linux
// goarch: amd64
// pkg: github.com/ehrktia/venn/internal
// cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
// BenchmarkDeDupString-8             10000            134634 ns/op
func BenchmarkDeDupString(b *testing.B) {
	in := make([]string, 0, 100)
	for i := 1; i <= 100; i++ {
		in = append(in, strconv.Itoa(i))
	}
	for i := 0; i <= b.N; i++ {
		_ = internal.DeDuplicateString(in)
	}
}
