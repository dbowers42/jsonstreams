package pkg

import "testing"

func TestSplitJsonStream(t *testing.T) {
	for _, tt := range testCases {
		actual := SplitJsonStream(tt.input)

		if len(actual) != tt.expectedLength {
			t.Fatalf("Expected length of %d. but got %d", tt.expectedLength, len(actual))
		}

		for i := 0; i < len(tt.expected); i++ {
			if actual[i] != tt.expected[i] {
				t.Fatalf("Expected %s, but got %s", tt.expected[i], actual[i])
			}
		}
	}
}

func BenchmarkHey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			SplitJsonStream(tt.input)
		}
	}
}