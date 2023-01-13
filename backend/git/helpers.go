package git

import "github.com/go-git/go-git/v5/plumbing/object"

func resolveAuthorsInfo(c *object.Commit, authors map[string]AuthorInfo) {
	s, _ := c.Stats()
	var modifications int
	for _, fs := range s {
		modifications += fs.Addition + fs.Deletion
	}

	entry, ok := authors[c.Author.Name]
	if ok {
		entry.Commits++
		entry.ModifiedLines += modifications
		authors[c.Author.Name] = entry
	} else {
		authors[c.Author.Name] = AuthorInfo{
			Name:          c.Author.Name,
			Commits:       1,
			ModifiedLines: modifications,
		}
	}
}
