package git

import (
	"github.com/go-git/go-git/v5/plumbing/object"
	"sort"
)

func analyzeAuthors(authors map[string]AuthorInfo) AuthorsInfo {
	var authorsDetails []AuthorInfo
	var numAuthors, numCommits, numModifications int
	for _, v := range authors {
		numAuthors++
		numCommits += v.Commits
		numModifications += v.ModifiedLines
		authorsDetails = append(authorsDetails, v)
	}

	return AuthorsInfo{
		TotalAuthors:       numAuthors,
		TotalCommits:       numCommits,
		TotalModifications: numModifications,
		AuthorDetails:      authorsDetails,
	}
}

func analyzeModifications(pathModifications map[string]ModificationsInfo) []ModificationsInfo {
	var modifications []ModificationsInfo
	for _, currentModification := range pathModifications {
		analyzeModification(&currentModification)
		modifications = append(modifications, currentModification)
	}
	sort.Sort(sortModificationsSliceByAuthors(modifications))

	return modifications
}

func analyzeModification(currentModification *ModificationsInfo) {
	summary := make(map[string]ContributorSummary)
	for _, a := range currentModification.AuthorsSummary {
		entry, ok := summary[a.Name]
		if ok {
			entry.Modifications += a.Modifications
			summary[a.Name] = entry
		} else {
			summary[a.Name] = ContributorSummary{
				Name:          a.Name,
				Modifications: a.Modifications,
			}
		}
	}
	currentModification.AuthorsSummary = []ContributorSummary{}
	var keys []string
	for _, cs := range summary {
		keys = append(keys, cs.Name)
	}
	sort.Strings(keys)
	for _, k := range keys {
		s := summary[k]
		currentModification.AuthorsSummary = append(currentModification.AuthorsSummary, s)
		currentModification.TotalModifications += s.Modifications
	}
}

func analyzeLog(log object.CommitIter) *LogDetails {
	logDetails := &LogDetails{
		AuthorInfo:            make(map[string]AuthorInfo),
		PathModificationsInfo: make(map[string]ModificationsInfo),
		FileModificationsInfo: make(map[string]ModificationsInfo),
	}

	log.ForEach(func(c *object.Commit) error {
		analyzeCommit(c, logDetails)

		return nil
	})

	return logDetails
}

func analyzeCommit(c *object.Commit, logDetails *LogDetails) {
	if c == nil {
		return
	}
	if logDetails == nil {
		return
	}

	stats, _ := c.Stats()
	resolveAuthorsInfo(stats, c.Author.Name, logDetails.AuthorInfo)

	for _, fs := range stats {
		extractPathModifications(c.Author.Name, fs, logDetails.PathModificationsInfo)
		extractFileModifications(c.Author.Name, fs, logDetails.FileModificationsInfo)
	}
}
