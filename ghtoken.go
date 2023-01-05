package ghtoken

import (
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type GitHubToken struct {
	tr     http.RoundTripper
	rand   *rand.Rand
	tokens []string
}

func NewGitHubToken(tr http.RoundTripper) *GitHubToken {
	return &GitHubToken{
		tr:     tr,
		tokens: strings.Split(os.Getenv("GITHUB_TOKEN"), ","),
		rand:   rand.New(rand.NewSource(time.Now().Unix())),
	}
}

func (gh *GitHubToken) RoundTrip(req *http.Request) (*http.Response, error) {
	if len(gh.tokens) == 0 {
		return gh.tr.RoundTrip(req)
	}
	token := gh.tokens[gh.rand.Intn(len(gh.tokens))]
	req.Header.Add("Authorization", "Bearer "+token)
	return gh.tr.RoundTrip(req)
}
