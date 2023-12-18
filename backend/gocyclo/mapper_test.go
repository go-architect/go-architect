package gocyclo_test

import (
	"github.com/go-architect/go-architect/backend/gocyclo"
	"reflect"
	"testing"
)

func Test_GoCyclo_Mapper(t *testing.T) {
	type test struct {
		title string
		input string
		want  *gocyclo.GoCycloOutput
	}

	tests := []test{
		{title: "Unexpected output", input: "gjsdflgsdfkgjklsdfjgkldfsjgdgkdfjglk", want: nil},
		{title: "Empty output", input: "", want: &gocyclo.GoCycloOutput{}},
		{title: "Expected output", input: `5 main mainFunction main.go:73:1
5 foo foobar1 foo/functions.go:54:1
5 foo foobar2 foo/functions.go:28:1
4 main (*App).CheckEnvironmentPath app.go:39:1
3 main main main.go:27:1
Average: 1.96`, want: &gocyclo.GoCycloOutput{
			AverageComplexity: 1.96,
			ComplexityDetails: []*gocyclo.ComplexityDetail{
				{Package: "main", File: "main.go:73:1", Function: "mainFunction", Complexity: 5},
				{Package: "foo", File: "functions.go:54:1", Function: "foobar1", Complexity: 5},
				{Package: "foo", File: "functions.go:28:1", Function: "foobar2", Complexity: 5},
				{Package: "main", File: "app.go:39:1", Function: "(*App).CheckEnvironmentPath", Complexity: 4},
				{Package: "main", File: "main.go:27:1", Function: "main", Complexity: 3},
			},
		}},
	}

	for _, tc := range tests {
		got := gocyclo.MapToGoCycloModel(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("For case: '%s' - expected: %v, got: %v", tc.title, tc.want, got)
		}
	}
}
