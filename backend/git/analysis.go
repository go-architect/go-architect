package git

import (
	"fmt"
	"github.com/go-git/go-git/v5/plumbing/object"
	"os"
	"strings"
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
		for _, cs := range summary {
			currentModification.AuthorsSummary = append(currentModification.AuthorsSummary, cs)
			currentModification.TotalModifications += cs.Modifications
		}
		modifications = append(modifications, currentModification)
	}
	return modifications
}

func analyzeLog(log object.CommitIter) (map[string]AuthorInfo, map[string]ModificationsInfo, map[string]ModificationsInfo) {
	authors := make(map[string]AuthorInfo)
	pathModifications := make(map[string]ModificationsInfo)
	fileModifications := make(map[string]ModificationsInfo)

	log.ForEach(func(c *object.Commit) error {
		resolveAuthorsInfo(c, authors)

		s, _ := c.Stats()
		for _, fs := range s {
			filePath := fmt.Sprintf(".%s", string(os.PathSeparator))
			if strings.Contains(fs.Name, string(os.PathSeparator)) {
				lastIndex := strings.LastIndex(fs.Name, string(os.PathSeparator))
				filePath = fs.Name[0 : lastIndex+1]
			}
			entry, ok := pathModifications[filePath]
			if ok {
				entry.AuthorsSummary = append(entry.AuthorsSummary, ContributorSummary{
					Name:          c.Author.Name,
					Modifications: fs.Deletion + fs.Addition,
				})
				pathModifications[filePath] = entry
			} else {
				pathModifications[filePath] = ModificationsInfo{
					Source: filePath,
					AuthorsSummary: append(entry.AuthorsSummary, ContributorSummary{
						Name:          c.Author.Name,
						Modifications: fs.Deletion + fs.Addition,
					}),
				}
			}
			entry, ok = fileModifications[fs.Name]
			if ok {
				entry.AuthorsSummary = append(entry.AuthorsSummary, ContributorSummary{
					Name:          c.Author.Name,
					Modifications: fs.Deletion + fs.Addition,
				})
				fileModifications[fs.Name] = entry
			} else {
				fileModifications[fs.Name] = ModificationsInfo{
					Source: fs.Name,
					AuthorsSummary: append(entry.AuthorsSummary, ContributorSummary{
						Name:          c.Author.Name,
						Modifications: fs.Deletion + fs.Addition,
					}),
				}
			}
		}

		return nil
	})

	return authors, pathModifications, fileModifications
}
