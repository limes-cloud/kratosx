package change

import "testing"

func TestParseGithubURL(t *testing.T) {
	urls := []struct {
		url   string
		owner string
		repo  string
	}{
		{url: "https://github.com/limes-cloud/kratosx.git", owner: "limes-cloud", repo: "kratosx"},
		{url: "https://github.com/limes-cloud/kratosx", owner: "limes-cloud", repo: "kratosx"},
		{url: "git@github.com:limes-cloud/kratosx.git", owner: "limes-cloud", repo: "kratosx"},
	}
	for _, url := range urls {
		owner, repo := ParseGithubURL(url.url)
		if owner != url.owner {
			t.Fatalf("owner want: %s, got: %s", url.owner, owner)
		}
		if repo != url.repo {
			t.Fatalf("repo want: %s, got: %s", url.repo, repo)
		}
	}
}
