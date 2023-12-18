package git

import (
	"sort"
	"testing"
)

func Test_Sort_Modifications(t *testing.T) {
	originalModifications := []ModificationsInfo{
		{
			Source:             "source1",
			TotalModifications: 123,
			AuthorsSummary: []ContributorSummary{
				{Name: "c1", Modifications: 5},
				{Name: "c2", Modifications: 3},
				{Name: "c3", Modifications: 14},
			},
		},
		{
			Source:             "source2",
			TotalModifications: 34,
			AuthorsSummary: []ContributorSummary{
				{Name: "c1", Modifications: 5},
				{Name: "c2", Modifications: 3},
				{Name: "c3", Modifications: 14},
				{Name: "c4", Modifications: 7},
			},
		},
		{
			Source:             "source3",
			TotalModifications: 87,
			AuthorsSummary: []ContributorSummary{
				{Name: "c3", Modifications: 87},
			},
		},
		{
			Source:             "source4",
			TotalModifications: 90,
			AuthorsSummary: []ContributorSummary{
				{Name: "c1", Modifications: 5},
				{Name: "c3", Modifications: 85},
			},
		},
		{
			Source:             "source5",
			TotalModifications: 43,
			AuthorsSummary: []ContributorSummary{
				{Name: "c3", Modifications: 43},
			},
		},
	}

	expectedResult := []string{"source3", "source5", "source4", "source1", "source2"}

	sort.Sort(sortModificationsSliceByAuthors(originalModifications))
	sourcesResult := modificationsToSourceSlice(originalModifications)

	if !equalSlices(sourcesResult, expectedResult) {
		t.Fatal("Sorted slice doesn't match expected result.")
	}
}

func modificationsToSourceSlice(modifications []ModificationsInfo) []string {
	var sources []string
	for _, m := range modifications {
		sources = append(sources, m.Source)
	}
	return sources
}

func equalSlices(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
