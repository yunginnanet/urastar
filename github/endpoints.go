package github

import (
	"fmt"
)

const EndpointBase = "https://api.github.com/"
const EndpointUser = EndpointBase + "users/%s"
const EndpointStars = EndpointUser + "/starred"
const EndpointRepoSearch = EndpointBase + "search/repositories?q=%s&l=%s&page=%d&per_page=%d&sort=%v&order=%v"

type SortBy string

const (
	SortBestMatch SortBy = ""
	SortForks     SortBy = "forks"
	SortUpdated   SortBy = "updated"
	SortStars     SortBy = "stars"
)

type SortOrder string

const (
	Ascending  SortOrder = "asc"
	Descending SortOrder = "desc"
)

func URLFromName(s string) string {
	return fmt.Sprintf(EndpointUser, s)
}

func URLStarsFromName(s string) string {
	return fmt.Sprintf(EndpointStars, s)
}

func URLSearchRepos(query, language string, page, limit int, sort SortBy, order SortOrder) string {
	return fmt.Sprintf(
		EndpointRepoSearch,
		query, language,
		page, limit,
		sort, order)
}

func URLCVESearch(query string) string {
	return fmt.Sprintf(
		URLSearchRepos(query+" repo:CVEproject/cvelist extension:json", "", 1, 0, SortUpdated, Descending),
	)
}
