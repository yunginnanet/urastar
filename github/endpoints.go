package github

import (
	"fmt"
)

const EndpointBase = "https://api.github.com/"
const EndpointUser = EndpointBase + "users/%s"
const EndpointStars = EndpointUser + "/starred"
const EndpointRepoSearch = EndpointBase + "search/repositories?q=%s&l=%s&page=%d&per_page=%d&sort=%v&order=%v"

type RepoSearchSort string

const (
	SortBestMatch RepoSearchSort = ""
	SortForks     RepoSearchSort = "forks"
	SortUpdated   RepoSearchSort = "updated"
	SortStars     RepoSearchSort = "stars"
)

type SortOrder string

const (
	Ascending  SortOrder = "asc"
	Descending SortOrder = "desc"
)

func URLStarsFromName(s string) string {
	return fmt.Sprintf(EndpointStars, s)
}

func URLSearchRepos(query, language string, page, limit int, sort RepoSearchSort, order SortOrder) string {
	return fmt.Sprintf(
		EndpointRepoSearch,
		query, language,
		page, limit,
		sort, order)
}
