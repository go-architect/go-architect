package git

import (
	"fmt"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_Helpers_ResolveAuthorsInfo(t *testing.T) {
	type test struct {
		title      string
		fileStats  object.FileStats
		authorName string
		authorsMap map[string]AuthorInfo
		want       map[string]AuthorInfo
	}

	tests := []test{
		{
			title:      "Empty FileStats",
			fileStats:  object.FileStats{},
			authorName: "foobar",
			authorsMap: make(map[string]AuthorInfo),
			want: map[string]AuthorInfo{
				"foobar": {Name: "foobar", Commits: 1, ModifiedLines: 0},
			},
		},
		{
			title: "Modifications Case 1 - from empty authors map",
			fileStats: object.FileStats{
				{Name: "fs01", Addition: 5, Deletion: 3},
				{Name: "fs02", Addition: 15, Deletion: 33},
			},
			authorName: "foobar",
			authorsMap: make(map[string]AuthorInfo),
			want: map[string]AuthorInfo{
				"foobar": {Name: "foobar", Commits: 1, ModifiedLines: 56},
			},
		},
		{
			title: "Modifications Case 2 - from existent authors map without current author",
			fileStats: object.FileStats{
				{Name: "fs01", Addition: 5, Deletion: 3},
				{Name: "fs02", Addition: 15, Deletion: 33},
			},
			authorName: "foobar",
			authorsMap: map[string]AuthorInfo{
				"testing": {Name: "testing", Commits: 5, ModifiedLines: 100},
				"barfoo":  {Name: "barfoo", Commits: 40, ModifiedLines: 3},
			},
			want: map[string]AuthorInfo{
				"testing": {Name: "testing", Commits: 5, ModifiedLines: 100},
				"foobar":  {Name: "foobar", Commits: 1, ModifiedLines: 56},
				"barfoo":  {Name: "barfoo", Commits: 40, ModifiedLines: 3},
			},
		},
		{
			title: "Modifications Case 3 - from existent authors map with current author",
			fileStats: object.FileStats{
				{Name: "fs01", Addition: 5, Deletion: 3},
				{Name: "fs02", Addition: 15, Deletion: 33},
			},
			authorName: "foobar",
			authorsMap: map[string]AuthorInfo{
				"testing": {Name: "testing", Commits: 5, ModifiedLines: 100},
				"foobar":  {Name: "foobar", Commits: 10, ModifiedLines: 1000},
				"barfoo":  {Name: "barfoo", Commits: 40, ModifiedLines: 3},
			},
			want: map[string]AuthorInfo{
				"testing": {Name: "testing", Commits: 5, ModifiedLines: 100},
				"foobar":  {Name: "foobar", Commits: 11, ModifiedLines: 1056},
				"barfoo":  {Name: "barfoo", Commits: 40, ModifiedLines: 3},
			},
		},
	}

	for _, tc := range tests {
		resolveAuthorsInfo(tc.fileStats, tc.authorName, tc.authorsMap)
		if !assert.Equal(t, tc.want, tc.authorsMap) {
			t.Fatalf("For case: '%s' - expected: %v, got: %v", tc.title, tc.want, tc.authorsMap)
		}
	}
}

func Test_Helpers_ExtractPathModifications(t *testing.T) {
	type test struct {
		title         string
		fileStat      object.FileStat
		authorName    string
		modifications map[string]ModificationsInfo
		want          map[string]ModificationsInfo
	}

	tests := []test{
		{
			title:         "Empty FileStat",
			fileStat:      object.FileStat{},
			authorName:    "foobar",
			modifications: make(map[string]ModificationsInfo),
			want: map[string]ModificationsInfo{
				fmt.Sprintf(".%s", string(os.PathSeparator)): {
					Source:             fmt.Sprintf(".%s", string(os.PathSeparator)),
					TotalModifications: 0,
					AuthorsSummary: []ContributorSummary{
						{
							Name:          "foobar",
							Modifications: 0,
						},
					},
				},
			},
		},
		{
			title: "Single element FileStat",
			fileStat: object.FileStat{
				Name:     fmt.Sprintf("blabla%sfs01", string(os.PathSeparator)),
				Addition: 5,
				Deletion: 7,
			},
			authorName:    "foobar",
			modifications: make(map[string]ModificationsInfo),
			want: map[string]ModificationsInfo{
				fmt.Sprintf("blabla%s", string(os.PathSeparator)): {
					Source:             fmt.Sprintf("blabla%s", string(os.PathSeparator)),
					TotalModifications: 0,
					AuthorsSummary: []ContributorSummary{
						{
							Name:          "foobar",
							Modifications: 12,
						},
					},
				},
			},
		},
		{
			title: "Path exists in modifications map",
			fileStat: object.FileStat{
				Name:     fmt.Sprintf("blabla%sfs01", string(os.PathSeparator)),
				Addition: 5,
				Deletion: 7,
			},
			authorName: "foobar",
			modifications: map[string]ModificationsInfo{
				fmt.Sprintf("blabla%s", string(os.PathSeparator)): {
					Source:             fmt.Sprintf("blabla%s", string(os.PathSeparator)),
					TotalModifications: 21,
					AuthorsSummary: []ContributorSummary{
						{
							Name:          "jojojo",
							Modifications: 21,
						},
					},
				},
			},
			want: map[string]ModificationsInfo{
				fmt.Sprintf("blabla%s", string(os.PathSeparator)): {
					Source:             fmt.Sprintf("blabla%s", string(os.PathSeparator)),
					TotalModifications: 21,
					AuthorsSummary: []ContributorSummary{
						{
							Name:          "jojojo",
							Modifications: 21,
						},
						{
							Name:          "foobar",
							Modifications: 12,
						},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		extractPathModifications(tc.authorName, tc.fileStat, tc.modifications)
		if !assert.Equal(t, tc.want, tc.modifications) {
			t.Fatalf("For case: '%s' - expected: %v, got: %v", tc.title, tc.want, tc.modifications)
		}
	}
}

func Test_Helpers_ExtractFileModifications(t *testing.T) {
	type test struct {
		title         string
		fileStat      object.FileStat
		authorName    string
		modifications map[string]ModificationsInfo
		want          map[string]ModificationsInfo
	}

	tests := []test{
		{
			title:         "Empty FileStat",
			fileStat:      object.FileStat{},
			authorName:    "foobar",
			modifications: make(map[string]ModificationsInfo),
			want: map[string]ModificationsInfo{
				"": {
					Source:             "",
					TotalModifications: 0,
					AuthorsSummary: []ContributorSummary{
						{
							Name:          "foobar",
							Modifications: 0,
						},
					},
				},
			},
		},
		{
			title: "Single element FileStat",
			fileStat: object.FileStat{
				Name:     fmt.Sprintf("blabla%sfs01", string(os.PathSeparator)),
				Addition: 5,
				Deletion: 7,
			},
			authorName:    "foobar",
			modifications: make(map[string]ModificationsInfo),
			want: map[string]ModificationsInfo{
				fmt.Sprintf("blabla%sfs01", string(os.PathSeparator)): {
					Source:             fmt.Sprintf("blabla%sfs01", string(os.PathSeparator)),
					TotalModifications: 0,
					AuthorsSummary: []ContributorSummary{
						{
							Name:          "foobar",
							Modifications: 12,
						},
					},
				},
			},
		},
		{
			title: "Path exists in modifications map",
			fileStat: object.FileStat{
				Name:     fmt.Sprintf("blabla%sfs01", string(os.PathSeparator)),
				Addition: 5,
				Deletion: 7,
			},
			authorName: "foobar",
			modifications: map[string]ModificationsInfo{
				fmt.Sprintf("blabla%sfs01", string(os.PathSeparator)): {
					Source:             fmt.Sprintf("blabla%sfs01", string(os.PathSeparator)),
					TotalModifications: 21,
					AuthorsSummary: []ContributorSummary{
						{
							Name:          "jojojo",
							Modifications: 21,
						},
					},
				},
			},
			want: map[string]ModificationsInfo{
				fmt.Sprintf("blabla%sfs01", string(os.PathSeparator)): {
					Source:             fmt.Sprintf("blabla%sfs01", string(os.PathSeparator)),
					TotalModifications: 21,
					AuthorsSummary: []ContributorSummary{
						{
							Name:          "jojojo",
							Modifications: 21,
						},
						{
							Name:          "foobar",
							Modifications: 12,
						},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		extractFileModifications(tc.authorName, tc.fileStat, tc.modifications)
		if !assert.Equal(t, tc.want, tc.modifications) {
			t.Fatalf("For case: '%s' - expected: %v, got: %v", tc.title, tc.want, tc.modifications)
		}
	}
}
