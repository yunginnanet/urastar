package github

import "testing"

type test struct {
	name    string
	args    any
	want    string
	neederr bool
}

func TestURLFromName(t *testing.T) {
	var tests = []test{
		{
			name: "UserURLBasic",
			args: "yunginnanet",
			want: "https://api.github.com/users/yunginnanet",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := URLFromName(tt.args.(string)); got != tt.want {
				t.Errorf("URLFromName() = %v, want %v", got, tt.want)
			}
		})
	}
}

type urepoArgs struct {
	user  string
	rtype UserRepoType
	sort  UserRepoSort
	order SortOrder
	max   int
}

func TestURLUserRepos(t *testing.T) {
	var tests = []test{
		{
			name: "UserRepoURLBasicOwner",
			args: urepoArgs{
				user:  "yunginnanet",
				rtype: UserRepoTypeOwner,
				sort:  UserRepoSortUpdated,
				order: Descending,
				max:   100,
			},
			want: "https://api.github.com/users/yunginnanet/repos?type=owner&sort=updated&direction=desc&per_page=100",
		},
		{
			name: "UserRepoMax",
			args: urepoArgs{
				user:  "yunginnanet",
				rtype: UserRepoTypeOwner,
				sort:  UserRepoSortUpdated,
				order: Descending,
				max:   101,
			},
			want:    "",
			neederr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := URLUserRepos(tt.args.(urepoArgs).user, tt.args.(urepoArgs).rtype, tt.args.(urepoArgs).sort, tt.args.(urepoArgs).order, tt.args.(urepoArgs).max)
			if got != tt.want {
				t.Fatalf("URLUserRepos() = %v, want %v", got, tt.want)
			}
			if tt.neederr && err == nil {
				t.Fatalf("URLUserRepos() needed error got nil")
			}
			t.Logf("%s successful result: %s", tt.name, got)
		})
	}
}
