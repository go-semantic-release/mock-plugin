package mock

import (
	"strings"

	"github.com/go-semantic-release/semantic-release/v2/pkg/analyzer"
	"github.com/go-semantic-release/semantic-release/v2/pkg/semrel"
)

var _ analyzer.CommitAnalyzer = &Mock{}

func analyzeCommit(rawCommit *semrel.RawCommit) *semrel.Commit {
	splitMessages := strings.Split(rawCommit.RawMessage, "\n")
	cType, msg, foundType := strings.Cut(splitMessages[0], ":")
	if !foundType {
		msg = splitMessages[0]
		cType = ""
	}
	rawCommit.Annotations["mock"] = "true"
	return &semrel.Commit{
		SHA:         rawCommit.SHA,
		Raw:         splitMessages,
		Annotations: rawCommit.Annotations,
		Type:        cType,
		Scope:       "mock",
		Message:     msg,
		Change: &semrel.Change{
			Patch: cType == "fix",
			Minor: cType == "feat",
			Major: cType == "feat!",
		},
	}
}

func (m *Mock) Analyze(commits []*semrel.RawCommit) []*semrel.Commit {
	m.Log.Printf("analyzing %d commits", len(commits))
	analyzedCommits := make([]*semrel.Commit, len(commits))
	for i, commit := range commits {
		analyzedCommits[i] = analyzeCommit(commit)
	}
	return analyzedCommits
}
