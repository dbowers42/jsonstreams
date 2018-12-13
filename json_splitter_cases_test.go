package jsonstreams

const (
	j1 = "{\"message\": \"hi\"}"
	j2 = "{\"message\": \"bye\"}"
	j3 = "{\"greeter\": {\"name\": \"Rob\", \"message\": \"hi\"}}"
)

var testCases = []struct {
	description    string
	input          string
	expected       []string
	expectedLength int
}{
	{"single json body", j1, []string{j1}, 1},
	{"multiple json bodies", j1 + j2, []string{j1, j2}, 2},
	{"single json body with nested json", j3, []string{j3}, 1},
	{"multiple json bodies with nested json", j1 + j2 + j3, []string{j1, j2, j3}, 3},
	{"complex stream with mix json bodies", j1 + j2 + j1 + j2 + j3, []string{j1, j2, j1, j2, j3}, 5},
}
