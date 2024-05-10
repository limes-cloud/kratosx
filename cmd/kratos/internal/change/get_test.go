package change

import "testing"

func TestParseGithubURL(t *testing.T) {
	urls := []struct {
		url   string
		owner string
		repo  string
	}{
		{url: "https://github.com/limes-cloud/kratosx.git", owner: "go-kratos", repo: "kratos"},
		{url: "https://github.com/limes-cloud/kratosx", owner: "go-kratos", repo: "kratos"},
		{url: "git@github.com:limes-cloud/kratosx.git", owner: "go-kratos", repo: "kratos"},
	}
	for _, url := range urls {
		owner, repo := ParseGithubURL(url.url)
		if owner != url.owner {
			t.Fatalf("owner want: %s, got: %s", owner, url.owner)
		}
		if repo != url.repo {
			t.Fatalf("repo want: %s, got: %s", repo, url.repo)
		}
	}
}
