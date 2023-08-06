package checks

import "testing"

func TestIsValidString(t *testing.T) {
	tests := []struct {
		args string
		want bool
	}{
		{"qwertyuuio", true},
		{"123456789", true},
		{"asdqwe5451232", true},
		{"1", true},
		{"", false},
		{"1564165....", true},
		{"!!!!....", true},
		{"@@@@@@aaa...", true},
		{"###%%%%...", true},
		{"+-()==", false},
		{"]][[", false},
		{"}}{{", false},
		{"....", true},
		{"++++", false},
		{"----", false},
		{"~~~~", false},
		{"~~~", false},
		{":::'''\"\"|||\\\\///", false},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			if got := isValidString(tt.args); got != tt.want {
				t.Errorf("f() = %v, want %v", got, tt.want)
			}
		})
	}
}
