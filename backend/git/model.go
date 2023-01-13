package git

import "time"

type VCSAnalysisInfo struct {
	Since             *time.Time
	AuthorsInfo       AuthorsInfo
	ModificationsInfo Modifications
}

type Modifications struct {
	PathModifications []ModificationsInfo
	FileModifications []ModificationsInfo
}

type ModificationsInfo struct {
	Source             string
	TotalModifications int
	AuthorsSummary     []ContributorSummary
}

type ContributorSummary struct {
	Name          string
	Modifications int
}

type AuthorsInfo struct {
	TotalAuthors       int
	TotalCommits       int
	TotalModifications int
	AuthorDetails      []AuthorInfo
}

type AuthorInfo struct {
	Name          string
	Commits       int
	ModifiedLines int
}
