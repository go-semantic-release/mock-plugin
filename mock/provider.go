package mock

import (
	"github.com/go-semantic-release/semantic-release/v2/pkg/provider"
	"github.com/go-semantic-release/semantic-release/v2/pkg/semrel"
)

var _ provider.Provider = &Mock{}

func (m *Mock) GetInfo() (*provider.RepositoryInfo, error) {
	m.Log.Println("getting repository info")
	return &provider.RepositoryInfo{
		Owner:         "owner",
		Repo:          "repo",
		DefaultBranch: "main",
		Private:       false,
	}, nil
}

func (m *Mock) GetCommits(fromSha, toSha string) ([]*semrel.RawCommit, error) {
	m.Log.Printf("getting commits from %s to %s", fromSha, toSha)
	annotations := map[string]string{
		"mock":        "true",
		"author_name": "mock",
	}
	return []*semrel.RawCommit{
		{
			SHA:         "1",
			RawMessage:  "feat: add feature",
			Annotations: annotations,
		},
		{
			SHA:         "2",
			RawMessage:  "fix: bug fix",
			Annotations: annotations,
		},
		{
			SHA:         "2",
			RawMessage:  "chore: update dependencies",
			Annotations: annotations,
		},
	}, nil
}

func (m *Mock) GetReleases(rawRe string) ([]*semrel.Release, error) {
	m.Log.Printf("getting releases (%s)", rawRe)
	return []*semrel.Release{
		{
			SHA:     "a",
			Version: "1.0.0",
		},
		{
			SHA:     "b",
			Version: "1.1.0",
		},
		{
			SHA:     "c",
			Version: "1.1.1",
		},
	}, nil
}

func (m *Mock) CreateRelease(rc *provider.CreateReleaseConfig) error {
	m.Log.Printf("creating release %s on branch %s", rc.NewVersion, rc.Branch)
	return nil
}
