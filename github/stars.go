package github

import (
	"time"
)

//easyjson:json
type Stars []struct {
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    struct {
		Login     string `json:"login"`
		ID        int    `json:"id"`
		AvatarURL string `json:"avatar_url"`
		URL       string `json:"url"`
		HTMLURL   string `json:"html_url"`
		Type      string `json:"type"`
		SiteAdmin bool   `json:"site_admin"`
	} `json:"owner"`
	HTMLURL         string    `json:"html_url"`
	Description     *string   `json:"description"`
	Fork            bool      `json:"fork"`
	URL             string    `json:"url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	PushedAt        time.Time `json:"pushed_at"`
	GitURL          string    `json:"git_url"`
	SSHURL          string    `json:"ssh_url"`
	CloneURL        string    `json:"clone_url"`
	SvnURL          string    `json:"svn_url"`
	Homepage        *string   `json:"homepage"`
	Size            int       `json:"size"`
	StargazersCount int       `json:"stargazers_count"`
	WatchersCount   int       `json:"watchers_count"`
	Language        *string   `json:"language"`
	HasIssues       bool      `json:"has_issues"`
	HasProjects     bool      `json:"has_projects"`
	HasDownloads    bool      `json:"has_downloads"`
	HasWiki         bool      `json:"has_wiki"`
	HasPages        bool      `json:"has_pages"`
	ForksCount      int       `json:"forks_count"`
	Archived        bool      `json:"archived"`
	Disabled        bool      `json:"disabled"`
	OpenIssuesCount int       `json:"open_issues_count"`
	License         *struct {
		Key    string  `json:"key"`
		Name   string  `json:"name"`
		SpdxID string  `json:"spdx_id"`
		URL    *string `json:"url"`
		NodeID string  `json:"node_id"`
	} `json:"license"`
	AllowForking  bool     `json:"allow_forking"`
	IsTemplate    bool     `json:"is_template"`
	Topics        []string `json:"topics"`
	Visibility    string   `json:"visibility"`
	Forks         int      `json:"forks"`
	OpenIssues    int      `json:"open_issues"`
	Watchers      int      `json:"watchers"`
	DefaultBranch string   `json:"default_branch"`
}
