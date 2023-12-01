package git

import (
	"fmt"
	"github.com/go-git/go-git/v5/plumbing/object"
	"os"
	"strings"
)

func resolveAuthorsInfo(s object.FileStats, authorName string, authors map[string]AuthorInfo) {
	var modifications int
	for _, fs := range s {
		modifications += fs.Addition + fs.Deletion
	}

	entry, ok := authors[authorName]
	if ok {
		entry.Commits++
		entry.ModifiedLines += modifications
		authors[authorName] = entry
	} else {
		authors[authorName] = AuthorInfo{
			Name:          authorName,
			Commits:       1,
			ModifiedLines: modifications,
		}
	}
}

func extractFileModifications(authorName string, fs object.FileStat, modificationsInfo map[string]ModificationsInfo) {
	entry, ok := modificationsInfo[fs.Name]
	if ok {
		entry.AuthorsSummary = append(entry.AuthorsSummary, ContributorSummary{
			Name:          authorName,
			Modifications: fs.Deletion + fs.Addition,
		})
		modificationsInfo[fs.Name] = entry
	} else {
		modificationsInfo[fs.Name] = ModificationsInfo{
			Source: fs.Name,
			AuthorsSummary: append(entry.AuthorsSummary, ContributorSummary{
				Name:          authorName,
				Modifications: fs.Deletion + fs.Addition,
			}),
		}
	}
}

func extractPathModifications(authorName string, fs object.FileStat, modificationsInfo map[string]ModificationsInfo) {
	filePath := resolveFilePath(fs)

	entry, ok := modificationsInfo[filePath]
	if ok {
		entry.AuthorsSummary = append(entry.AuthorsSummary, ContributorSummary{
			Name:          authorName,
			Modifications: fs.Deletion + fs.Addition,
		})
		modificationsInfo[filePath] = entry
	} else {
		modificationsInfo[filePath] = ModificationsInfo{
			Source: filePath,
			AuthorsSummary: append(entry.AuthorsSummary, ContributorSummary{
				Name:          authorName,
				Modifications: fs.Deletion + fs.Addition,
			}),
		}
	}
}

func resolveFilePath(fs object.FileStat) string {
	filePath := fmt.Sprintf(".%s", string(os.PathSeparator))
	if strings.Contains(fs.Name, string(os.PathSeparator)) {
		lastIndex := strings.LastIndex(fs.Name, string(os.PathSeparator))
		filePath = fs.Name[0 : lastIndex+1]
	}
	return filePath
}
