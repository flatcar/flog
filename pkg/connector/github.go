package connector

import (
	"context"
	"fmt"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/google/go-github/v39/github"
)

type GH struct{}

func (g *GH) FetchChangelog(name, version string) (string, string, error) {
	// TODO: private repository are not in the v1
	c := github.NewClient(nil)

	repo := strings.Split(name, "/")
	if len(repo) != 2 {
		return "", "", fmt.Errorf("name should be owner/repo")
	}

	release, _, err := c.Repositories.GetReleaseByTag(context.TODO(), repo[0], repo[1], version)
	if err != nil {
		return "", "", fmt.Errorf("getting release by tag: %w", err)
	}

	if release.Body == nil {
		return "", "", nil
	}

	changelog := markdown.ToHTML([]byte(*release.Body), nil, nil)

	return string(changelog), *release.HTMLURL, nil
}
