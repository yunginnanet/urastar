package github

import (
	"errors"
	"fmt"
)

const SuffixUserRepos = "/repos?type=%s&sort=%s&direction=%s&per_page=%d"

type UserRepoType string

const (
	UserRepoTypeOwner  UserRepoType = "owner"
	UserRepoTypeMember UserRepoType = "member"
	UserRepoTypeAll    UserRepoType = "all"
)

type UserRepoSort string

const (
	UserRepoSortName    UserRepoSort = "full_name"
	UserRepoSortCreated UserRepoSort = "created"
	UserRepoSortPushed  UserRepoSort = "pushed"
	UserRepoSortUpdated UserRepoSort = "updated"
)

func URLFromName(s string) string {
	return fmt.Sprintf(EndpointUser, s)
}

func URLUserRepos(user string, rtype UserRepoType, sort UserRepoSort, order SortOrder, max int) (string, error) {
	if max > 100 {
		return "", errors.New("specified max is too large, limit is 100")
	}
	return URLFromName(user) + fmt.Sprintf(SuffixUserRepos, rtype, sort, order, max), nil
}
