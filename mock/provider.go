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
	return []*semrel.RawCommit{
		{
			SHA:        "1",
			RawMessage: "feat: add feature",
			Annotations: map[string]string{
				"author_name": "mock",
			},
		},
		{
			SHA:        "2",
			RawMessage: "fix: bug fix",
			Annotations: map[string]string{
				"author_name": "mock",
			},
		},
	}, nil
}

func (m *Mock) GetReleases(re string) ([]*semrel.Release, error) {
	m.Log.Printf("getting releases: %s", re)
	return []*semrel.Release{
		{
			SHA:     "0",
			Version: "1.0.0",
		},
	}, nil
}

func (m *Mock) CreateRelease(rc *provider.CreateReleaseConfig) error {
	m.Log.Printf("creating release: %s", rc.NewVersion)
	return nil
}
