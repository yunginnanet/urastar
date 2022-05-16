package github

func URLCVESearch(query string) string {
	return URLSearchRepos(query+" repo:CVEproject/cvelist extension:json", "", 1, 0, SortUpdated, Descending)
}
