package mock

import (
	"strings"

	"github.com/go-semantic-release/semantic-release/v2/pkg/analyzer"
	"github.com/go-semantic-release/semantic-release/v2/pkg/semrel"
)

var _ analyzer.CommitAnalyzer = &Mock{}

func (m *Mock) Analyze(commits []*semrel.RawCommit) []*semrel.Commit {
	m.Log.Printf("analyzing %d commits", len(commits))
	analyzedCommits := make([]*semrel.Commit, len(commits))
	for i, commit := range commits {
		analyzedCommits[i] = &semrel.Commit{
			SHA:     commit.SHA,
			Raw:     strings.Split(commit.RawMessage, "\n"),
			Type:    commit.RawMessage,
			Scope:   "mock",
			Message: commit.RawMessage,
			Change: &semrel.Change{
				Patch: commit.RawMessage == "fix",
				Minor: commit.RawMessage == "feat",
			},
			Annotations: commit.Annotations,
		}
	}
	return analyzedCommits
}
