package router

import "testing"

func TestIsDomain(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"openai.com", true},
		{"example.org", true},
		{"example.edu", true},
		{"example.net", false},
		{"https://openai.com", false},
		{"open ai.com", false},
		{"", false},
	}

	for _, tc := range cases {
		if got := IsDomain(tc.input); got != tc.want {
			t.Fatalf("IsDomain(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}
