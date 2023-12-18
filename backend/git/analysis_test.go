package git

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Analyze_Authors(t *testing.T) {
	type test struct {
		title string
		input map[string]AuthorInfo
		want  AuthorsInfo
	}

	tests := []test{
		{title: "Empty Map", input: map[string]AuthorInfo{}, want: AuthorsInfo{}},
		{
			title: "Authors Info Map", input: map[string]AuthorInfo{
				"case1": {Name: "foo", Commits: 10, ModifiedLines: 6},
				"case2": {Name: "bar", Commits: 34, ModifiedLines: 21},
				"case3": {Name: "foobar", Commits: 18, ModifiedLines: 11},
			},
			want: AuthorsInfo{
				TotalAuthors:       3,
				TotalCommits:       62,
				TotalModifications: 38,
				AuthorDetails: []AuthorInfo{
					{Name: "foo", Commits: 10, ModifiedLines: 6},
					{Name: "bar", Commits: 34, ModifiedLines: 21},
					{Name: "foobar", Commits: 18, ModifiedLines: 11},
				},
			},
		},
	}

	for _, tc := range tests {
		got := analyzeAuthors(tc.input)
		if !authorsEquals(t, tc.want, got) {
			t.Fatalf("For case: '%s' - expected: %v, got: %v", tc.title, tc.want, got)
		}
	}
}

func authorsEquals(t *testing.T, want AuthorsInfo, got AuthorsInfo) bool {
	a := want.TotalModifications == got.TotalModifications
	b := want.TotalAuthors == got.TotalAuthors
	c := want.TotalCommits == got.TotalCommits
	d := assert.ElementsMatch(t, want.AuthorDetails, got.AuthorDetails)

	return a && b && c && d
}

func Test_Analyze_Modifications(t *testing.T) {
	var emptyInfoSlice []ModificationsInfo
	type test struct {
		title string
		input map[string]ModificationsInfo
		want  []ModificationsInfo
	}

	tests := []test{
		{
			title: "Empty Modifications Map",
			input: map[string]ModificationsInfo{},
			want:  emptyInfoSlice,
		},
		{
			title: "Modifications Info Map", input: map[string]ModificationsInfo{
				"case1": {
					Source:             "foo",
					TotalModifications: 0,
					AuthorsSummary: []ContributorSummary{
						{Name: "user1", Modifications: 4},
						{Name: "user3", Modifications: 6},
					},
				},
				"case2": {
					Source:             "bar",
					TotalModifications: 0,
					AuthorsSummary: []ContributorSummary{
						{Name: "user2", Modifications: 12},
						{Name: "user3", Modifications: 20},
					},
				},
			},
			want: []ModificationsInfo{
				{
					Source:             "foo",
					TotalModifications: 10,
					AuthorsSummary: []ContributorSummary{
						{Name: "user1", Modifications: 4},
						{Name: "user3", Modifications: 6},
					},
				},
				{
					Source:             "bar",
					TotalModifications: 32,
					AuthorsSummary: []ContributorSummary{
						{Name: "user2", Modifications: 12},
						{Name: "user3", Modifications: 20},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		got := analyzeModifications(tc.input)

		if !assert.ElementsMatch(t, tc.want, got) {
			t.Fatalf("For case: '%s' - expected: %v, got: %v", tc.title, tc.want, got)
		}
	}
}
