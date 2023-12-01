package git

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Prepare_Result(t *testing.T) {
	type test struct {
		title           string
		inputSince      *time.Time
		inputLogDetails *LogDetails
		want            *VCSAnalysisInfo
	}

	now := time.Now()

	tests := []test{
		{
			title:           "Nil Since Param",
			inputSince:      nil,
			inputLogDetails: nil,
			want:            nil,
		},
		{
			title:           "Nil LogDetails Param",
			inputSince:      &now,
			inputLogDetails: nil,
			want:            nil,
		},
		{
			title:      "Successful case",
			inputSince: &now,
			inputLogDetails: &LogDetails{
				AuthorInfo: map[string]AuthorInfo{
					"foo": {
						Name:          "foo",
						Commits:       5,
						ModifiedLines: 2,
					},
					"bar": {
						Name:          "bar",
						Commits:       15,
						ModifiedLines: 7,
					},
				},
				PathModificationsInfo: map[string]ModificationsInfo{
					"pm1": {
						Source: "foobar",
						AuthorsSummary: []ContributorSummary{
							{
								Name:          "foobar",
								Modifications: 13,
							},
							{
								Name:          "barfoo",
								Modifications: 10,
							},
						},
					},
				},
				FileModificationsInfo: map[string]ModificationsInfo{
					"fm1": {
						Source: "foobar.go",
						AuthorsSummary: []ContributorSummary{
							{
								Name:          "foobar",
								Modifications: 13,
							},
							{
								Name:          "barfoo",
								Modifications: 10,
							},
						},
					},
					"fm2": {
						Source: "barfoo.go",
						AuthorsSummary: []ContributorSummary{
							{
								Name:          "barfoo",
								Modifications: 15,
							},
						},
					},
				},
			},
			want: &VCSAnalysisInfo{
				Since: &now,
				AuthorsInfo: AuthorsInfo{
					TotalAuthors:       2,
					TotalCommits:       20,
					TotalModifications: 9,
					AuthorDetails: []AuthorInfo{
						{
							Name:          "foo",
							Commits:       5,
							ModifiedLines: 2,
						},
						{
							Name:          "bar",
							Commits:       15,
							ModifiedLines: 7,
						},
					},
				},
				ModificationsInfo: Modifications{
					PathModifications: []ModificationsInfo{
						{
							Source:             "foobar",
							TotalModifications: 23,
							AuthorsSummary: []ContributorSummary{
								{
									Name:          "barfoo",
									Modifications: 10,
								},
								{
									Name:          "foobar",
									Modifications: 13,
								},
							},
						},
					},
					FileModifications: []ModificationsInfo{
						{
							Source:             "barfoo.go",
							TotalModifications: 15,
							AuthorsSummary: []ContributorSummary{
								{
									Name:          "barfoo",
									Modifications: 15,
								},
							},
						},
						{
							Source:             "foobar.go",
							TotalModifications: 23,
							AuthorsSummary: []ContributorSummary{
								{
									Name:          "barfoo",
									Modifications: 10,
								},
								{
									Name:          "foobar",
									Modifications: 13,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		got := prepareResult(tc.inputSince, tc.inputLogDetails)
		if !equalsVcsAnalysisInfo(t, tc.want, got) {
			t.Fatalf("For case: '%s' - expected: %v, got: %v", tc.title, tc.want, got)
		}
	}
}

func equalsVcsAnalysisInfo(t *testing.T, a *VCSAnalysisInfo, b *VCSAnalysisInfo) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil && b != nil {
		return false
	}
	if a != nil && b == nil {
		return false
	}
	if a.Since != b.Since {
		fmt.Printf("False1\n")
		return false
	}

	if !authorsEquals(t, a.AuthorsInfo, b.AuthorsInfo) {
		fmt.Printf("False2 [%+v][%+v]\n", a.AuthorsInfo, b.AuthorsInfo)
		return false
	}
	if !assert.ElementsMatch(t, a.ModificationsInfo.PathModifications, b.ModificationsInfo.PathModifications) {
		fmt.Printf("False3 [%+v][%+v]\n", a.ModificationsInfo.PathModifications, b.ModificationsInfo.PathModifications)
		return false
	}
	if !assert.ElementsMatch(t, a.ModificationsInfo.FileModifications, b.ModificationsInfo.FileModifications) {
		fmt.Printf("False4 [%+v][%+v]\n", a.ModificationsInfo.FileModifications, b.ModificationsInfo.FileModifications)
		return false
	}
	return true
}
