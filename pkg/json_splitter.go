package pkg

func SplitJsonStream(stream string) []string {
	var payloads []string
	leftIndex := 0
	rightIndex := 0
	openBraces := 0

	for index, c := range stream {
		if c == '{' {
			if openBraces == 0 {
				leftIndex = index
			}

			openBraces++
		}

		if c == '}' {
			openBraces--

			if openBraces == 0 {
				rightIndex = index
			}
		}

		if openBraces == 0 && rightIndex > 0 {
			payloads = append(payloads, stream[leftIndex:rightIndex+1])
		}
	}

	return payloads
}
