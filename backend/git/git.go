package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"strings"
	"time"
)

func LoadRepositoryTeamCohesion(path string, since *time.Time) (*VCSAnalysisInfo, error) {
	authors := make(map[string]AuthorInfo)
	pathModifications := make(map[string]ModificationsInfo)
	fileModifications := make(map[string]ModificationsInfo)
	var authorsDetails []AuthorInfo
	result := &VCSAnalysisInfo{
		Since:             since,
		ModificationsInfo: Modifications{},
	}
	fmt.Println("Search for VCSAnalysisInfo into: " + path)
	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Repo: %+v\n", repo)
	log, _ := repo.Log(&git.LogOptions{
		Order: git.LogOrderCommitterTime,
		Since: since,
	})
	log.ForEach(func(c *object.Commit) error {
		resolveAuthorsInfo(c, authors)

		s, _ := c.Stats()
		for _, fs := range s {
			fmt.Printf("LogDetail: (%s)[%s](%v/%v)\n", c.Author.Name, fs.Name, fs.Deletion, fs.Addition)
			filePath := "./"
			if strings.Contains(fs.Name, "/") {
				lastIndex := strings.LastIndex(fs.Name, "/")
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

	for _, modValue := range pathModifications {
		fmt.Printf("PathModification: %+v\n", modValue)
		summary := make(map[string]ContributorSummary)
		for _, a := range modValue.AuthorsSummary {
			fmt.Printf("Summarize author data: %+v\n", a)
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
		modValue.AuthorsSummary = []ContributorSummary{}
		for _, cs := range summary {
			modValue.AuthorsSummary = append(modValue.AuthorsSummary, cs)
			modValue.TotalModifications += cs.Modifications
		}
		result.ModificationsInfo.PathModifications = append(result.ModificationsInfo.PathModifications, modValue)
	}
	for _, modValue := range fileModifications {
		fmt.Printf("FileModification: %+v\n", modValue)
		summary := make(map[string]ContributorSummary)
		for _, a := range modValue.AuthorsSummary {
			fmt.Printf("Summarize author data: %+v\n", a)
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
		modValue.AuthorsSummary = []ContributorSummary{}
		for _, cs := range summary {
			modValue.AuthorsSummary = append(modValue.AuthorsSummary, cs)
			modValue.TotalModifications += cs.Modifications
		}
		result.ModificationsInfo.FileModifications = append(result.ModificationsInfo.FileModifications, modValue)
	}

	var numAuthors, numCommits, numModifications int
	for _, v := range authors {
		numAuthors++
		numCommits += v.Commits
		numModifications += v.ModifiedLines
		authorsDetails = append(authorsDetails, v)
	}
	result.AuthorsInfo = AuthorsInfo{
		TotalAuthors:       numAuthors,
		TotalCommits:       numCommits,
		TotalModifications: numModifications,
		AuthorDetails:      authorsDetails,
	}

	return result, nil
}
