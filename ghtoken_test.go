package ghtoken

import (
	"context"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-github/github"
)

func TestGitHubToken(t *testing.T) {
	tr := NewGitHubToken(http.DefaultTransport)
	client := github.NewClient(&http.Client{Transport: tr})
	user, _, err := client.Users.Get(context.Background(), os.Getenv("GITHUB_ACTOR"))
	if err != nil {
		t.Fatal(err)
	}
	log.Println(user, err)
}
