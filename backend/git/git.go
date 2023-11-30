package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"sort"
	"time"
)

func LoadRepositoryTeamCohesion(path string, since *time.Time) (*VCSAnalysisInfo, error) {
	result := &VCSAnalysisInfo{
		Since:             since,
		ModificationsInfo: Modifications{},
	}
	fmt.Println("Search for VCSAnalysisInfo into: " + path)

	log, err := getGitLog(path, since)
	if err != nil {
		return nil, err
	}
	authors, pathModifications, fileModifications := analyzeLog(log)

	result.ModificationsInfo.PathModifications = analyzeModifications(pathModifications)
	result.ModificationsInfo.FileModifications = analyzeModifications(fileModifications)
	result.AuthorsInfo = analyzeAuthors(authors)

	sort.Sort(sortModificationsSliceByAuthors(result.ModificationsInfo.FileModifications))
	sort.Sort(sortModificationsSliceByAuthors(result.ModificationsInfo.PathModifications))

	return result, nil
}

func getGitLog(path string, since *time.Time) (object.CommitIter, error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	log, err := repo.Log(&git.LogOptions{
		Order: git.LogOrderCommitterTime,
		Since: since,
	})

	return log, err
}
