package internal

import "testing"

func Test_Union(t *testing.T) {
	t.Run("union string successfully", func(t *testing.T) {
		tt := []struct {
			Name   string
			i1     []string
			i2     []string
			expLen int
		}{
			{
				Name:   "equal string list",
				i1:     []string{"1", "2"},
				i2:     []string{"1", "2"},
				expLen: 2,
			},
			{
				Name:   "non eq str list",
				i1:     []string{"1", "2", "3"},
				i2:     []string{"1", "2"},
				expLen: 3,
			},
			{
				Name:   "non eq str list",
				i1:     []string{},
				i2:     []string{},
				expLen: 0,
			},
		}
		for _, tc := range tt {
			t.Run(tc.Name, func(t *testing.T) {
				got := UnionString(tc.i1, tc.i2)
				if len(got) != tc.expLen {
					t.Fatalf("exp:%v,got:%v", tc.expLen, len(got))
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
		}{
			{
				Name:   "equal int list",
				i1:     []int{1, 2},
				i2:     []int{1, 2},
				expLen: 2,
			},
			{
				Name:   "non eq int list",
				i1:     []int{1, 2, 3},
				i2:     []int{1, 2},
				expLen: 3,
			},
			{
				Name:   "non eq int list",
				i1:     []int{},
				i2:     []int{},
				expLen: 0,
			},
		}
		for _, tc := range tt {
			t.Run(tc.Name, func(t *testing.T) {
				got := UnionInt(tc.i1, tc.i2)
				if len(got) != tc.expLen {
					t.Fatalf("exp:%v,got:%v", tc.expLen, len(got))
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
		}{
			{
				Name:   "equal float list",
				i1:     []float64{1, 2},
				i2:     []float64{1, 2},
				expLen: 2,
			},
			{
				Name:   "non eq float list",
				i1:     []float64{1, 2, 3},
				i2:     []float64{1, 2},
				expLen: 3,
			},
			{
				Name:   "empty float list",
				i1:     []float64{},
				i2:     []float64{},
				expLen: 0,
			},
		}
		for _, tc := range tt {
			t.Run(tc.Name, func(t *testing.T) {
				got := UnionFloat64(tc.i1, tc.i2)
				if len(got) != tc.expLen {
					t.Fatalf("exp:%v,got:%v", tc.expLen, len(got))
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
