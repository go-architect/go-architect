package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"time"
)

func LoadRepositoryTeamCohesion(path string, since *time.Time) (*VCSAnalysisInfo, error) {
	fmt.Println("Search for VCSAnalysisInfo into: " + path)

	log, err := getGitLog(path, since)
	if err != nil {
		return nil, err
	}

	result := prepareResult(since, analyzeLog(log))

	return result, nil
}

func prepareResult(since *time.Time, logDetails *LogDetails) *VCSAnalysisInfo {
	if since == nil {
		return nil
	}
	if logDetails == nil {
		return nil
	}
	return &VCSAnalysisInfo{
		Since:       since,
		AuthorsInfo: analyzeAuthors(logDetails.AuthorInfo),
		ModificationsInfo: Modifications{
			PathModifications: analyzeModifications(logDetails.PathModificationsInfo),
			FileModifications: analyzeModifications(logDetails.FileModificationsInfo),
		},
	}
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
