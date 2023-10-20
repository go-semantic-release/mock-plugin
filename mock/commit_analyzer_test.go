package mock

import (
	"testing"

	"github.com/go-semantic-release/semantic-release/v2/pkg/semrel"
	"github.com/stretchr/testify/require"
)

func TestAnalyze(t *testing.T) {
	testCases := []struct {
		RawMessage string
		Type       string
		Change     *semrel.Change
	}{
		{"feat: add feature", "feat", &semrel.Change{Minor: true}},
		{"feat!: add feature", "feat!", &semrel.Change{Major: true}},
		{"fix: bug fix", "fix", &semrel.Change{Patch: true}},
		{"chore: update dependencies", "chore", &semrel.Change{}},
		{"update readme", "", &semrel.Change{}},
	}

	for _, tc := range testCases {
		t.Run(tc.RawMessage, func(t *testing.T) {
			commit := analyzeCommit(&semrel.RawCommit{
				RawMessage: tc.RawMessage,
			})
			require.Equal(t, tc.Type, commit.Type)
			require.Equal(t, tc.Change.Patch, commit.Change.Patch)
			require.Equal(t, tc.Change.Minor, commit.Change.Minor)
			require.Equal(t, tc.Change.Major, commit.Change.Major)
		})
	}
}
