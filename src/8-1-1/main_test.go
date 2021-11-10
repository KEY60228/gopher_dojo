package mypkg_test

import (
	"dojo/8-1-1/mypkg"
	"testing"
)

func Test_IsEven(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		in   int
		want bool
	}{
		"+even": {4, true},
		"+odd":  {5, false},
		"-even": {-4, true},
		"-odd":  {5, false},
		"zero":  {0, false},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := mypkg.IsEven(tt.in); tt.want != got {
				t.Errorf("want IsEven(%d) = %v, got %v", tt.in, tt.want, got)
			}
		})
	}
}
