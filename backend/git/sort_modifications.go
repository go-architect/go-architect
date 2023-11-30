package git

type sortModificationsSliceByAuthors []ModificationsInfo

func (a sortModificationsSliceByAuthors) Len() int { return len(a) }
func (a sortModificationsSliceByAuthors) Less(i, j int) bool {
	if len(a[i].AuthorsSummary) == len(a[j].AuthorsSummary) {
		return a[i].TotalModifications > a[j].TotalModifications
	}
	return len(a[i].AuthorsSummary) < len(a[j].AuthorsSummary)
}
func (a sortModificationsSliceByAuthors) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
