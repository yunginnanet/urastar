// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package github

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	time "time"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson619ac83dDecodeGitTcpDirectKayosUrastarGithub(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "login":
			out.Login = string(in.String())
		case "id":
			out.Id = int(in.Int())
		case "avatar_url":
			out.AvatarURL = string(in.String())
		case "url":
			out.Url = string(in.String())
		case "html_url":
			out.HTMLURL = string(in.String())
		case "followers_url":
			out.FollowersUrl = string(in.String())
		case "following_url":
			out.FollowingUrl = string(in.String())
		case "gists_url":
			out.GistsUrl = string(in.String())
		case "starred_url":
			out.StarredUrl = string(in.String())
		case "subscriptions_url":
			out.SubscriptionsUrl = string(in.String())
		case "organizations_url":
			out.OrganizationsUrl = string(in.String())
		case "repos_url":
			out.ReposUrl = string(in.String())
		case "events_url":
			out.EventsUrl = string(in.String())
		case "received_events_url":
			out.ReceivedEventsUrl = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "site_admin":
			out.SiteAdmin = bool(in.Bool())
		case "name":
			out.Name = string(in.String())
		case "company":
			out.Company = string(in.String())
		case "blog":
			out.Blog = string(in.String())
		case "location":
			out.Location = string(in.String())
		case "email":
			if m, ok := out.Email.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Email.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Email = in.Interface()
			}
		case "hireable":
			if m, ok := out.Hireable.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Hireable.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Hireable = in.Interface()
			}
		case "bio":
			if m, ok := out.Bio.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Bio.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Bio = in.Interface()
			}
		case "twitter_username":
			out.TwitterUsername = string(in.String())
		case "public_repos":
			out.PublicRepos = int(in.Int())
		case "public_gists":
			out.PublicGists = int(in.Int())
		case "followers":
			out.Followers = int(in.Int())
		case "following":
			out.Following = int(in.Int())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updated_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson619ac83dEncodeGitTcpDirectKayosUrastarGithub(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Login != "" {
		const prefix string = ",\"login\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Login))
	}
	if in.Id != 0 {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Id))
	}
	if in.AvatarURL != "" {
		const prefix string = ",\"avatar_url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AvatarURL))
	}
	if in.Url != "" {
		const prefix string = ",\"url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Url))
	}
	if in.HTMLURL != "" {
		const prefix string = ",\"html_url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.HTMLURL))
	}
	if in.FollowersUrl != "" {
		const prefix string = ",\"followers_url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FollowersUrl))
	}
	if in.FollowingUrl != "" {
		const prefix string = ",\"following_url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FollowingUrl))
	}
	if in.GistsUrl != "" {
		const prefix string = ",\"gists_url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.GistsUrl))
	}
	if in.StarredUrl != "" {
		const prefix string = ",\"starred_url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.StarredUrl))
	}
	if in.SubscriptionsUrl != "" {
		const prefix string = ",\"subscriptions_url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SubscriptionsUrl))
	}
	if in.OrganizationsUrl != "" {
		const prefix string = ",\"organizations_url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.OrganizationsUrl))
	}
	if in.ReposUrl != "" {
		const prefix string = ",\"repos_url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ReposUrl))
	}
	if in.EventsUrl != "" {
		const prefix string = ",\"events_url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.EventsUrl))
	}
	if in.ReceivedEventsUrl != "" {
		const prefix string = ",\"received_events_url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ReceivedEventsUrl))
	}
	if in.Type != "" {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	if in.SiteAdmin {
		const prefix string = ",\"site_admin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.SiteAdmin))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.Company != "" {
		const prefix string = ",\"company\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Company))
	}
	if in.Blog != "" {
		const prefix string = ",\"blog\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Blog))
	}
	if in.Location != "" {
		const prefix string = ",\"location\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Location))
	}
	if in.Email != nil {
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if m, ok := in.Email.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Email.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Email))
		}
	}
	if in.Hireable != nil {
		const prefix string = ",\"hireable\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if m, ok := in.Hireable.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Hireable.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Hireable))
		}
	}
	if in.Bio != nil {
		const prefix string = ",\"bio\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if m, ok := in.Bio.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Bio.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Bio))
		}
	}
	if in.TwitterUsername != "" {
		const prefix string = ",\"twitter_username\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.TwitterUsername))
	}
	if in.PublicRepos != 0 {
		const prefix string = ",\"public_repos\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.PublicRepos))
	}
	if in.PublicGists != 0 {
		const prefix string = ",\"public_gists\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.PublicGists))
	}
	if in.Followers != 0 {
		const prefix string = ",\"followers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Followers))
	}
	if in.Following != 0 {
		const prefix string = ",\"following\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Following))
	}
	if true {
		const prefix string = ",\"created_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	if true {
		const prefix string = ",\"updated_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson619ac83dEncodeGitTcpDirectKayosUrastarGithub(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson619ac83dEncodeGitTcpDirectKayosUrastarGithub(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson619ac83dDecodeGitTcpDirectKayosUrastarGithub(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson619ac83dDecodeGitTcpDirectKayosUrastarGithub(l, v)
}
func easyjson619ac83dDecodeGitTcpDirectKayosUrastarGithub1(in *jlexer.Lexer, out *Stars) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Stars, 0, 0)
			} else {
				*out = Stars{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 struct {
				ID       int    `json:"id"`
				NodeID   string `json:"node_id"`
				Name     string `json:"name"`
				FullName string `json:"full_name"`
				Private  bool   `json:"private"`
				Owner    struct {
					Login     string `json:"login"`
					ID        int    `json:"id"`
					AvatarURL string `json:"avatar_url"`
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
			easyjson619ac83dDecode(in, &v1)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson619ac83dEncodeGitTcpDirectKayosUrastarGithub1(out *jwriter.Writer, in Stars) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			easyjson619ac83dEncode(out, v3)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Stars) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson619ac83dEncodeGitTcpDirectKayosUrastarGithub1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Stars) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson619ac83dEncodeGitTcpDirectKayosUrastarGithub1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Stars) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson619ac83dDecodeGitTcpDirectKayosUrastarGithub1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Stars) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson619ac83dDecodeGitTcpDirectKayosUrastarGithub1(l, v)
}
func easyjson619ac83dDecode(in *jlexer.Lexer, out *struct {
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    struct {
		Login     string `json:"login"`
		ID        int    `json:"id"`
		AvatarURL string `json:"avatar_url"`
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
}) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = int(in.Int())
		case "node_id":
			out.NodeID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "full_name":
			out.FullName = string(in.String())
		case "private":
			out.Private = bool(in.Bool())
		case "owner":
			easyjson619ac83dDecode1(in, &out.Owner)
		case "html_url":
			out.HTMLURL = string(in.String())
		case "description":
			if in.IsNull() {
				in.Skip()
				out.Description = nil
			} else {
				if out.Description == nil {
					out.Description = new(string)
				}
				*out.Description = string(in.String())
			}
		case "fork":
			out.Fork = bool(in.Bool())
		case "url":
			out.URL = string(in.String())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updated_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
			}
		case "pushed_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.PushedAt).UnmarshalJSON(data))
			}
		case "git_url":
			out.GitURL = string(in.String())
		case "ssh_url":
			out.SSHURL = string(in.String())
		case "clone_url":
			out.CloneURL = string(in.String())
		case "svn_url":
			out.SvnURL = string(in.String())
		case "homepage":
			if in.IsNull() {
				in.Skip()
				out.Homepage = nil
			} else {
				if out.Homepage == nil {
					out.Homepage = new(string)
				}
				*out.Homepage = string(in.String())
			}
		case "size":
			out.Size = int(in.Int())
		case "stargazers_count":
			out.StargazersCount = int(in.Int())
		case "watchers_count":
			out.WatchersCount = int(in.Int())
		case "language":
			if in.IsNull() {
				in.Skip()
				out.Language = nil
			} else {
				if out.Language == nil {
					out.Language = new(string)
				}
				*out.Language = string(in.String())
			}
		case "has_issues":
			out.HasIssues = bool(in.Bool())
		case "has_projects":
			out.HasProjects = bool(in.Bool())
		case "has_downloads":
			out.HasDownloads = bool(in.Bool())
		case "has_wiki":
			out.HasWiki = bool(in.Bool())
		case "has_pages":
			out.HasPages = bool(in.Bool())
		case "forks_count":
			out.ForksCount = int(in.Int())
		case "archived":
			out.Archived = bool(in.Bool())
		case "disabled":
			out.Disabled = bool(in.Bool())
		case "open_issues_count":
			out.OpenIssuesCount = int(in.Int())
		case "license":
			if in.IsNull() {
				in.Skip()
				out.License = nil
			} else {
				if out.License == nil {
					out.License = new(struct {
						Key    string  `json:"key"`
						Name   string  `json:"name"`
						SpdxID string  `json:"spdx_id"`
						URL    *string `json:"url"`
						NodeID string  `json:"node_id"`
					})
				}
				easyjson619ac83dDecode2(in, out.License)
			}
		case "allow_forking":
			out.AllowForking = bool(in.Bool())
		case "is_template":
			out.IsTemplate = bool(in.Bool())
		case "topics":
			if in.IsNull() {
				in.Skip()
				out.Topics = nil
			} else {
				in.Delim('[')
				if out.Topics == nil {
					if !in.IsDelim(']') {
						out.Topics = make([]string, 0, 4)
					} else {
						out.Topics = []string{}
					}
				} else {
					out.Topics = (out.Topics)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Topics = append(out.Topics, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "visibility":
			out.Visibility = string(in.String())
		case "forks":
			out.Forks = int(in.Int())
		case "open_issues":
			out.OpenIssues = int(in.Int())
		case "watchers":
			out.Watchers = int(in.Int())
		case "default_branch":
			out.DefaultBranch = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson619ac83dEncode(out *jwriter.Writer, in struct {
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    struct {
		Login     string `json:"login"`
		ID        int    `json:"id"`
		AvatarURL string `json:"avatar_url"`
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
}) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	if in.NodeID != "" {
		const prefix string = ",\"node_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.NodeID))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.FullName != "" {
		const prefix string = ",\"full_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FullName))
	}
	if in.Private {
		const prefix string = ",\"private\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Private))
	}
	if true {
		const prefix string = ",\"owner\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson619ac83dEncode1(out, in.Owner)
	}
	if in.HTMLURL != "" {
		const prefix string = ",\"html_url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.HTMLURL))
	}
	if in.Description != nil {
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Description))
	}
	if in.Fork {
		const prefix string = ",\"fork\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Fork))
	}
	if in.URL != "" {
		const prefix string = ",\"url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.URL))
	}
	if true {
		const prefix string = ",\"created_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	if true {
		const prefix string = ",\"updated_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	if true {
		const prefix string = ",\"pushed_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.PushedAt).MarshalJSON())
	}
	if in.GitURL != "" {
		const prefix string = ",\"git_url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.GitURL))
	}
	if in.SSHURL != "" {
		const prefix string = ",\"ssh_url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SSHURL))
	}
	if in.CloneURL != "" {
		const prefix string = ",\"clone_url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CloneURL))
	}
	if in.SvnURL != "" {
		const prefix string = ",\"svn_url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SvnURL))
	}
	if in.Homepage != nil {
		const prefix string = ",\"homepage\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Homepage))
	}
	if in.Size != 0 {
		const prefix string = ",\"size\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Size))
	}
	if in.StargazersCount != 0 {
		const prefix string = ",\"stargazers_count\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.StargazersCount))
	}
	if in.WatchersCount != 0 {
		const prefix string = ",\"watchers_count\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.WatchersCount))
	}
	if in.Language != nil {
		const prefix string = ",\"language\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Language))
	}
	if in.HasIssues {
		const prefix string = ",\"has_issues\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.HasIssues))
	}
	if in.HasProjects {
		const prefix string = ",\"has_projects\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.HasProjects))
	}
	if in.HasDownloads {
		const prefix string = ",\"has_downloads\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.HasDownloads))
	}
	if in.HasWiki {
		const prefix string = ",\"has_wiki\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.HasWiki))
	}
	if in.HasPages {
		const prefix string = ",\"has_pages\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.HasPages))
	}
	if in.ForksCount != 0 {
		const prefix string = ",\"forks_count\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.ForksCount))
	}
	if in.Archived {
		const prefix string = ",\"archived\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Archived))
	}
	if in.Disabled {
		const prefix string = ",\"disabled\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Disabled))
	}
	if in.OpenIssuesCount != 0 {
		const prefix string = ",\"open_issues_count\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.OpenIssuesCount))
	}
	if in.License != nil {
		const prefix string = ",\"license\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson619ac83dEncode2(out, *in.License)
	}
	if in.AllowForking {
		const prefix string = ",\"allow_forking\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.AllowForking))
	}
	if in.IsTemplate {
		const prefix string = ",\"is_template\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsTemplate))
	}
	if len(in.Topics) != 0 {
		const prefix string = ",\"topics\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Topics {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	if in.Visibility != "" {
		const prefix string = ",\"visibility\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Visibility))
	}
	if in.Forks != 0 {
		const prefix string = ",\"forks\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Forks))
	}
	if in.OpenIssues != 0 {
		const prefix string = ",\"open_issues\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.OpenIssues))
	}
	if in.Watchers != 0 {
		const prefix string = ",\"watchers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Watchers))
	}
	if in.DefaultBranch != "" {
		const prefix string = ",\"default_branch\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.DefaultBranch))
	}
	out.RawByte('}')
}
func easyjson619ac83dDecode2(in *jlexer.Lexer, out *struct {
	Key    string  `json:"key"`
	Name   string  `json:"name"`
	SpdxID string  `json:"spdx_id"`
	URL    *string `json:"url"`
	NodeID string  `json:"node_id"`
}) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "key":
			out.Key = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "spdx_id":
			out.SpdxID = string(in.String())
		case "url":
			if in.IsNull() {
				in.Skip()
				out.URL = nil
			} else {
				if out.URL == nil {
					out.URL = new(string)
				}
				*out.URL = string(in.String())
			}
		case "node_id":
			out.NodeID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson619ac83dEncode2(out *jwriter.Writer, in struct {
	Key    string  `json:"key"`
	Name   string  `json:"name"`
	SpdxID string  `json:"spdx_id"`
	URL    *string `json:"url"`
	NodeID string  `json:"node_id"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Key != "" {
		const prefix string = ",\"key\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Key))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.SpdxID != "" {
		const prefix string = ",\"spdx_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SpdxID))
	}
	if in.URL != nil {
		const prefix string = ",\"url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.URL))
	}
	if in.NodeID != "" {
		const prefix string = ",\"node_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.NodeID))
	}
	out.RawByte('}')
}
func easyjson619ac83dDecode1(in *jlexer.Lexer, out *struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	AvatarURL string `json:"avatar_url"`
	HTMLURL   string `json:"html_url"`
	Type      string `json:"type"`
	SiteAdmin bool   `json:"site_admin"`
}) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "login":
			out.Login = string(in.String())
		case "id":
			out.ID = int(in.Int())
		case "avatar_url":
			out.AvatarURL = string(in.String())
		case "html_url":
			out.HTMLURL = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "site_admin":
			out.SiteAdmin = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson619ac83dEncode1(out *jwriter.Writer, in struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	AvatarURL string `json:"avatar_url"`
	HTMLURL   string `json:"html_url"`
	Type      string `json:"type"`
	SiteAdmin bool   `json:"site_admin"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Login != "" {
		const prefix string = ",\"login\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Login))
	}
	if in.ID != 0 {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.ID))
	}
	if in.AvatarURL != "" {
		const prefix string = ",\"avatar_url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AvatarURL))
	}
	if in.HTMLURL != "" {
		const prefix string = ",\"html_url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.HTMLURL))
	}
	if in.Type != "" {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	if in.SiteAdmin {
		const prefix string = ",\"site_admin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.SiteAdmin))
	}
	out.RawByte('}')
}
