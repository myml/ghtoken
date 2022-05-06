package ghtoken

import (
	"net/http"
	"os"
)

type GitHubToken struct {
	tr http.RoundTripper
}

func NewGitHubToken(tr http.RoundTripper) *GitHubToken {
	return &GitHubToken{tr: tr}
}

func (gh *GitHubToken) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+os.Getenv("GITHUB_TOKEN"))
	return gh.tr.RoundTrip(req)
}
