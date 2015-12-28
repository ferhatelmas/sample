package sample

import "testing"

func TestFindIndexOfEqualOrClosestLower(t *testing.T) {
	for i, f := range [...]struct {
		s   []int64
		n   int64
		err error
		out int
	}{
		{[]int64{}, 1, errEmptySlice, 0},
		{[]int64{1}, 0, errInsufficientDataLow, 0},
		{[]int64{1}, 1, nil, 0},
		{[]int64{1}, 2, nil, 0},
		{[]int64{1, 3}, 0, errInsufficientDataLow, 0},
		{[]int64{1, 3}, 1, nil, 0},
		{[]int64{1, 3}, 2, nil, 0},
		{[]int64{1, 3}, 3, nil, 1},
		{[]int64{1, 3}, 4, nil, 1},
		{[]int64{1, 3, 5}, 0, errInsufficientDataLow, 0},
		{[]int64{1, 3, 5}, 1, nil, 0},
		{[]int64{1, 3, 5}, 2, nil, 0},
		{[]int64{1, 3, 5}, 3, nil, 1},
		{[]int64{1, 3, 5}, 4, nil, 1},
		{[]int64{1, 3, 5}, 5, nil, 2},
		{[]int64{1, 3, 5}, 6, nil, 2},
		{[]int64{1, 3, 5, 7}, 0, errInsufficientDataLow, 0},
		{[]int64{1, 3, 5, 7}, 1, nil, 0},
		{[]int64{1, 3, 5, 7}, 2, nil, 0},
		{[]int64{1, 3, 5, 7}, 3, nil, 1},
		{[]int64{1, 3, 5, 7}, 4, nil, 1},
		{[]int64{1, 3, 5, 7}, 5, nil, 2},
		{[]int64{1, 3, 5, 7}, 6, nil, 2},
		{[]int64{1, 3, 5, 7}, 7, nil, 3},
		{[]int64{1, 3, 5, 7}, 8, nil, 3},
		{[]int64{1, 3, 5, 7, 9}, 0, errInsufficientDataLow, 0},
		{[]int64{1, 3, 5, 7, 9}, 1, nil, 0},
		{[]int64{1, 3, 5, 7, 9}, 2, nil, 0},
		{[]int64{1, 3, 5, 7, 9}, 3, nil, 1},
		{[]int64{1, 3, 5, 7, 9}, 4, nil, 1},
		{[]int64{1, 3, 5, 7, 9}, 5, nil, 2},
		{[]int64{1, 3, 5, 7, 9}, 6, nil, 2},
		{[]int64{1, 3, 5, 7, 9}, 7, nil, 3},
		{[]int64{1, 3, 5, 7, 9}, 8, nil, 3},
		{[]int64{1, 3, 5, 7, 9}, 9, nil, 4},
		{[]int64{1, 3, 5, 7, 9}, 10, nil, 4},
		{[]int64{1, 3, 5, 7, 9, 11}, 0, errInsufficientDataLow, 0},
		{[]int64{1, 3, 5, 7, 9, 11}, 1, nil, 0},
		{[]int64{1, 3, 5, 7, 9, 11}, 2, nil, 0},
		{[]int64{1, 3, 5, 7, 9, 11}, 3, nil, 1},
		{[]int64{1, 3, 5, 7, 9, 11}, 4, nil, 1},
		{[]int64{1, 3, 5, 7, 9, 11}, 5, nil, 2},
		{[]int64{1, 3, 5, 7, 9, 11}, 6, nil, 2},
		{[]int64{1, 3, 5, 7, 9, 11}, 7, nil, 3},
		{[]int64{1, 3, 5, 7, 9, 11}, 8, nil, 3},
		{[]int64{1, 3, 5, 7, 9, 11}, 9, nil, 4},
		{[]int64{1, 3, 5, 7, 9, 11}, 10, nil, 4},
		{[]int64{1, 3, 5, 7, 9, 11}, 11, nil, 5},
		{[]int64{1, 3, 5, 7, 9, 11}, 12, nil, 5},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 0, errInsufficientDataLow, 0},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 1, nil, 0},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 2, nil, 0},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 3, nil, 1},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 4, nil, 1},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 5, nil, 2},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 6, nil, 2},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 7, nil, 3},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 8, nil, 3},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 9, nil, 4},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 10, nil, 4},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 11, nil, 5},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 12, nil, 5},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 13, nil, 6},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 14, nil, 6},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 15, nil, 7},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 16, nil, 7},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 17, nil, 8},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 18, nil, 8},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 19, nil, 9},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 20, nil, 9},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 21, nil, 10},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 22, nil, 10},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 0, errInsufficientDataLow, 0},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 1, nil, 0},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 2, nil, 0},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 3, nil, 1},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 4, nil, 1},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 5, nil, 2},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 6, nil, 2},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 7, nil, 3},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 8, nil, 3},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 9, nil, 4},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 10, nil, 4},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 11, nil, 5},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 12, nil, 5},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 13, nil, 6},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 14, nil, 6},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 15, nil, 7},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 16, nil, 7},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 17, nil, 8},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 18, nil, 8},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 19, nil, 9},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 20, nil, 9},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 21, nil, 10},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 22, nil, 10},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 23, nil, 11},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 24, nil, 11},
	} {
		got, err := findIndexOfEqualOrClosestLower(f.n, f.s)
		if f.err != nil {
			if err == nil {
				t.Errorf("%d) s=%v, n=%d, expected error %q, none returned, got=%d instead", i, f.s, f.n, f.err, got)
			}
			if err != f.err {
				t.Errorf("%d) s=%v, n=%d, expected error %q, but error %q was returned", i, f.s, f.n, f.err, err)
			}
			// correct error returned
		} else { // no error was expected
			if err != nil {
				t.Errorf("%d) s=%v, n=%d, out=%d, got unexpected error: %s", i, f.s, f.n, f.out, err)
			}
			if f.s[got] != f.s[f.out] {
				t.Errorf("%d) s=%v, n=%d, out=%d, wrong index returned: %d", i, f.s, f.n, f.out, got)
			}
			// nil error and correct output returned
		}
	}
}
