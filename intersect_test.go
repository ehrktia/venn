package venn

import (
	"reflect"
	"testing"
)

func Test_intersect_string(t *testing.T) {
	tt := []struct {
		name   string
		i1     []string
		i2     []string
		expLen int
	}{
		{
			name:   "eq len str list with common",
			i1:     []string{"a", "ab"},
			i2:     []string{"a", "c"},
			expLen: 1,
		},
		{
			name:   "no eq inp with no common",
			i1:     []string{"a", "ab"},
			i2:     []string{"c"},
			expLen: 0,
		},
		{
			name:   "nil inps",
			i1:     nil,
			i2:     nil,
			expLen: 0,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := IntersectString(tc.i1, tc.i2)
			if len(got) != tc.expLen {
				t.Fatalf("exp:%v,got:%v", tc.expLen, len(got))
			}
		})
	}
}

func Test_convert_map(t *testing.T) {
	tt := []struct {
		name   string
		inp    []string
		exp    stringLookUp
		expLen int
	}{
		{
			name:   "valid no dup string list",
			inp:    []string{"a", "b", "c"},
			expLen: 3,
			exp:    stringLookUp{"a": {}, "b": {}, "c": {}},
		},
		{
			name:   "nil input",
			inp:    nil,
			expLen: 0,
			exp:    stringLookUp{},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := stringconvertToMap(tc.inp)
			if len(got) != tc.expLen {
				t.Fatalf("exp:%v,got:%v", tc.expLen, len(got))
			}
			if !reflect.DeepEqual(got, tc.exp) {
				t.Fatalf("exp:%v,got:%v", tc.exp, got)

			}
		})
	}

}
func Test_intersect_int(t *testing.T) {
	tt := []struct {
		name   string
		i1     []int
		i2     []int
		exp    []int
		expLen int
	}{
		{
			name:   "eq len valid inp",
			i1:     []int{1, 2, 3},
			i2:     []int{1},
			exp:    []int{1},
			expLen: 1,
		},
		{
			name:   "no eq no intersect inp",
			i1:     []int{1, 2, 3},
			i2:     []int{4},
			exp:    []int{},
			expLen: 0,
		},
		{
			name:   "nil inp",
			i1:     nil,
			i2:     nil,
			exp:    []int{},
			expLen: 0,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := IntersectInt(tc.i1, tc.i2)
			if len(got) != tc.expLen {
				t.Fatalf("exp:%v,got:%v", tc.expLen, len(got))
			}
			if !reflect.DeepEqual(got, tc.exp) {
				t.Fatalf("exp:%v,got:%v", tc.exp, got)

			}
		})
	}

}
func Test_intersect_float(t *testing.T) {
	tt := []struct {
		name   string
		i1     []float64
		i2     []float64
		exp    []float64
		expLen int
	}{
		{
			name:   "eq len valid inp",
			i1:     []float64{1, 2, 3},
			i2:     []float64{1},
			exp:    []float64{1},
			expLen: 1,
		},
		{
			name:   "no eq no float64ersect inp",
			i1:     []float64{1, 2, 3},
			i2:     []float64{4},
			exp:    []float64{},
			expLen: 0,
		},
		{
			name:   "nil inp",
			i1:     nil,
			i2:     nil,
			exp:    []float64{},
			expLen: 0,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := IntersectFloat64(tc.i1, tc.i2)
			if len(got) != tc.expLen {
				t.Fatalf("exp:%v,got:%v", tc.expLen, len(got))
			}
			if !reflect.DeepEqual(got, tc.exp) {
				t.Fatalf("exp:%v,got:%v", tc.exp, got)

			}
		})
	}

}
